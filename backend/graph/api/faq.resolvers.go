package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"

	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/graph/api/generated"
	"github.com/bcc-code/brunstadtv/backend/graph/api/model"
	"github.com/bcc-code/brunstadtv/backend/utils"
	"github.com/google/uuid"
)

// Categories is the resolver for the categories field.
func (r *fAQResolver) Categories(ctx context.Context, obj *model.Faq, first *int, offset *int) (*model.FAQCategoryPagination, error) {
	ids, err := r.GetFilteredLoaders(ctx).FAQCategoryIDsLoader(ctx)
	if err != nil {
		return nil, err
	}
	page := utils.Paginate(ids, first, offset, nil)

	items, err := r.GetLoaders().FAQCategoryLoader.GetMany(ctx, page.Items)
	if err != nil {
		return nil, err
	}

	return &model.FAQCategoryPagination{
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
		Items:  utils.MapWithCtx(ctx, items, model.FAQCategoryFrom),
	}, nil
}

// Category is the resolver for the category field.
func (r *fAQResolver) Category(ctx context.Context, obj *model.Faq, id string) (*model.FAQCategory, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return resolverFor(ctx, &itemLoaders[uuid.UUID, common.FAQCategory]{
		Item: r.Loaders.FAQCategoryLoader,
	}, uid, model.FAQCategoryFrom)
}

// Question is the resolver for the question field.
func (r *fAQResolver) Question(ctx context.Context, obj *model.Faq, id string) (*model.Question, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return resolverFor(ctx, &itemLoaders[uuid.UUID, common.Question]{
		Item: r.Loaders.QuestionLoader,
	}, uid, model.QuestionFrom)
}

// Questions is the resolver for the questions field.
func (r *fAQCategoryResolver) Questions(ctx context.Context, obj *model.FAQCategory, first *int, offset *int) (*model.QuestionPagination, error) {
	itemIDs, err := r.GetFilteredLoaders(ctx).FAQQuestionsLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}

	page := utils.Paginate(itemIDs, first, offset, nil)

	items, err := r.Loaders.QuestionLoader.GetMany(ctx, utils.PointerArrayToArray(page.Items))
	if err != nil {
		return nil, err
	}

	return &model.QuestionPagination{
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
		Items:  utils.MapWithCtx(ctx, items, model.QuestionFrom),
	}, nil
}

// Category is the resolver for the category field.
func (r *questionResolver) Category(ctx context.Context, obj *model.Question) (*model.FAQCategory, error) {
	return r.FAQ().Category(ctx, nil, obj.Category.ID)
}

// FAQ returns generated.FAQResolver implementation.
func (r *Resolver) FAQ() generated.FAQResolver { return &fAQResolver{r} }

// FAQCategory returns generated.FAQCategoryResolver implementation.
func (r *Resolver) FAQCategory() generated.FAQCategoryResolver { return &fAQCategoryResolver{r} }

// Question returns generated.QuestionResolver implementation.
func (r *Resolver) Question() generated.QuestionResolver { return &questionResolver{r} }

type fAQResolver struct{ *Resolver }
type fAQCategoryResolver struct{ *Resolver }
type questionResolver struct{ *Resolver }
