package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	cache "github.com/Code-Hex/go-generics-cache"
	merry "github.com/ansel1/merry/v2"
	"github.com/bcc-code/bcc-media-platform/backend/applications"
	"github.com/bcc-code/bcc-media-platform/backend/auth0"
	"github.com/bcc-code/bcc-media-platform/backend/common"
	"github.com/bcc-code/bcc-media-platform/backend/export"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/generated"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/model"
	"github.com/bcc-code/bcc-media-platform/backend/memorycache"
	"github.com/bcc-code/bcc-media-platform/backend/ratelimit"
	"github.com/bcc-code/bcc-media-platform/backend/sqlc"
	"github.com/bcc-code/bcc-media-platform/backend/user"
	"github.com/bcc-code/bcc-media-platform/backend/utils"
	"github.com/bcc-code/mediabank-bridge/log"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/samber/lo"
	null "gopkg.in/guregu/null.v4"
)

// Application is the resolver for the application field.
func (r *queryRootResolver) Application(ctx context.Context, timestamp *string) (*model.Application, error) {
	ginCtx, err := utils.GinCtx(ctx)
	if err != nil {
		return nil, err
	}
	ctxApp, err := applications.GetFromCtx(ginCtx)
	if err != nil {
		return nil, err
	}

	if timestamp != nil {
		withTimestampExpiration(ctx, "application:"+strconv.Itoa(ctxApp.ID), timestamp, func() {
			r.Loaders.ApplicationLoader.Clear(ctx, ctxApp.ID)
		})
	}

	app, err := r.Loaders.ApplicationLoader.Get(ctx, ctxApp.ID)
	if err != nil {
		return nil, err
	}

	u := user.GetFromCtx(ginCtx)
	livestreamEnabled := len(app.LivestreamRoles) == 0 || len(lo.Intersect(app.LivestreamRoles, u.Roles)) > 0

	var page *model.Page
	if app.DefaultPageID.Valid {
		page = &model.Page{
			ID: strconv.Itoa(int(app.DefaultPageID.Int64)),
		}
	}
	var searchPage *model.Page
	if app.SearchPageID.Valid {
		searchPage = &model.Page{
			ID: strconv.Itoa(int(app.SearchPageID.Int64)),
		}
	}
	var gamesPage *model.Page
	if app.GamesPageID.Valid {
		gamesPage = &model.Page{
			ID: strconv.Itoa(int(app.GamesPageID.Int64)),
		}
	}

	return &model.Application{
		ID:                strconv.Itoa(app.ID),
		Code:              app.Code,
		Page:              page,
		SearchPage:        searchPage,
		GamesPage:         gamesPage,
		ClientVersion:     app.ClientVersion,
		LivestreamEnabled: livestreamEnabled,
	}, nil
}

// Languages is the resolver for the languages field.
func (r *queryRootResolver) Languages(ctx context.Context) ([]string, error) {
	languages, err := memorycache.GetOrSet(ctx, "languages", r.Queries.GetLanguageKeys, cache.WithExpiration(time.Minute*5))
	if err != nil {
		return nil, err
	}
	return languages, nil
}

// Export is the resolver for the export field.
func (r *queryRootResolver) Export(ctx context.Context, groups []string) (*model.Export, error) {
	ginCtx, err := utils.GinCtx(ctx)
	if err != nil {
		return nil, err
	}

	profile := user.GetProfileFromCtx(ginCtx)
	if profile == nil {
		return nil, merry.New(
			"Not authorized",
			merry.WithUserMessage("you are not authorized for this query"),
		)
	}

	url, err := export.DoExport(ctx, r, r.AWSConfig.GetTempStorageBucket())
	if err != nil {
		return nil, err
	}

	return &model.Export{
		URL:       url,
		DbVersion: export.SQLiteExportDBVersion,
	}, nil
}

// Redirect is the resolver for the redirect field.
func (r *queryRootResolver) Redirect(ctx context.Context, id string) (*model.RedirectLink, error) {
	ginCtx, err := utils.GinCtx(ctx)
	if err != nil {
		return nil, err
	}

	profile := user.GetProfileFromCtx(ginCtx)
	if profile == nil {
		return nil, merry.New(
			"Not authorized",
			merry.WithUserMessage("you are not authorized for this query"),
		)
	}

	redirID, err := r.Loaders.RedirectIDFromCodeLoader.Get(ctx, id)
	if err != nil {
		return nil, merry.Wrap(err, merry.WithUserMessage("Failed to retrieve data"))
	}

	if redirID == nil {
		return nil, merry.New("no rows", merry.WithUserMessage("Code not found"))
	}

	redir, err := r.Loaders.RedirectLoader.Get(ctx, *redirID)
	if err != nil {
		return nil, merry.Wrap(err, merry.WithUserMessage("Failed to retrieve data"))
	}

	usr := user.GetFromCtx(ginCtx)

	// Add JWT to url
	url, err := url.Parse(redir.TargetURL)
	if err != nil {
		return nil, merry.Wrap(err, merry.WithUserMessage("Internal server error. URL-PARSE"))
	}

	if redir.IncludeToken {
		// Build a JWT!
		tok, err := jwt.NewBuilder().
			Claim("person_id", profile.UserID).
			Claim("gender", usr.Gender).
			Claim("first_name", usr.FirstName).
			Issuer("https://api.brunstad.tv/").
			IssuedAt(time.Now()).
			Expiration(time.Now().Add(30 * time.Second)).
			Build()
		if err != nil {
			return nil, merry.Wrap(err, merry.WithUserMessage("Internal server error. TOKEN-ERROR"))
		}

		// Sign a JWT!
		signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, r.RedirectConfig.GetPrivateKey()))
		if err != nil {
			return nil, merry.Wrap(err, merry.WithUserMessage("Internal server error. TOKEN-SIGN-ERROR"))
		}
		q := url.Query()
		q.Add("token", string(signed))
		url.RawQuery = q.Encode()
	}

	return &model.RedirectLink{
		URL:                    url.String(),
		RequiresAuthentication: redir.RequiresAuthentication,
	}, nil
}

// Page is the resolver for the page field.
func (r *queryRootResolver) Page(ctx context.Context, id *string, code *string) (*model.Page, error) {
	if id != nil {
		return resolverForIntID(ctx, &itemLoaders[int, common.Page]{
			Item:        r.Loaders.PageLoader,
			Permissions: r.Loaders.PagePermissionLoader,
		}, *id, model.PageFrom)
	}
	if code != nil {
		intID, err := r.Loaders.PageIDFromCodeLoader.Get(ctx, *code)
		if err != nil {
			return nil, err
		}
		if intID == nil {
			return nil, merry.Sentinel("No page found with that code")
		}
		return resolverFor(ctx, &itemLoaders[int, common.Page]{
			Item:        r.Loaders.PageLoader,
			Permissions: r.Loaders.PagePermissionLoader,
		}, *intID, model.PageFrom)
	}
	return nil, merry.Sentinel("Must specify either ID or code", merry.WithHTTPCode(400))
}

// Section is the resolver for the section field.
func (r *queryRootResolver) Section(ctx context.Context, id string, timestamp *string) (model.Section, error) {
	if timestamp != nil {
		intID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, err
		}
		ginCtx, _ := utils.GinCtx(ctx)
		ginCtx.Set(timestampContextKey, *timestamp)
		withTimestampExpiration(ctx, "section:"+id, timestamp, func() {
			r.Loaders.SectionLoader.Clear(ctx, int(intID))
		})
	}
	return resolverForIntID(ctx, &itemLoaders[int, common.Section]{
		Item:        r.Loaders.SectionLoader,
		Permissions: r.Loaders.SectionPermissionLoader,
	}, id, model.SectionFrom)
}

// Show is the resolver for the show field.
func (r *queryRootResolver) Show(ctx context.Context, id string) (*model.Show, error) {
	return resolverForIntID(ctx, &itemLoaders[int, common.Show]{
		Item:        r.Loaders.ShowLoader,
		Permissions: r.Loaders.ShowPermissionLoader,
	}, id, model.ShowFrom)
}

// Season is the resolver for the season field.
func (r *queryRootResolver) Season(ctx context.Context, id string) (*model.Season, error) {
	return resolverForIntID(ctx, &itemLoaders[int, common.Season]{
		Item:        r.Loaders.SeasonLoader,
		Permissions: r.Loaders.SeasonPermissionLoader,
	}, id, model.SeasonFrom)
}

// Episode is the resolver for the episode field.
func (r *queryRootResolver) Episode(ctx context.Context, id string, context *model.EpisodeContext) (*model.Episode, error) {
	ginCtx, _ := utils.GinCtx(ctx)
	if context != nil {
		var collectionID null.Int
		if context.PlaylistID != nil {
			playlist, err := r.Loaders.PlaylistLoader.Get(ctx, utils.AsUuid(*context.PlaylistID))
			if err == nil && playlist != nil {
				collectionID = playlist.CollectionID
			}
		}
		if !collectionID.Valid {
			collectionID = utils.AsNullInt(context.CollectionID)
		}
		ginCtx.Set(episodeContextKey, common.EpisodeContext{
			CollectionID: collectionID,
			Cursor:       null.StringFromPtr(context.Cursor),
			Shuffle:      null.BoolFromPtr(context.Shuffle),
		})
	}
	episodeID, err := r.episodeIDResolver(ctx, id)
	if err != nil {
		return nil, err
	}
	return resolverForIntID(ctx, &itemLoaders[int, common.Episode]{
		Item:        r.Loaders.EpisodeLoader,
		Permissions: r.Loaders.EpisodePermissionLoader,
	}, episodeID, model.EpisodeFrom)
}

// Episodes is the resolver for the episodes field.
func (r *queryRootResolver) Episodes(ctx context.Context, ids []string) ([]*model.Episode, error) {

	resolved := make([]*model.Episode, len(ids))
	ch := make(chan *model.Episode, len(ids))
	errCh := make(chan error, len(ids))
	defer close(ch)
	defer close(errCh)

	for _, id := range ids {
		go func(id string) {
			episodeID, err := r.episodeIDResolver(ctx, id)
			if err != nil {
				errCh <- err
				return
			}
			episode, err := resolverForIntID(ctx, &itemLoaders[int, common.Episode]{
				Item:        r.Loaders.EpisodeLoader,
				Permissions: r.Loaders.EpisodePermissionLoader,
			}, episodeID, model.EpisodeFrom)
			if err != nil {
				errCh <- err
				return
			}
			ch <- episode
		}(id)
	}

	for i := 0; i < len(ids); i++ {
		select {
		case episode := <-ch:
			resolved[i] = episode
		case err := <-errCh:
			return nil, err
		}
	}

	return resolved, nil
}

// Playlist is the resolver for the playlist field.
func (r *queryRootResolver) Playlist(ctx context.Context, id string) (*model.Playlist, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	// ignoring permissions as items in the playlist is also checked for permissions,@
	// and it won't show up in lists if the user does not have access
	return resolverFor(ctx, &itemLoaders[uuid.UUID, common.Playlist]{
		Item: r.Loaders.PlaylistLoader,
	}, uid, model.PlaylistFrom)
}

// Search is the resolver for the search field.
func (r *queryRootResolver) Search(ctx context.Context, queryString string, first *int, offset *int, typeArg *string, minScore *int) (*model.SearchResult, error) {
	return searchResolver(r, ctx, queryString, first, offset, typeArg, minScore)
}

// Game is the resolver for the game field.
func (r *queryRootResolver) Game(ctx context.Context, id string) (*model.Game, error) {
	return uuidItemLoader(ctx, r.Loaders.GameLoader, model.GameFrom, id)
}

// Short is the resolver for the short field.
func (r *queryRootResolver) Short(ctx context.Context, id string) (*model.Short, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	shortIDSegments, err := r.GetFilteredLoaders(ctx).ShortIDsLoader(ctx)
	if err != nil {
		return nil, err
	}
	shortIDs := lo.Flatten(shortIDSegments)
	if !lo.Contains(shortIDs, uid) {
		return nil, ErrItemNotFound
	}
	short, err := r.Loaders.ShortLoader.Get(ctx, uid)
	if err != nil {
		return nil, err
	}
	return shortToShort(ctx, short), nil
}

// Shorts is the resolver for the shorts field.
func (r *queryRootResolver) Shorts(ctx context.Context, cursor *string, limit *int, initialShortID *string) (*model.ShortsPagination, error) {
	return r.getShorts(ctx, cursor, limit, initialShortID)
}

// PendingAchievements is the resolver for the pendingAchievements field.
func (r *queryRootResolver) PendingAchievements(ctx context.Context) ([]*model.Achievement, error) {
	p, err := getProfile(ctx)
	if err != nil {
		return nil, err
	}
	ids, err := r.Loaders.UnconfirmedAchievementsLoader.Get(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	items, err := r.Loaders.AchievementLoader.GetMany(ctx, utils.PointerArrayToArray(ids))
	if err != nil {
		return nil, err
	}
	return utils.MapWithCtx(ctx, items, model.AchievementFrom), nil
}

// Achievement is the resolver for the achievement field.
func (r *queryRootResolver) Achievement(ctx context.Context, id string) (*model.Achievement, error) {
	return uuidItemLoader(ctx, r.Loaders.AchievementLoader, model.AchievementFrom, id)
}

// AchievementGroup is the resolver for the achievementGroup field.
func (r *queryRootResolver) AchievementGroup(ctx context.Context, id string) (*model.AchievementGroup, error) {
	return uuidItemLoader(ctx, r.Loaders.AchievementGroupLoader, model.AchievementGroupFrom, id)
}

// AchievementGroups is the resolver for the achievementGroups field.
func (r *queryRootResolver) AchievementGroups(ctx context.Context, first *int, offset *int) (*model.AchievementGroupPagination, error) {
	ids, err := memorycache.GetOrSet(ctx, "achievement_groups", r.Queries.ListAchievementGroups)
	if err != nil {
		return nil, err
	}
	if ids == nil {
		ids = []uuid.UUID{}
	}
	page := utils.Paginate(ids, first, offset, nil)

	groups, err := r.Loaders.AchievementGroupLoader.GetMany(ctx, page.Items)
	if err != nil {
		return nil, err
	}

	return &model.AchievementGroupPagination{
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
		Items:  utils.MapWithCtx(ctx, groups, model.AchievementGroupFrom),
	}, nil
}

// StudyTopic is the resolver for the studyTopic field.
func (r *queryRootResolver) StudyTopic(ctx context.Context, id string) (*model.StudyTopic, error) {
	_, err := getProfile(ctx)
	if err != nil {
		return nil, err
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	ginCtx, _ := utils.GinCtx(ctx)
	languages := user.GetLanguagesFromCtx(ginCtx)
	return resolverFor(ctx, &itemLoaders[uuid.UUID, common.StudyTopic]{
		Item: r.Loaders.StudyTopicLoader,
	}, uid, func(ctx context.Context, topic *common.StudyTopic) *model.StudyTopic {
		return &model.StudyTopic{
			ID:    topic.ID.String(),
			Title: topic.Title.Get(languages),
		}
	})
}

// StudyLesson is the resolver for the studyLesson field.
func (r *queryRootResolver) StudyLesson(ctx context.Context, id string) (*model.Lesson, error) {
	_, err := getProfile(ctx)
	if err != nil {
		return nil, err
	}
	return uuidItemLoader(ctx, r.Loaders.StudyLessonLoader, model.LessonFrom, id)
}

// Calendar is the resolver for the calendar field.
func (r *queryRootResolver) Calendar(ctx context.Context) (*model.Calendar, error) {
	return &model.Calendar{}, nil
}

// Event is the resolver for the event field.
func (r *queryRootResolver) Event(ctx context.Context, id string) (*model.Event, error) {
	return resolverForIntID(ctx, &itemLoaders[int, common.Event]{
		Item: r.Loaders.EventLoader,
	}, id, model.EventFrom)
}

// Faq is the resolver for the faq field.
func (r *queryRootResolver) Faq(ctx context.Context) (*model.Faq, error) {
	return &model.Faq{}, nil
}

// Me is the resolver for the me field.
func (r *queryRootResolver) Me(ctx context.Context) (*model.User, error) {
	gc, err := utils.GinCtx(ctx)
	if err != nil {
		return nil, err
	}

	usr := user.GetFromCtx(gc)

	roles := user.GetRolesFromCtx(gc)

	u := &model.User{
		Anonymous:   usr.IsAnonymous(),
		BccMember:   usr.IsActiveBCC(),
		Roles:       roles,
		DisplayName: usr.DisplayName,
		FirstName:   usr.FirstName,
		Analytics:   &model.Analytics{},
	}

	if pid := gc.GetString(auth0.CtxUserID); pid != "" {
		u.ID = &pid
	}

	switch usr.Gender {
	case "male":
		u.Gender = model.GenderMale
	case "female":
		u.Gender = model.GenderFemale
	default:
		u.Gender = model.GenderUnknown
	}

	//if aud := gc.GetString(auth0.CtxAudience); aud != "" {
	//	u.Audience = &aud
	//}

	if usr.Email != "" {
		u.Email = &usr.Email
		u.EmailVerified = usr.EmailVerified
	}

	u.CompletedRegistration = usr.IsActiveBCC() || usr.CompletedRegistration

	return u, nil
}

// MyList is the resolver for the myList field.
func (r *queryRootResolver) MyList(ctx context.Context) (*model.UserCollection, error) {
	p, err := getProfile(ctx)
	if err != nil {
		return nil, err
	}
	ginCtx, err := utils.GinCtx(ctx)
	if err != nil {
		return nil, err
	}
	app, err := applications.GetFromCtx(ginCtx)
	if err != nil {
		return nil, err
	}
	l := r.GetLoaders().ProfileMyListCollectionID
	id, err := l.Get(ctx, p.ID)
	if id == nil {
		uc := common.UserCollection{
			ID:                 uuid.New(),
			ApplicationGroupID: app.GroupID,
			Title:              "my-list",
			Metadata: common.UserCollectionMetadata{
				MyList: true,
			},
			ProfileID: p.ID,
		}
		id = &uc.ID
		l.Clear(ctx, p.ID)
		l.Prime(ctx, p.ID, id)
		err = r.Queries.UpsertUserCollection(ctx, sqlc.UpsertUserCollectionParams{
			ID:                 uc.ID,
			ApplicationgroupID: uc.ApplicationGroupID,
			MyList:             uc.Metadata.MyList,
			ProfileID:          uc.ProfileID,
			Title:              uc.Title,
		})
		if err != nil {
			return nil, err
		}
	}
	return r.QueryRoot().UserCollection(ctx, id.String())
}

// UserCollection is the resolver for the userCollection field.
func (r *queryRootResolver) UserCollection(ctx context.Context, id string) (*model.UserCollection, error) {
	p, err := getProfile(ctx)
	if err != nil {
		return nil, err
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, common.ErrInvalidUUID
	}
	col, err := r.Loaders.UserCollectionLoader.Get(ctx, uid)
	if err != nil {
		return nil, err
	}
	if col == nil {
		return nil, common.ErrItemNotFound
	}
	if col.ProfileID != p.ID {
		return nil, common.ErrItemNoAccess
	}

	return &model.UserCollection{
		ID:    col.ID.String(),
		Title: col.Title,
	}, nil
}

// Config is the resolver for the config field.
func (r *queryRootResolver) Config(ctx context.Context) (*model.Config, error) {
	return &model.Config{}, nil
}

// Profiles is the resolver for the profiles field.
func (r *queryRootResolver) Profiles(ctx context.Context) ([]*model.Profile, error) {
	ginCtx, err := utils.GinCtx(ctx)
	if err != nil {
		return nil, err
	}
	profiles := user.GetProfilesFromCtx(ginCtx)

	return lo.Map(profiles, func(i common.Profile, _ int) *model.Profile {
		return &model.Profile{
			ID:   i.ID.String(),
			Name: i.Name,
		}
	}), nil
}

// Profile is the resolver for the profile field.
func (r *queryRootResolver) Profile(ctx context.Context) (*model.Profile, error) {
	ginCtx, err := utils.GinCtx(ctx)
	if err != nil {
		return nil, err
	}
	profile := user.GetProfileFromCtx(ginCtx)

	return &model.Profile{
		ID:   profile.ID.String(),
		Name: profile.Name,
	}, nil
}

// LegacyIDLookup is the resolver for the legacyIDLookup field.
func (r *queryRootResolver) LegacyIDLookup(ctx context.Context, options *model.LegacyIDLookupOptions) (*model.LegacyIDLookup, error) {
	var id *int
	var err error
	if options.EpisodeID != nil {
		id, err = r.Loaders.EpisodeIDFromLegacyIDLoader.Get(ctx, *options.EpisodeID)
	}
	if options.ProgramID != nil {
		id, err = r.Loaders.EpisodeIDFromLegacyProgramIDLoader.Get(ctx, *options.ProgramID)
	}
	if err != nil {
		return nil, err
	}
	if id == nil {
		return nil, ErrItemNotFound
	}
	return &model.LegacyIDLookup{
		ID: strconv.Itoa(*id),
	}, nil
}

// Prompts is the resolver for the prompts field.
func (r *queryRootResolver) Prompts(ctx context.Context, timestamp *string) ([]model.Prompt, error) {
	loaders := r.FilteredLoaders(ctx)
	if timestamp != nil {
		withTimestampExpiration(ctx, "prompts:"+loaders.Key, timestamp, func() {
			ids, err := loaders.PromptIDsLoader(ctx)
			if err != nil {
				log.L.Error().Err(err).Send()
			}
			for _, id := range ids {
				r.Loaders.PromptLoader.Clear(ctx, id)
			}
			memorycache.Delete(fmt.Sprintf("promptIDs:roles:%s", loaders.Key))
		})
	}
	ids, err := loaders.PromptIDsLoader(ctx)
	if err != nil {
		return nil, err
	}
	surveys, err := r.Loaders.PromptLoader.GetMany(ctx, ids)
	if err != nil {
		return nil, err
	}
	return utils.MapWithCtx(ctx, surveys, model.PromptFrom), nil
}

// Subscriptions is the resolver for the subscriptions field.
func (r *queryRootResolver) Subscriptions(ctx context.Context) ([]model.SubscriptionTopic, error) {
	p, err := getProfile(ctx)
	if err != nil {
		return nil, err
	}
	err = ratelimit.Endpoint(ctx, "subscriptions", 10, false)
	if err != nil {
		return nil, err
	}
	subscriptions, err := r.GetQueries().ListSubscriptions(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	return lo.Filter(lo.Map(subscriptions, func(i string, _ int) model.SubscriptionTopic {
		return model.SubscriptionTopic(i)
	}), func(i model.SubscriptionTopic, _ int) bool {
		return i.IsValid()
	}), nil
}

// QueryRoot returns generated.QueryRootResolver implementation.
func (r *Resolver) QueryRoot() generated.QueryRootResolver { return &queryRootResolver{r} }

type queryRootResolver struct{ *Resolver }
