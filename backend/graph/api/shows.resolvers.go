package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	merry "github.com/ansel1/merry/v2"
	"github.com/bcc-code/brunstadtv/backend/graph/api/generated"
	"github.com/bcc-code/brunstadtv/backend/graph/api/model"
	"github.com/bcc-code/brunstadtv/backend/items/show"
	"github.com/bcc-code/brunstadtv/backend/utils"
)

// Image is the resolver for the image field.
func (r *showResolver) Image(ctx context.Context, obj *model.Show, style *model.ImageStyle) (*string, error) {
	e, err := r.Loaders.ShowLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return nil, err
	}
	return imageOrFallback(ctx, e.Images, style), nil
}

// EpisodeCount is the resolver for the episodeCount field.
func (r *showResolver) EpisodeCount(ctx context.Context, obj *model.Show) (int, error) {
	seasonIDs, err := r.FilteredLoaders(ctx).SeasonsLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return 0, err
	}
	el := r.FilteredLoaders(ctx).EpisodesLoader
	for _, id := range seasonIDs {
		el.Load(ctx, *id)
	}

	count := 0
	for _, id := range seasonIDs {
		episodeIDs, err := el.Get(ctx, *id)
		if err != nil {
			return 0, err
		}
		count += len(episodeIDs)
	}
	return count, nil
}

// SeasonCount is the resolver for the seasonCount field.
func (r *showResolver) SeasonCount(ctx context.Context, obj *model.Show) (int, error) {
	seasonIDs, err := r.FilteredLoaders(ctx).SeasonsLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return 0, err
	}
	return len(seasonIDs), nil
}

// Seasons is the resolver for the seasons field.
func (r *showResolver) Seasons(ctx context.Context, obj *model.Show, first *int, offset *int, dir *string) (*model.SeasonPagination, error) {
	intID, err := strconv.ParseInt(obj.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	itemIDs, err := r.FilteredLoaders(ctx).SeasonsLoader.Get(ctx, int(intID))
	if err != nil {
		return nil, err
	}

	page := utils.Paginate(itemIDs, first, offset, dir)

	seasons, err := r.Loaders.SeasonLoader.GetMany(ctx, utils.PointerIntArrayToIntArray(page.Items))
	if err != nil {
		return nil, err
	}
	return &model.SeasonPagination{
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
		Items:  utils.MapWithCtx(ctx, seasons, model.SeasonFrom),
	}, nil
}

// DefaultEpisode is the resolver for the defaultEpisode field.
func (r *showResolver) DefaultEpisode(ctx context.Context, obj *model.Show) (*model.Episode, error) {
	s, err := r.Loaders.ShowLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return nil, err
	}
	ls := r.FilteredLoaders(ctx)
	eID, err := show.DefaultEpisodeID(ctx, ls, s)
	if err != nil {
		return nil, err
	}
	if eID == nil {
		return nil, merry.New("invalid default episode")
	}
	return r.QueryRoot().Episode(ctx, strconv.Itoa(*eID), nil)
}

// Show returns generated.ShowResolver implementation.
func (r *Resolver) Show() generated.ShowResolver { return &showResolver{r} }

type showResolver struct{ *Resolver }