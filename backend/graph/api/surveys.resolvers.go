package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bcc-code/brunstadtv/backend/graph/api/generated"
	"github.com/bcc-code/brunstadtv/backend/graph/api/model"
	"github.com/bcc-code/brunstadtv/backend/utils"
)

// Questions is the resolver for the questions field.
func (r *surveyResolver) Questions(ctx context.Context, obj *model.Survey, first *int, offset *int) (*model.SurveyQuestionPagination, error) {
	items, err := r.GetFilteredLoaders(ctx).SurveyQuestionsLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	page := utils.Paginate(items, first, offset, nil)
	questions, err := r.GetLoaders().SurveyQuestionLoader.GetMany(ctx, utils.PointerArrayToArray(page.Items))
	if err != nil {
		return nil, err
	}
	return &model.SurveyQuestionPagination{
		First:  page.First,
		Offset: page.Offset,
		Total:  page.Total,
		Items:  utils.MapWithCtx(ctx, questions, model.SurveyQuestionFrom),
	}, nil
}

// Survey returns generated.SurveyResolver implementation.
func (r *Resolver) Survey() generated.SurveyResolver { return &surveyResolver{r} }

type surveyResolver struct{ *Resolver }