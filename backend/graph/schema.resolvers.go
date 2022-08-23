package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	merry "github.com/ansel1/merry/v2"
	"github.com/bcc-code/brunstadtv/backend/auth0"
	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/graph/generated"
	gqlmodel "github.com/bcc-code/brunstadtv/backend/graph/model"
	"github.com/bcc-code/brunstadtv/backend/user"
	"github.com/bcc-code/brunstadtv/backend/utils"
)

// Streams is the resolver for the streams field.
func (r *episodeResolver) Streams(ctx context.Context, obj *gqlmodel.Episode) ([]*gqlmodel.Stream, error) {
	intID, _ := strconv.ParseInt(obj.ID, 10, 32)
	streams, err := common.GetFromLoaderForKey(ctx, r.Resolver.Loaders.StreamsLoader, int(intID))
	if err != nil {
		return nil, err
	}

	out := []*gqlmodel.Stream{}
	for _, s := range streams {
		out = append(out, gqlmodel.StreamFrom(ctx, r.Resolver.APIConfig.GetVOD2Domain(), s))
	}

	return out, nil
}

// Files is the resolver for the files field.
func (r *episodeResolver) Files(ctx context.Context, obj *gqlmodel.Episode) ([]*gqlmodel.File, error) {
	intID, err := strconv.ParseInt(obj.ID, 10, 32)
	if err != nil {
		return nil, err
	}
	items, err := common.GetFromLoaderForKey(ctx, r.Resolver.Loaders.FilesLoader, int(intID))
	if err != nil {
		return nil, err
	}
	return utils.MapWithCtx(ctx, items, gqlmodel.FileFrom), nil
}

// Season is the resolver for the season field.
func (r *episodeResolver) Season(ctx context.Context, obj *gqlmodel.Episode) (*gqlmodel.Season, error) {
	if obj.Season != nil {
		return r.QueryRoot().Season(ctx, obj.Season.ID)
	}
	return nil, nil
}

// Show is the resolver for the show field.
func (r *episodeSearchItemResolver) Show(ctx context.Context, obj *gqlmodel.EpisodeSearchItem) (*gqlmodel.Show, error) {
	return r.QueryRoot().Show(ctx, obj.Show.ID)
}

// Season is the resolver for the season field.
func (r *episodeSearchItemResolver) Season(ctx context.Context, obj *gqlmodel.EpisodeSearchItem) (*gqlmodel.Season, error) {
	return r.QueryRoot().Season(ctx, obj.Season.ID)
}

// Page is the resolver for the page field.
func (r *itemSectionResolver) Page(ctx context.Context, obj *gqlmodel.ItemSection) (*gqlmodel.Page, error) {
	return r.QueryRoot().Page(ctx, &obj.Page.ID, nil)
}

// Items is the resolver for the items field.
func (r *itemSectionResolver) Items(ctx context.Context, obj *gqlmodel.ItemSection, first *int, offset *int) (*gqlmodel.CollectionItemPagination, error) {
	items, err := collectionItemResolver(ctx, r.Resolver, obj.ID)
	if err != nil {
		return nil, err
	}
	pagination := utils.Paginate(items, first, offset)
	return &gqlmodel.CollectionItemPagination{
		Total:  pagination.Total,
		First:  pagination.First,
		Offset: pagination.Offset,
		Items:  pagination.Items,
	}, nil
}

// Sections is the resolver for the sections field.
func (r *pageResolver) Sections(ctx context.Context, obj *gqlmodel.Page, first *int, offset *int) (*gqlmodel.SectionPagination, error) {
	sections, err := itemsResolverForIntID(ctx, &itemLoaders[int, common.Section]{
		Item: r.Loaders.SectionLoader,
	}, r.Loaders.SectionsLoader, obj.ID, gqlmodel.SectionFrom)
	if err != nil {
		return nil, err
	}
	pagination := utils.Paginate(sections, first, offset)
	return &gqlmodel.SectionPagination{
		Total:  pagination.Total,
		First:  pagination.First,
		Offset: pagination.Offset,
		Items:  pagination.Items,
	}, nil
}

// Page is the resolver for the page field.
func (r *queryRootResolver) Page(ctx context.Context, id *string, code *string) (*gqlmodel.Page, error) {
	if id != nil {
		return resolverForIntID(ctx, &itemLoaders[int, common.Page]{
			Item:        r.Loaders.PageLoader,
			Permissions: r.Loaders.PagePermissionLoader,
		}, *id, gqlmodel.PageFrom)
	}
	if code != nil {
		intID, err := common.GetFromLoaderByID(ctx, r.Loaders.PageIDFromCodeLoader, *code)
		if err != nil {
			return nil, err
		}
		return resolverFor(ctx, &itemLoaders[int, common.Page]{
			Item:        r.Loaders.PageLoader,
			Permissions: r.Loaders.PagePermissionLoader,
		}, *intID, gqlmodel.PageFrom)
	}
	return nil, merry.Sentinel("Must specify either ID or code", merry.WithHTTPCode(400))
}

// Section is the resolver for the section field.
func (r *queryRootResolver) Section(ctx context.Context, id string) (gqlmodel.Section, error) {
	return resolverForIntID(ctx, &itemLoaders[int, common.Section]{
		Item: r.Loaders.SectionLoader,
	}, id, gqlmodel.SectionFrom)
}

// Show is the resolver for the show field.
func (r *queryRootResolver) Show(ctx context.Context, id string) (*gqlmodel.Show, error) {
	return resolverForIntID(ctx, &itemLoaders[int, common.Show]{
		Item:        r.Loaders.ShowLoader,
		Permissions: r.Loaders.ShowPermissionLoader,
	}, id, gqlmodel.ShowFrom)
}

// Season is the resolver for the season field.
func (r *queryRootResolver) Season(ctx context.Context, id string) (*gqlmodel.Season, error) {
	return resolverForIntID(ctx, &itemLoaders[int, common.Season]{
		Item:        r.Loaders.SeasonLoader,
		Permissions: r.Loaders.SeasonPermissionLoader,
	}, id, gqlmodel.SeasonFrom)
}

// Episode is the resolver for the episode field.
func (r *queryRootResolver) Episode(ctx context.Context, id string) (*gqlmodel.Episode, error) {
	return resolverForIntID(ctx, &itemLoaders[int, common.Episode]{
		Item:        r.Loaders.EpisodeLoader,
		Permissions: r.Loaders.EpisodePermissionLoader,
	}, id, gqlmodel.EpisodeFrom)
}

// Search is the resolver for the search field.
func (r *queryRootResolver) Search(ctx context.Context, queryString string, first *int, offset *int) (*gqlmodel.SearchResult, error) {
	return searchResolver(r, ctx, queryString, first, offset)
}

// Calendar is the resolver for the calendar field.
func (r *queryRootResolver) Calendar(ctx context.Context) (*gqlmodel.Calendar, error) {
	return &gqlmodel.Calendar{}, nil
}

// Event is the resolver for the event field.
func (r *queryRootResolver) Event(ctx context.Context, id string) (*gqlmodel.Event, error) {
	return resolverForIntID(ctx, &itemLoaders[int, common.Event]{
		Item: r.Loaders.EventLoader,
	}, id, gqlmodel.EventFrom)
}

// Faq is the resolver for the faq field.
func (r *queryRootResolver) Faq(ctx context.Context) (*gqlmodel.Faq, error) {
	return &gqlmodel.Faq{}, nil
}

// Me is the resolver for the me field.
func (r *queryRootResolver) Me(ctx context.Context) (*gqlmodel.User, error) {
	gc, err := utils.GinCtx(ctx)
	if err != nil {
		return nil, err
	}

	usr := user.GetFromCtx(gc)

	u := &gqlmodel.User{
		Anonymous: usr.IsAnonymous(),
		BccMember: usr.IsActiveBCC(),
		Roles:     usr.Roles,
	}

	if pid := gc.GetString(auth0.CtxPersonID); pid != "" {
		u.ID = &pid
	}

	if aud := gc.GetString(auth0.CtxJWTAudience); aud != "" {
		u.Audience = &aud
	}

	if usr.Email != "" {
		u.Email = &usr.Email
	}

	return u, nil
}

// Show is the resolver for the show field.
func (r *seasonResolver) Show(ctx context.Context, obj *gqlmodel.Season) (*gqlmodel.Show, error) {
	return r.QueryRoot().Show(ctx, obj.Show.ID)
}

// Episodes is the resolver for the episodes field.
func (r *seasonResolver) Episodes(ctx context.Context, obj *gqlmodel.Season, first *int, offset *int) (*gqlmodel.EpisodePagination, error) {
	items, err := itemsResolverForIntID(ctx, toItemLoaders(r.Loaders.EpisodeLoader, r.Loaders.EpisodePermissionLoader), r.Resolver.Loaders.EpisodesLoader, obj.ID, gqlmodel.EpisodeFrom)
	if err != nil {
		return nil, err
	}

	page := utils.Paginate(items, first, offset)

	return &gqlmodel.EpisodePagination{
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
		Items:  page.Items,
	}, nil
}

// Show is the resolver for the show field.
func (r *seasonSearchItemResolver) Show(ctx context.Context, obj *gqlmodel.SeasonSearchItem) (*gqlmodel.Show, error) {
	return r.QueryRoot().Show(ctx, obj.Show.ID)
}

// Seasons is the resolver for the seasons field.
func (r *showResolver) Seasons(ctx context.Context, obj *gqlmodel.Show, first *int, offset *int) (*gqlmodel.SeasonPagination, error) {
	seasons, err := itemsResolverForIntID(ctx, toItemLoaders(r.Loaders.SeasonLoader, r.Loaders.SeasonPermissionLoader), r.Resolver.Loaders.SeasonsLoader, obj.ID, gqlmodel.SeasonFrom)
	if err != nil {
		return nil, err
	}
	pagination := utils.Paginate(seasons, first, offset)
	return &gqlmodel.SeasonPagination{
		Total:  pagination.Total,
		First:  pagination.First,
		Offset: pagination.Offset,
		Items:  pagination.Items,
	}, nil
}

// Episode returns generated.EpisodeResolver implementation.
func (r *Resolver) Episode() generated.EpisodeResolver { return &episodeResolver{r} }

// EpisodeSearchItem returns generated.EpisodeSearchItemResolver implementation.
func (r *Resolver) EpisodeSearchItem() generated.EpisodeSearchItemResolver {
	return &episodeSearchItemResolver{r}
}

// ItemSection returns generated.ItemSectionResolver implementation.
func (r *Resolver) ItemSection() generated.ItemSectionResolver { return &itemSectionResolver{r} }

// Page returns generated.PageResolver implementation.
func (r *Resolver) Page() generated.PageResolver { return &pageResolver{r} }

// QueryRoot returns generated.QueryRootResolver implementation.
func (r *Resolver) QueryRoot() generated.QueryRootResolver { return &queryRootResolver{r} }

// Season returns generated.SeasonResolver implementation.
func (r *Resolver) Season() generated.SeasonResolver { return &seasonResolver{r} }

// SeasonSearchItem returns generated.SeasonSearchItemResolver implementation.
func (r *Resolver) SeasonSearchItem() generated.SeasonSearchItemResolver {
	return &seasonSearchItemResolver{r}
}

// Show returns generated.ShowResolver implementation.
func (r *Resolver) Show() generated.ShowResolver { return &showResolver{r} }

type episodeResolver struct{ *Resolver }
type episodeSearchItemResolver struct{ *Resolver }
type itemSectionResolver struct{ *Resolver }
type pageResolver struct{ *Resolver }
type queryRootResolver struct{ *Resolver }
type seasonResolver struct{ *Resolver }
type seasonSearchItemResolver struct{ *Resolver }
type showResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryRootResolver) Pages(ctx context.Context, first *int, offset *int) (*gqlmodel.PagePagination, error) {
	panic(nil)
}
