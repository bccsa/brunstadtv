package graph

import (
	"context"
	"github.com/bcc-code/bcc-media-platform/backend/common"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/model"
	"github.com/bcc-code/bcc-media-platform/backend/items/collection"
	"github.com/bcc-code/bcc-media-platform/backend/user"
	"github.com/bcc-code/bcc-media-platform/backend/utils"
	"github.com/bcc-code/mediabank-bridge/log"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"strconv"
)

func preloadEntryLoaders(ctx context.Context, loaders *common.BatchLoaders, entries []collection.Entry) {
	for _, e := range entries {
		switch e.Collection {
		case common.CollectionShows:
			loaders.ShowLoader.Load(ctx, utils.AsInt(e.ID))
		case common.CollectionSeasons:
			loaders.SeasonLoader.Load(ctx, utils.AsInt(e.ID))
		case common.CollectionEpisodes:
			loaders.EpisodeLoader.Load(ctx, utils.AsInt(e.ID))
		case common.CollectionPages:
			loaders.PageLoader.Load(ctx, utils.AsInt(e.ID))
		case common.CollectionSections:
			loaders.SectionLoader.Load(ctx, utils.AsInt(e.ID))
		case common.CollectionStudyTopics:
			loaders.StudyTopicLoader.Load(ctx, utils.AsUuid(e.ID))
		case common.CollectionLinks:
			loaders.LinkLoader.Load(ctx, utils.AsInt(e.ID))
		case common.CollectionPlaylists:
			loaders.PlaylistLoader.Load(ctx, utils.AsUuid(e.ID))
		case common.CollectionGames:
			loaders.GameLoader.Load(ctx, utils.AsUuid(e.ID))
		}
	}
}

func sectionStyleToImageStyle(style string) common.ImageStyle {
	switch style {
	case "icon_grid", "icons":
		return common.ImageStyleIcon
	case "poster_grid", "posters":
		return common.ImageStylePoster
	case "featured":
		return common.ImageStyleFeatured
	default:
		return common.ImageStyleDefault
	}
}

func filterWithIds(col *common.Collection, entries []collection.Entry, ids []*int) []collection.Entry {
	limit := 20
	if col.Filter != nil && col.Filter.Limit != nil {
		limit = *col.Filter.Limit
	}
	var newEntries []collection.Entry
	for _, id := range utils.PointerArrayToArray(ids) {
		entry, found := lo.Find(entries, func(e collection.Entry) bool {
			return e.Collection == common.CollectionEpisodes && e.ID == strconv.Itoa(id)
		})
		if found {
			newEntries = append(newEntries, entry)
			if len(newEntries) >= limit {
				break
			}
		}
	}
	return newEntries
}

func resolveContinueWatchingCollection(ctx context.Context, ls *common.BatchLoaders) ([]*int, error) {
	ginCtx, err := utils.GinCtx(ctx)
	if err != nil {
		return nil, err
	}
	profile := user.GetProfileFromCtx(ginCtx)
	if profile == nil {
		return nil, err
	}
	ids, err := ls.EpisodeProgressLoader.Get(ctx, profile.ID)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func resolveMyListCollection(ctx context.Context, ls *common.BatchLoaders) ([]*int, error) {
	ginCtx, err := utils.GinCtx(ctx)
	if err != nil {
		return nil, err
	}
	profile := user.GetProfileFromCtx(ginCtx)
	if profile == nil {
		return nil, nil
	}
	myListID, err := ls.ProfileMyListCollectionID.Get(ctx, profile.ID)
	if err != nil || myListID == nil {
		return nil, err
	}
	entryIDs, err := ls.UserCollectionEntryIDsLoader.Get(ctx, *myListID)
	if err != nil {
		return nil, err
	}
	collectionEntries, err := ls.UserCollectionEntryLoader.GetMany(ctx, utils.PointerArrayToArray(entryIDs))
	if err != nil {
		return nil, err
	}
	// UUIDs, but only for episodes
	uuids := lo.Map(lo.Filter(collectionEntries, func(i *common.UserCollectionEntry, _ int) bool {
		return i.Type == "episode"
	}), func(i *common.UserCollectionEntry, _ int) uuid.UUID {
		return i.ItemID
	})
	ids, err := ls.EpisodeIDFromUuidLoader.GetMany(ctx, uuids)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func mapCollectionEntriesToSectionItems(ctx context.Context, ls *common.BatchLoaders, entries []collection.Entry, imageStyle string, numberInTitle bool) ([]*model.SectionItem, error) {
	var items []*model.SectionItem
	for _, e := range entries {
		var item *model.SectionItem
		switch e.Collection {
		case common.CollectionPages:
			i, err := ls.PageLoader.Get(ctx, utils.AsInt(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				log.L.Debug().Str("id", e.ID).Str("type", e.Collection.Value).Msg("Item with id not found")
				continue
			}
			item = model.PageSectionItemFrom(ctx, i, e.Sort, imageStyle)
		case common.CollectionShows:
			i, err := ls.ShowLoader.Get(ctx, utils.AsInt(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				log.L.Debug().Str("id", e.ID).Str("type", e.Collection.Value).Msg("Item with id not found")
				continue
			}
			item = model.ShowSectionItemFrom(ctx, i, e.Sort, imageStyle)
		case common.CollectionSeasons:
			i, err := ls.SeasonLoader.Get(ctx, utils.AsInt(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				log.L.Debug().Str("id", e.ID).Str("type", e.Collection.Value).Msg("Item with id not found")
				continue
			}
			item = model.SeasonSectionItemFrom(ctx, i, e.Sort, imageStyle)
		case common.CollectionEpisodes:
			i, err := ls.EpisodeLoader.Get(ctx, utils.AsInt(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				log.L.Debug().Str("id", e.ID).Str("type", e.Collection.Value).Msg("Item with id not found")
				continue
			}
			item = model.EpisodeSectionItemFrom(ctx, i, e.Sort, imageStyle, numberInTitle)
		case common.CollectionLinks:
			i, err := ls.LinkLoader.Get(ctx, utils.AsInt(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				log.L.Debug().Str("id", e.ID).Str("type", e.Collection.Value).Msg("Item with id not found")
				continue
			}
			item = model.LinkSectionItemFrom(ctx, i, e.Sort, imageStyle)
		case common.CollectionStudyTopics:
			i, err := ls.StudyTopicLoader.Get(ctx, utils.AsUuid(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				log.L.Debug().Str("id", e.ID).Str("type", e.Collection.Value).Msg("Item with id not found")
				continue
			}
			item = model.StudyTopicSectionItemFrom(ctx, i, e.Sort, imageStyle)
		case common.CollectionGames:
			i, err := ls.GameLoader.Get(ctx, utils.AsUuid(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				log.L.Debug().Str("id", e.ID).Str("type", e.Collection.Value).Msg("Item with id not found")
				continue
			}
			item = model.GameSectionItemFrom(ctx, i, e.Sort, imageStyle)
		case common.CollectionPlaylists:
			i, err := ls.PlaylistLoader.Get(ctx, utils.AsUuid(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				log.L.Debug().Str("id", e.ID).Str("type", e.Collection.Value).Msg("Item with id not found")
				continue
			}
			item = model.PlaylistSectionItemFrom(ctx, i, e.Sort, imageStyle)
		}
		if item != nil {
			items = append(items, item)
		}
	}
	return items, nil
}

func sectionCollectionEntryResolver(
	ctx context.Context,
	ls *common.BatchLoaders,
	filteredLoaders *common.FilteredLoaders,
	section *common.Section,
	first *int,
	offset *int,
) (*utils.PaginationResult[*model.SectionItem], error) {
	if !section.CollectionID.Valid {
		return &utils.PaginationResult[*model.SectionItem]{}, nil
	}

	collectionId := int(section.CollectionID.ValueOrZero())

	col, err := ls.CollectionLoader.Get(ctx, collectionId)
	if err != nil {
		return nil, err
	}

	entries, err := collection.GetCollectionEntries(ctx, ls, filteredLoaders, collectionId)
	if err != nil {
		return nil, err
	}

	switch col.AdvancedType.String {
	case "continue_watching":
		ids, err := resolveContinueWatchingCollection(ctx, ls)
		if err != nil {
			return nil, err
		}
		entries = filterWithIds(col, entries, ids)
	case "my_list":
		ids, err := resolveMyListCollection(ctx, ls)
		if err != nil {
			return nil, err
		}
		entries = filterWithIds(col, entries, ids)
	}

	pagination := utils.Paginate(entries, first, offset, nil)

	imageStyle := sectionStyleToImageStyle(section.Style)

	preloadEntryLoaders(ctx, ls, pagination.Items)

	items, err := mapCollectionEntriesToSectionItems(ctx, ls, pagination.Items, imageStyle, col.NumberInTitles)
	if err != nil {
		return nil, err
	}

	return &utils.PaginationResult[*model.SectionItem]{
		Total:  pagination.Total,
		First:  pagination.First,
		Offset: pagination.Offset,
		Items:  items,
	}, nil
}

func sectionCollectionItemResolver(ctx context.Context, r *Resolver, id string, first *int, offset *int) (*model.SectionItemPagination, error) {
	int64ID, _ := strconv.ParseInt(id, 10, 32)

	section, err := r.Loaders.SectionLoader.Get(ctx, int(int64ID))
	if err != nil {
		return nil, err
	}

	pagination, err := sectionCollectionEntryResolver(ctx, r.Loaders, r.FilteredLoaders(ctx), section, first, offset)
	if err != nil {
		return nil, err
	}

	return &model.SectionItemPagination{
		Total:  pagination.Total,
		First:  pagination.First,
		Offset: pagination.Offset,
		Items:  pagination.Items,
	}, nil
}
