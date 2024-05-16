package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/bcc-code/bcc-media-platform/backend/common"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/generated"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/model"
	"github.com/bcc-code/bcc-media-platform/backend/user"
	"github.com/bcc-code/bcc-media-platform/backend/utils"
	"github.com/samber/lo"
)

// Title is the resolver for the title field.
func (r *contributionTypeResolver) Title(ctx context.Context, obj *model.ContributionType) (string, error) {
	ginCtx, _ := utils.GinCtx(ctx)
	languages := user.GetLanguagesFromCtx(ginCtx)
	phrase, _ := r.Loaders.PhraseLoader.Get(ctx, obj.Code)

	if phrase == nil {
		return obj.Code, nil
	}

	val := phrase.Value.Get(languages)

	return val, nil
}

// Image is the resolver for the image field.
func (r *personResolver) Image(ctx context.Context, obj *model.Person, style *model.ImageStyle) (*string, error) {
	e, err := r.Loaders.PersonLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	return imageOrFallback(ctx, e.Images, style), nil
}

// ContributionTypes is the resolver for the contributionTypes field.
func (r *personResolver) ContributionTypes(ctx context.Context, obj *model.Person) ([]*model.ContributionTypeCount, error) {
	items, err := r.FilteredLoaders(ctx).ContributionsForPersonLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}

	countsByType := make(map[string]int)
	for _, c := range items {
		countsByType[c.Type]++
	}

	var mapped []*model.ContributionTypeCount
	for k, v := range countsByType {
		mapped = append(mapped, &model.ContributionTypeCount{
			Type:  &model.ContributionType{Code: k},
			Count: v,
		})
	}

	return mapped, err
}

// Contributions is the resolver for the contributions field.
func (r *personResolver) Contributions(ctx context.Context, obj *model.Person, first *int, offset *int, types []string, shuffle *bool) (*model.ContributionsPagination, error) {
	items, err := r.FilteredLoaders(ctx).ContributionsForPersonLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}

	if len(types) > 0 {
		items = lo.Filter(items, func(c *common.Contribution, _ int) bool {
			for _, t := range types {
				if c.Type == t {
					return true
				}
			}
			return false
		})
	}

	if shuffle != nil && *shuffle {
		items = lo.Shuffle(items)
	}

	page := utils.Paginate(items, first, offset, nil)
	if err != nil {
		return nil, err
	}

	var result []*model.Contribution
	var wg sync.WaitGroup
	var errors []error
	mu := sync.Mutex{}
	wg.Add(len(page.Items))
	for _, c := range page.Items {
		go func(c *common.Contribution) {
			defer wg.Done()
			contribution, err := resolveContribution(ctx, c, r.Loaders)
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				errors = append(errors, err)
				return
			}
			if contribution == nil {
				return
			}
			result = append(result, contribution)
		}(c)
	}
	wg.Wait()

	for _, e := range errors {
		graphql.AddError(ctx, e)
	}

	return &model.ContributionsPagination{
		Offset: page.Offset,
		First:  page.First,
		Total:  page.Total,
		Items:  result,
	}, nil
}

// ContributionType returns generated.ContributionTypeResolver implementation.
func (r *Resolver) ContributionType() generated.ContributionTypeResolver {
	return &contributionTypeResolver{r}
}

// Person returns generated.PersonResolver implementation.
func (r *Resolver) Person() generated.PersonResolver { return &personResolver{r} }

type contributionTypeResolver struct{ *Resolver }
type personResolver struct{ *Resolver }