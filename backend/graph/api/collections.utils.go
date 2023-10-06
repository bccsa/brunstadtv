package graph

import (
	"context"
	"github.com/bcc-code/bcc-media-platform/backend/common"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/model"
	"github.com/bcc-code/bcc-media-platform/backend/items/collection"
	"github.com/bcc-code/bcc-media-platform/backend/utils"
	"github.com/samber/lo"
)

func collectionEntriesToModels(ctx context.Context, ls *common.BatchLoaders, entries []collection.Entry) ([]model.Item, error) {
	preloadEntryLoaders(ctx, ls, entries)

	var items []model.Item
	for _, e := range entries {
		switch e.Collection {
		case common.CollectionShows:
			i, err := ls.ShowLoader.Get(ctx, utils.AsInt(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				continue
			}
			items = append(items, model.ShowFrom(ctx, i))
		case common.CollectionSeasons:
			i, err := ls.SeasonLoader.Get(ctx, utils.AsInt(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				continue
			}
			items = append(items, model.SeasonFrom(ctx, i))
		case common.CollectionEpisodes:
			i, err := ls.EpisodeLoader.Get(ctx, utils.AsInt(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				continue
			}
			items = append(items, model.EpisodeFrom(ctx, i))
		case common.CollectionPlaylists:
			i, err := ls.PlaylistLoader.Get(ctx, utils.AsUuid(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				continue
			}
			items = append(items, model.PlaylistFrom(ctx, i))
		case common.CollectionGames:
			i, err := ls.GameLoader.Get(ctx, utils.AsUuid(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				continue
			}
			items = append(items, model.GameFrom(ctx, i))
		case common.CollectionSections:
			i, err := ls.SectionLoader.Get(ctx, utils.AsInt(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				continue
			}
			items = append(items, model.SectionFrom(ctx, i))
		case common.CollectionStudyTopics:
			i, err := ls.StudyTopicLoader.Get(ctx, utils.AsUuid(e.ID))
			if err != nil {
				return nil, err
			}
			if i == nil {
				continue
			}
			items = append(items, model.StudyTopicFrom(ctx, i))
		}
	}
	return items, nil
}

func (r *Resolver) getEntriesFromCollection(ctx context.Context, collectionID int) ([]collection.Entry, error) {
	col, err := r.Loaders.CollectionLoader.Get(ctx, collectionID)
	if err != nil {
		return nil, err
	}

	entries, err := collection.GetCollectionEntries(ctx, r.Loaders, r.GetFilteredLoaders(ctx), collectionID)
	if err != nil {
		return nil, err
	}

	switch col.AdvancedType.String {
	case "continue_watching":
		ids, err := resolveContinueWatchingCollection(ctx, r.Loaders)
		if err != nil {
			return nil, err
		}
		entries = filterWithIds(col, entries, ids)
	case "my_list":
		ids, err := resolveMyListCollection(ctx, r.Loaders)
		if err != nil {
			return nil, err
		}
		entries = filterWithIds(col, entries, ids)
	}

	return entries, nil
}

func getItemsPageAs[T any](ctx context.Context, r *Resolver, collectionID int, first, offset *int, collections ...common.ItemCollection) (*utils.PaginationResult[T], error) {
	entries, err := r.getEntriesFromCollection(ctx, collectionID)
	if err != nil {
		return nil, err
	}

	// In case we want to make sure to only use entries of a certain collection
	if len(collections) > 0 {
		entries = lo.Filter(entries, func(e collection.Entry, _ int) bool {
			return lo.Contains(collections, e.Collection)
		})
	}

	pagination := utils.Paginate(entries, first, offset, nil)

	items, err := collectionEntriesToModels(ctx, r.Loaders, pagination.Items)
	if err != nil {
		return nil, err
	}

	var result []T
	for _, item := range items {
		if v, ok := item.(T); ok {
			result = append(result, v)
		}
	}

	return &utils.PaginationResult[T]{
		Total:  pagination.Total,
		First:  pagination.First,
		Offset: pagination.Offset,
		Items:  result,
	}, nil
}

func (r *Resolver) getPlaylistItemsPage(ctx context.Context, collectionID int, first, offset *int) (*model.PlaylistItemPagination, error) {
	p, err := getItemsPageAs[model.PlaylistItem](ctx, r, collectionID, first, offset, common.CollectionEpisodes)
	if err != nil {
		return nil, err
	}
	return &model.PlaylistItemPagination{
		Total:  p.Total,
		First:  p.First,
		Offset: p.Offset,
		Items:  p.Items,
	}, nil
}
