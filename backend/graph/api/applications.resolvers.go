package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"

	"github.com/bcc-code/bcc-media-platform/backend/graph/api/generated"
	gqlmodel "github.com/bcc-code/bcc-media-platform/backend/graph/api/model"
	"github.com/bcc-code/bcc-media-platform/backend/utils"
)

// Page is the resolver for the page field.
func (r *applicationResolver) Page(ctx context.Context, obj *gqlmodel.Application) (*gqlmodel.Page, error) {
	featureFlags := utils.GetFeatureFlags(ctx)
	if f, ok := featureFlags.GetVariant("application-page"); ok && f != "" {
		page, err := r.QueryRoot().Page(ctx, nil, &f)
		if err == nil {
			return page, nil
		}
	}
	if obj.Page != nil {
		return r.QueryRoot().Page(ctx, &obj.Page.ID, nil)
	}
	return nil, nil
}

// SearchPage is the resolver for the searchPage field.
func (r *applicationResolver) SearchPage(ctx context.Context, obj *gqlmodel.Application) (*gqlmodel.Page, error) {
	if obj.SearchPage != nil {
		return r.QueryRoot().Page(ctx, &obj.SearchPage.ID, nil)
	}
	return nil, nil
}

// GamesPage is the resolver for the gamesPage field.
func (r *applicationResolver) GamesPage(ctx context.Context, obj *gqlmodel.Application) (*gqlmodel.Page, error) {
	if obj.GamesPage != nil {
		return r.QueryRoot().Page(ctx, &obj.GamesPage.ID, nil)
	}
	return nil, nil
}

// Application returns generated.ApplicationResolver implementation.
func (r *Resolver) Application() generated.ApplicationResolver { return &applicationResolver{r} }

type applicationResolver struct{ *Resolver }
