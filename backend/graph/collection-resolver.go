package graph

import (
	"context"
	"github.com/bcc-code/brunstadtv/backend/common"
	gqlmodel "github.com/bcc-code/brunstadtv/backend/graph/model"
	"github.com/bcc-code/brunstadtv/backend/items/collection"
	"github.com/bcc-code/brunstadtv/backend/user"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/samber/lo"
	"strconv"
)

func itemsResolverForIntID[t any, r any](ctx context.Context, id string, loader *dataloader.Loader[int, []*t], converter func(context.Context, *t) r) ([]r, error) {
	intID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, err
	}
	return itemsResolverFor(ctx, int(intID), loader, converter)
}

func collectionItemResolver(ctx context.Context, r *Resolver, id string) ([]gqlmodel.Item, error) {
	int64ID, _ := strconv.ParseInt(id, 10, 32)

	section, err := common.GetFromLoaderByID(ctx, r.Loaders.SectionLoader, int(int64ID))
	if err != nil {
		return nil, err
	}

	if !section.CollectionID.Valid {
		return nil, nil
	}

	items, err := collection.GetCollectionItems(ctx, r.Loaders, int(section.CollectionID.ValueOrZero()))
	if err != nil {
		return nil, err
	}

	items = lo.Filter(items, func(i collection.Item, _ int) bool {
		if t, ok := i.(restrictedItem); ok {
			return user.ValidateAccess(ctx, t) == nil
		}
		return true
	})
	return lo.Map(items, func(i collection.Item, _ int) gqlmodel.Item {
		switch t := i.(type) {
		case *common.Page:
			return gqlmodel.PageItemFrom(ctx, t)
		case *common.Show:
			return gqlmodel.ShowItemFrom(ctx, t)
		case *common.Season:
			return gqlmodel.SeasonItemFrom(ctx, t)
		case *common.Episode:
			return gqlmodel.EpisodeItemFrom(ctx, t)
		}
		return gqlmodel.PageItem{}
	}), nil
}
