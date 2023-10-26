package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"time"

	"github.com/bcc-code/bcc-media-platform/backend/graph/api/generated"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/model"
)

// Global is the resolver for the global field.
func (r *configResolver) Global(ctx context.Context, obj *model.Config, timestamp *string) (*model.GlobalConfig, error) {
	conf, err := withCacheAndTimestamp(ctx, "global_config", r.Queries.GetGlobalConfig, time.Second*30, timestamp)
	if err != nil {
		return nil, err
	}
	return &model.GlobalConfig{
		LiveOnline:  conf.LiveOnline,
		NpawEnabled: conf.NPAWEnabled,
	}, nil
}

// Config returns generated.ConfigResolver implementation.
func (r *Resolver) Config() generated.ConfigResolver { return &configResolver{r} }

type configResolver struct{ *Resolver }
