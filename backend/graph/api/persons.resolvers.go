package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/bcc-code/bcc-media-platform/backend/common"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/generated"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/model"
	"github.com/bcc-code/bcc-media-platform/backend/user"
	"github.com/bcc-code/bcc-media-platform/backend/utils"
	"github.com/bcc-code/mediabank-bridge/log"
	"github.com/samber/lo"
)

// Title is the resolver for the title field.
func (r *contentTypeResolver) Title(ctx context.Context, obj *model.ContentType) (string, error) {
	ginCtx, _ := utils.GinCtx(ctx)
	languages := user.GetLanguagesFromCtx(ginCtx)
	phrase, _ := r.Loaders.PhraseLoader.Get(ctx, obj.Code)

	if phrase == nil {
		return obj.Code, nil
	}

	val := phrase.Value.Get(languages)

	return val, nil
}

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

// ContributionContentTypes is the resolver for the contributionContentTypes field.
func (r *personResolver) ContributionContentTypes(ctx context.Context, obj *model.Person) ([]*model.ContentTypeCount, error) {
	items, err := r.FilteredLoaders(ctx).ContributionsForPersonLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}

	countsByType := make(map[common.ContentType]int)
	for _, c := range items {
		t := common.ContentTypes.Parse(c.ContentType)
		if t == nil {
			log.L.Warn().Msgf("Unknown content type: %s", c.ContentType)
			continue
		}
		countsByType[*t]++
	}

	var mapped []*model.ContentTypeCount
	for _, t := range common.OrderedContentTypes {
		count, ok := countsByType[t]
		if !ok {
			continue
		}
		mapped = append(mapped, &model.ContentTypeCount{
			Type:  &model.ContentType{Code: t.Value},
			Count: count,
		})
	}

	return mapped, err
}

// Contributions is the resolver for the contributions field.
func (r *personResolver) Contributions(ctx context.Context, obj *model.Person, first *int, offset *int, types []string, contentTypes []string, shuffle *bool) (*model.ContributionsPagination, error) {
	items, err := r.FilteredLoaders(ctx).ContributionsForPersonLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}

	if len(types) > 0 {
		items = lo.Filter(items, func(c *common.Contribution, _ int) bool {
			return lo.Contains(types, c.Type)
		})
	}

	if len(contentTypes) > 0 {
		items = lo.Filter(items, func(c *common.Contribution, _ int) bool {
			return lo.Contains(contentTypes, c.ContentType)
		})
	}

	if shuffle != nil && *shuffle {
		items = lo.Shuffle(items)
	}

	page := utils.Paginate(items, first, offset, nil)

	var wg sync.WaitGroup
	var errors []error
	mu := sync.Mutex{}
	wg.Add(len(page.Items))
	result := make([]*model.Contribution, len(page.Items))
	for i, c := range page.Items {
		i := i
		c := c
		go func() {
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
			result[i] = contribution
		}()
	}
	wg.Wait()

	for _, e := range errors {
		graphql.AddError(ctx, e)
	}

	return &model.ContributionsPagination{
		Offset: page.Offset,
		First:  page.First,
		Total:  page.Total,
		Items: lo.Filter(result, func(c *model.Contribution, _ int) bool {
			return c != nil
		}),
	}, nil
}

// ContentType returns generated.ContentTypeResolver implementation.
func (r *Resolver) ContentType() generated.ContentTypeResolver { return &contentTypeResolver{r} }

// ContributionType returns generated.ContributionTypeResolver implementation.
func (r *Resolver) ContributionType() generated.ContributionTypeResolver {
	return &contributionTypeResolver{r}
}

// Person returns generated.PersonResolver implementation.
func (r *Resolver) Person() generated.PersonResolver { return &personResolver{r} }

type contentTypeResolver struct{ *Resolver }
type contributionTypeResolver struct{ *Resolver }
type personResolver struct{ *Resolver }
