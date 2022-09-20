package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bcc-code/brunstadtv/backend/graph/generated"
	gqlmodel "github.com/bcc-code/brunstadtv/backend/graph/model"
)

// Show is the resolver for the show field.
func (r *episodeSearchItemResolver) Show(ctx context.Context, obj *gqlmodel.EpisodeSearchItem) (*gqlmodel.Show, error) {
	return r.QueryRoot().Show(ctx, obj.Show.ID)
}

// Season is the resolver for the season field.
func (r *episodeSearchItemResolver) Season(ctx context.Context, obj *gqlmodel.EpisodeSearchItem) (*gqlmodel.Season, error) {
	return r.QueryRoot().Season(ctx, obj.Season.ID)
}

// Show is the resolver for the show field.
func (r *seasonSearchItemResolver) Show(ctx context.Context, obj *gqlmodel.SeasonSearchItem) (*gqlmodel.Show, error) {
	return r.QueryRoot().Show(ctx, obj.Show.ID)
}

// EpisodeSearchItem returns generated.EpisodeSearchItemResolver implementation.
func (r *Resolver) EpisodeSearchItem() generated.EpisodeSearchItemResolver {
	return &episodeSearchItemResolver{r}
}

// SeasonSearchItem returns generated.SeasonSearchItemResolver implementation.
func (r *Resolver) SeasonSearchItem() generated.SeasonSearchItemResolver {
	return &seasonSearchItemResolver{r}
}

type episodeSearchItemResolver struct{ *Resolver }
type seasonSearchItemResolver struct{ *Resolver }
