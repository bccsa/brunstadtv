package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/bcc-code/brunstadtv/backend/graph/api/generated"
	"github.com/bcc-code/brunstadtv/backend/graph/api/model"
	"github.com/bcc-code/brunstadtv/backend/utils"
)

// Items is the resolver for the items field.
func (r *cardListSectionResolver) Items(ctx context.Context, obj *model.CardListSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	return sectionCollectionItemResolver(ctx, r.Resolver, obj.ID, first, offset)
}

// Items is the resolver for the items field.
func (r *cardSectionResolver) Items(ctx context.Context, obj *model.CardSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	return sectionCollectionItemResolver(ctx, r.Resolver, obj.ID, first, offset)
}

// Items is the resolver for the items field.
func (r *defaultGridSectionResolver) Items(ctx context.Context, obj *model.DefaultGridSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	return sectionCollectionItemResolver(ctx, r.Resolver, obj.ID, first, offset)
}

// Items is the resolver for the items field.
func (r *defaultSectionResolver) Items(ctx context.Context, obj *model.DefaultSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	return sectionCollectionItemResolver(ctx, r.Resolver, obj.ID, first, offset)
}

// Items is the resolver for the items field.
func (r *featuredSectionResolver) Items(ctx context.Context, obj *model.FeaturedSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	return sectionCollectionItemResolver(ctx, r.Resolver, obj.ID, first, offset)
}

// Items is the resolver for the items field.
func (r *iconGridSectionResolver) Items(ctx context.Context, obj *model.IconGridSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	return sectionCollectionItemResolver(ctx, r.Resolver, obj.ID, first, offset)
}

// Items is the resolver for the items field.
func (r *iconSectionResolver) Items(ctx context.Context, obj *model.IconSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	return sectionCollectionItemResolver(ctx, r.Resolver, obj.ID, first, offset)
}

// Items is the resolver for the items field.
func (r *labelSectionResolver) Items(ctx context.Context, obj *model.LabelSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	return sectionCollectionItemResolver(ctx, r.Resolver, obj.ID, first, offset)
}

// Items is the resolver for the items field.
func (r *listSectionResolver) Items(ctx context.Context, obj *model.ListSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	return sectionCollectionItemResolver(ctx, r.Resolver, obj.ID, first, offset)
}

// Messages is the resolver for the messages field.
func (r *messageSectionResolver) Messages(ctx context.Context, obj *model.MessageSection) ([]*model.Message, error) {
	s, err := r.Loaders.SectionLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return nil, err
	}
	return resolveMessageSection(ctx, r, s)
}

// Items is the resolver for the items field.
func (r *posterGridSectionResolver) Items(ctx context.Context, obj *model.PosterGridSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	return sectionCollectionItemResolver(ctx, r.Resolver, obj.ID, first, offset)
}

// Items is the resolver for the items field.
func (r *posterSectionResolver) Items(ctx context.Context, obj *model.PosterSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	return sectionCollectionItemResolver(ctx, r.Resolver, obj.ID, first, offset)
}

// Image is the resolver for the image field.
func (r *sectionItemResolver) Image(ctx context.Context, obj *model.SectionItem) (*string, error) {
	fieldCtx := graphql.GetFieldContext(ctx)
	style := model.ImageStyleDefault
	if fieldCtx.Parent != nil && fieldCtx.Parent.Parent != nil && fieldCtx.Parent.Parent.Parent != nil {
		switch fieldCtx.Parent.Parent.Parent.Object {
		case "IconSection", "IconGridSection":
			style = "icon"
		case "PosterSection", "PosterGridSection":
			style = model.ImageStylePoster
		case "FeaturedSection":
			style = model.ImageStyleFeatured
		}
	}
	switch t := obj.Item.(type) {
	case *model.Episode:
		return r.Episode().Image(ctx, t, &style)
	case *model.Season:
		return r.Season().Image(ctx, t, &style)
	case *model.Show:
		return r.Show().Image(ctx, t, &style)
	}
	return obj.Image, nil
}

// CardListSection returns generated.CardListSectionResolver implementation.
func (r *Resolver) CardListSection() generated.CardListSectionResolver {
	return &cardListSectionResolver{r}
}

// CardSection returns generated.CardSectionResolver implementation.
func (r *Resolver) CardSection() generated.CardSectionResolver { return &cardSectionResolver{r} }

// DefaultGridSection returns generated.DefaultGridSectionResolver implementation.
func (r *Resolver) DefaultGridSection() generated.DefaultGridSectionResolver {
	return &defaultGridSectionResolver{r}
}

// DefaultSection returns generated.DefaultSectionResolver implementation.
func (r *Resolver) DefaultSection() generated.DefaultSectionResolver {
	return &defaultSectionResolver{r}
}

// FeaturedSection returns generated.FeaturedSectionResolver implementation.
func (r *Resolver) FeaturedSection() generated.FeaturedSectionResolver {
	return &featuredSectionResolver{r}
}

// IconGridSection returns generated.IconGridSectionResolver implementation.
func (r *Resolver) IconGridSection() generated.IconGridSectionResolver {
	return &iconGridSectionResolver{r}
}

// IconSection returns generated.IconSectionResolver implementation.
func (r *Resolver) IconSection() generated.IconSectionResolver { return &iconSectionResolver{r} }

// LabelSection returns generated.LabelSectionResolver implementation.
func (r *Resolver) LabelSection() generated.LabelSectionResolver { return &labelSectionResolver{r} }

// ListSection returns generated.ListSectionResolver implementation.
func (r *Resolver) ListSection() generated.ListSectionResolver { return &listSectionResolver{r} }

// MessageSection returns generated.MessageSectionResolver implementation.
func (r *Resolver) MessageSection() generated.MessageSectionResolver {
	return &messageSectionResolver{r}
}

// PosterGridSection returns generated.PosterGridSectionResolver implementation.
func (r *Resolver) PosterGridSection() generated.PosterGridSectionResolver {
	return &posterGridSectionResolver{r}
}

// PosterSection returns generated.PosterSectionResolver implementation.
func (r *Resolver) PosterSection() generated.PosterSectionResolver { return &posterSectionResolver{r} }

// SectionItem returns generated.SectionItemResolver implementation.
func (r *Resolver) SectionItem() generated.SectionItemResolver { return &sectionItemResolver{r} }

type cardListSectionResolver struct{ *Resolver }
type cardSectionResolver struct{ *Resolver }
type defaultGridSectionResolver struct{ *Resolver }
type defaultSectionResolver struct{ *Resolver }
type featuredSectionResolver struct{ *Resolver }
type iconGridSectionResolver struct{ *Resolver }
type iconSectionResolver struct{ *Resolver }
type labelSectionResolver struct{ *Resolver }
type listSectionResolver struct{ *Resolver }
type messageSectionResolver struct{ *Resolver }
type posterGridSectionResolver struct{ *Resolver }
type posterSectionResolver struct{ *Resolver }
type sectionItemResolver struct{ *Resolver }