package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/graph/api/generated"
	"github.com/bcc-code/brunstadtv/backend/graph/api/model"
	"github.com/bcc-code/brunstadtv/backend/utils"
	"github.com/samber/lo"
)

// Entries is the resolver for the entries field.
func (r *userCollectionResolver) Entries(ctx context.Context, obj *model.UserCollection, first *int, offset *int) (*model.UserCollectionEntryPagination, error) {
	ids, err := r.Loaders.UserCollectionEntryIDsLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	page := utils.Paginate(ids, first, offset, nil)
	entries, err := r.Loaders.UserCollectionEntryLoader.GetMany(ctx, utils.PointerArrayToArray(page.Items))
	if err != nil {
		return nil, err
	}
	return &model.UserCollectionEntryPagination{
		Total:  page.Total,
		Offset: page.Offset,
		First:  page.First,
		Items: lo.Map(entries, func(i *common.UserCollectionEntry, _ int) *model.UserCollectionEntry {
			return &model.UserCollectionEntry{
				ID: i.ID.String(),
			}
		}),
	}, nil
}

// Item is the resolver for the item field.
func (r *userCollectionEntryResolver) Item(ctx context.Context, obj *model.UserCollectionEntry) (model.UserCollectionEntryItem, error) {
	e, err := r.Loaders.UserCollectionEntryLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	switch e.Type {
	case "show":
		id, err := r.Loaders.ShowIDFromUuidLoader.Get(ctx, e.ItemID)
		if err != nil {
			return nil, err
		}
		if id == nil {
			return &model.Show{}, nil
		}
		return r.QueryRoot().Show(ctx, strconv.Itoa(*id))
	case "episode":
		id, err := r.Loaders.EpisodeIDFromUuidLoader.Get(ctx, e.ItemID)
		if err != nil {
			return nil, err
		}
		if id == nil {
			return &model.Episode{}, nil
		}
		return r.QueryRoot().Episode(ctx, strconv.Itoa(*id), nil)
	}
	return nil, common.ErrItemNotFound
}

// UserCollection returns generated.UserCollectionResolver implementation.
func (r *Resolver) UserCollection() generated.UserCollectionResolver {
	return &userCollectionResolver{r}
}

// UserCollectionEntry returns generated.UserCollectionEntryResolver implementation.
func (r *Resolver) UserCollectionEntry() generated.UserCollectionEntryResolver {
	return &userCollectionEntryResolver{r}
}

type userCollectionResolver struct{ *Resolver }
type userCollectionEntryResolver struct{ *Resolver }
