package graph

import (
	"context"
	"github.com/bcc-code/brunstadtv/backend/common"
	gqlmodel "github.com/bcc-code/brunstadtv/backend/graph/model"
	"github.com/bcc-code/brunstadtv/backend/items/collection"
	"github.com/bcc-code/brunstadtv/backend/user"
	"github.com/bcc-code/brunstadtv/backend/utils"
	"github.com/bcc-code/mediabank-bridge/log"
	"strconv"
)

func preloadLoaders(ctx context.Context, loaders *common.BatchLoaders, entries []collection.Entry) {
	for _, e := range entries {
		switch e.Type {
		case "show":
			loaders.ShowLoader.Load(ctx, e.ID)
		case "season":
			loaders.SeasonLoader.Load(ctx, e.ID)
		case "episode":
			loaders.EpisodeLoader.Load(ctx, e.ID)
		case "page":
			loaders.PageLoader.Load(ctx, e.ID)
		case "section":
			loaders.SectionLoader.Load(ctx, e.ID)
		}
	}
}

func collectionEntryResolver(ctx context.Context, loaders *common.BatchLoaders, collectionId int, first *int, offset *int) (*utils.PaginationResult[gqlmodel.Item], error) {
	entries, err := collection.GetCollectionEntries(ctx, loaders, collectionId)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		switch e.Type {
		case common.TypePage:
			loaders.PagePermissionLoader.Load(ctx, e.ID)
		case common.TypeShow:
			loaders.ShowPermissionLoader.Load(ctx, e.ID)
		case common.TypeSeason:
			loaders.SeasonPermissionLoader.Load(ctx, e.ID)
		case common.TypeEpisode:
			loaders.EpisodePermissionLoader.Load(ctx, e.ID)
		case common.TypeSection:
			loaders.SectionPermissionLoader.Load(ctx, e.ID)
		}
	}

	var returnEntries []collection.Entry
	for _, e := range entries {
		var success bool
		switch e.Type {
		case common.TypePage:
			success = user.ValidateAccess(ctx, loaders.PagePermissionLoader, e.ID) == nil
		case common.TypeShow:
			success = user.ValidateAccess(ctx, loaders.ShowPermissionLoader, e.ID) == nil
		case common.TypeSeason:
			success = user.ValidateAccess(ctx, loaders.SeasonPermissionLoader, e.ID) == nil
		case common.TypeEpisode:
			success = user.ValidateAccess(ctx, loaders.EpisodePermissionLoader, e.ID) == nil
		case common.TypeSection:
			success = user.ValidateAccess(ctx, loaders.SectionPermissionLoader, e.ID) == nil
		default:
			log.L.Error().Str("type", string(e.Type)).Msg("Invalid/unsupported entry type in collection")
		}
		if success {
			returnEntries = append(returnEntries, e)
		}
	}

	pagination := utils.Paginate(returnEntries, first, offset)

	preloadLoaders(ctx, loaders, pagination.Items)

	var items []gqlmodel.Item
	for _, e := range pagination.Items {
		var item gqlmodel.Item
		switch e.Type {
		case "page":
			i, err := common.GetFromLoaderByID(ctx, loaders.PageLoader, e.ID)
			if err != nil {
				return nil, err
			}
			item = gqlmodel.PageItemFrom(ctx, i, e.Sort)
		case "show":
			i, err := common.GetFromLoaderByID(ctx, loaders.ShowLoader, e.ID)
			if err != nil {
				return nil, err
			}
			item = gqlmodel.ShowItemFrom(ctx, i, e.Sort)
		case "season":
			i, err := common.GetFromLoaderByID(ctx, loaders.SeasonLoader, e.ID)
			if err != nil {
				return nil, err
			}
			item = gqlmodel.SeasonItemFrom(ctx, i, e.Sort)
		case "episode":
			i, err := common.GetFromLoaderByID(ctx, loaders.EpisodeLoader, e.ID)
			if err != nil {
				return nil, err
			}
			item = gqlmodel.EpisodeItemFrom(ctx, i, e.Sort)
		}
		if item != nil {
			items = append(items, item)
		}
	}

	return &utils.PaginationResult[gqlmodel.Item]{
		Total:  pagination.Total,
		First:  pagination.First,
		Offset: pagination.Offset,
		Items:  items,
	}, nil
}

func collectionItemResolver(ctx context.Context, r *Resolver, id string, first *int, offset *int) (*utils.PaginationResult[gqlmodel.Item], error) {
	int64ID, _ := strconv.ParseInt(id, 10, 32)

	section, err := common.GetFromLoaderByID(ctx, r.Loaders.SectionLoader, int(int64ID))
	if err != nil {
		return nil, err
	}

	if !section.CollectionID.Valid {
		return nil, nil
	}

	return collectionEntryResolver(ctx, r.Loaders, int(section.CollectionID.Int64), first, offset)
}

func collectionItemResolverFromCollection(ctx context.Context, r *Resolver, id string, first *int, offset *int) (*utils.PaginationResult[gqlmodel.Item], error) {
	int64ID, _ := strconv.ParseInt(id, 10, 32)

	return collectionEntryResolver(ctx, r.Loaders, int(int64ID), first, offset)
}
