package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"strconv"

	"github.com/bcc-code/bcc-media-platform/backend/common"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/generated"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/model"
	"github.com/bcc-code/bcc-media-platform/backend/user"
	"github.com/bcc-code/bcc-media-platform/backend/utils"
	"github.com/samber/lo"
)

// URL is the resolver for the url field.
func (r *linkResolver) URL(ctx context.Context, obj *model.Link) (string, error) {
	l, err := r.Loaders.LinkLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return "", err
	}
	if !l.ComputedDataGroupID.Valid {
		return l.URL, nil
	}
	ginCtx, _ := utils.GinCtx(ctx)
	u := user.GetFromCtx(ginCtx)
	if u == nil || !u.ActiveBCC {
		return l.URL, nil
	}

	data, err := r.Loaders.ComputedDataLoader.Get(ctx, l.ComputedDataGroupID.UUID)
	if err != nil {
		return "", err
	}
	for _, i := range data {
		if lo.EveryBy(i.Conditions, func(i common.ComputedCondition) bool {
			switch i.Type {
			case "user_church":
				switch i.Operator {
				case "==":
					intValue, _ := strconv.ParseInt(i.Value, 10, 64)
					return lo.Contains(u.ChurchIDs, int(intValue))
				}
			case "user_age":
				intValue, _ := strconv.ParseInt(i.Value, 10, 64)
				switch i.Operator {
				case "<":
					return u.Age < int(intValue)
				case ">":
					return u.Age > int(intValue)
				case "==":
					return u.Age == int(intValue)
				}
			}
			return false
		}) {
			return i.Result, nil
		}
	}
	return l.URL, nil
}

// Image is the resolver for the image field.
func (r *linkResolver) Image(ctx context.Context, obj *model.Link, style *model.ImageStyle) (*string, error) {
	l, err := r.Loaders.LinkLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return nil, err
	}
	if l == nil {
		return nil, nil
	}
	if style == nil {
		s := model.ImageStyleDefault
		style = &s
	}

	ginCtx, _ := utils.GinCtx(ctx)
	languages := user.GetLanguagesFromCtx(ginCtx)

	return l.Images.GetDefault(languages, style.String()), nil
}

// Link returns generated.LinkResolver implementation.
func (r *Resolver) Link() generated.LinkResolver { return &linkResolver{r} }

type linkResolver struct{ *Resolver }
