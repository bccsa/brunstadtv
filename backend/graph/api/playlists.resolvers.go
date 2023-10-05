package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/bcc-code/bcc-media-platform/backend/graph/api/generated"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/model"
)

// Image is the resolver for the image field.
func (r *playlistResolver) Image(ctx context.Context, obj *model.Playlist, style *model.ImageStyle) (*string, error) {
	panic(fmt.Errorf("not implemented: Image - image"))
}

// Items is the resolver for the items field.
func (r *playlistResolver) Items(ctx context.Context, obj *model.Playlist, first *int, offset *int) (*model.PlaylistItemPagination, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Playlist returns generated.PlaylistResolver implementation.
func (r *Resolver) Playlist() generated.PlaylistResolver { return &playlistResolver{r} }

type playlistResolver struct{ *Resolver }
