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
	"github.com/bcc-code/bcc-media-platform/backend/utils"
	"github.com/samber/lo"
)

// Items is the resolver for the items field.
func (r *contextCollectionResolver) Items(ctx context.Context, obj *model.ContextCollection, first *int, offset *int) (*model.SectionItemPagination, error) {
	pagination, err := r.sectionCollectionEntryResolver(ctx, &common.Section{
		Style:        "default",
		CollectionID: utils.AsNullInt(&obj.ID),
	}, first, offset)
	if err != nil {
		return nil, err
	}

	return &model.SectionItemPagination{
		Total:  pagination.Total,
		First:  pagination.First,
		Offset: pagination.Offset,
		Items:  pagination.Items,
	}, nil
}

// Image is the resolver for the image field.
func (r *pageResolver) Image(ctx context.Context, obj *model.Page, style *model.ImageStyle) (*string, error) {
	e, err := r.Loaders.PageLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return nil, err
	}
	return imageOrFallback(ctx, e.Images, style), nil
}

// Sections is the resolver for the sections field.
func (r *pageResolver) Sections(ctx context.Context, obj *model.Page, first *int, offset *int) (*model.SectionPagination, error) {
	intID, err := strconv.ParseInt(obj.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	itemIDs, err := r.FilteredLoaders(ctx).SectionsLoader.Get(ctx, int(intID))
	if err != nil {
		return nil, err
	}

	page := utils.Paginate(itemIDs, first, offset, nil)

	sections, err := r.Loaders.SectionLoader.GetMany(ctx, utils.PointerIntArrayToIntArray(page.Items))
	if err != nil {
		return nil, err
	}
	sections = lo.Filter(sections, func(i *common.Section, _ int) bool {
		return i != nil
	})
	return &model.SectionPagination{
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
		Items:  utils.MapWithCtx(ctx, sections, model.SectionFrom),
	}, nil
}

// ContextCollection returns generated.ContextCollectionResolver implementation.
func (r *Resolver) ContextCollection() generated.ContextCollectionResolver {
	return &contextCollectionResolver{r}
}

// Page returns generated.PageResolver implementation.
func (r *Resolver) Page() generated.PageResolver { return &pageResolver{r} }

type contextCollectionResolver struct{ *Resolver }
type pageResolver struct{ *Resolver }
