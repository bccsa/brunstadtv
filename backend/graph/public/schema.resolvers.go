package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	merry "github.com/ansel1/merry/v2"
	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/graph/public/generated"
	"github.com/bcc-code/brunstadtv/backend/graph/public/model"
	"github.com/bcc-code/brunstadtv/backend/version"
)

// Episode is the resolver for the episode field.
func (r *queryRootResolver) Episode(ctx context.Context, id string) (*model.Episode, error) {
	intID, _ := strconv.ParseInt(id, 10, 64)
	item, err := r.Loaders.EpisodeLoader.Get(ctx, int(intID))
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, merry.New("item not found", merry.WithUserMessage("item not found"))
	}

	languages := []string{"en"}
	var season *model.Season
	if item.SeasonID.Valid {
		season = &model.Season{
			ID: strconv.Itoa(int(item.SeasonID.Int64)),
		}
	}

	var num *int
	if item.Number.Valid {
		n := int(item.Number.Int64)
		num = &n
	}

	title := item.PublicTitle.String
	if title == "" {
		title = item.Title.Get(languages)
	}

	return &model.Episode{
		ID:     strconv.Itoa(item.ID),
		Index:  !item.PreventPublicIndexing,
		Title:  title,
		Number: num,
		Season: season,
		Image:  item.Images.GetDefault(languages, common.ImageStyleDefault),
	}, nil
}

// Season is the resolver for the season field.
func (r *queryRootResolver) Season(ctx context.Context, id string) (*model.Season, error) {
	intID, _ := strconv.ParseInt(id, 10, 64)
	item, err := r.Loaders.SeasonLoader.Get(ctx, int(intID))
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, merry.New("item not found", merry.WithUserMessage("item not found"))
	}

	languages := []string{"en"}

	title := item.PublicTitle.String
	if title == "" {
		title = item.Title.Get(languages)
	}

	return &model.Season{
		ID:     strconv.Itoa(item.ID),
		Title:  title,
		Number: item.Number,
		Image:  item.Images.GetDefault(languages, common.ImageStyleDefault),
		Show: &model.Show{
			ID: strconv.Itoa(item.ShowID),
		},
	}, nil
}

// Show is the resolver for the show field.
func (r *queryRootResolver) Show(ctx context.Context, id string) (*model.Show, error) {
	intID, _ := strconv.ParseInt(id, 10, 64)
	item, err := r.Loaders.ShowLoader.Get(ctx, int(intID))
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, merry.New("item not found", merry.WithUserMessage("item not found"))
	}

	languages := []string{"en"}

	title := item.PublicTitle.String
	if title == "" {
		title = item.Title.Get(languages)
	}

	return &model.Show{
		ID:    strconv.Itoa(item.ID),
		Title: title,
		Image: item.Images.GetDefault(languages, common.ImageStyleDefault),
	}, nil
}

// Version is the resolver for the version field.
func (r *queryRootResolver) Version(ctx context.Context) (*model.Version, error) {
	return version.GQLHandler()
}

// QueryRoot returns generated.QueryRootResolver implementation.
func (r *Resolver) QueryRoot() generated.QueryRootResolver { return &queryRootResolver{r} }

type queryRootResolver struct{ *Resolver }
