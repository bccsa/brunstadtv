package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bcc-code/bcc-media-platform/backend/applications"
	"github.com/bcc-code/bcc-media-platform/backend/common"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/generated"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/model"
	"github.com/bcc-code/bcc-media-platform/backend/user"
	"github.com/bcc-code/bcc-media-platform/backend/utils"
	"github.com/samber/lo"
	null "gopkg.in/guregu/null.v4"
)

// Locked is the resolver for the locked field.
func (r *episodeResolver) Locked(ctx context.Context, obj *model.Episode) (bool, error) {
	e, err := r.Loaders.EpisodeLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return false, err
	}
	perms, err := r.Loaders.EpisodePermissionLoader.Get(ctx, e.ID)
	if err != nil {
		return false, err
	}
	ginCtx, _ := utils.GinCtx(ctx)
	roles := user.GetRolesFromCtx(ginCtx)
	if e.PublishDate.After(time.Now()) && len(lo.Intersect(perms.Roles.EarlyAccess, roles)) == 0 {
		return true, nil
	}
	return false, nil
}

// AvailableFrom is the resolver for the availableFrom field.
func (r *episodeResolver) AvailableFrom(ctx context.Context, obj *model.Episode) (string, error) {
	perms, err := r.Loaders.EpisodePermissionLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return "", err
	}
	ginCtx, err := utils.GinCtx(ctx)
	if err != nil {
		return "", err
	}
	roles := user.GetRolesFromCtx(ginCtx)
	if len(lo.Intersect(roles, perms.Roles.EarlyAccess)) == 0 {
		return perms.Availability.From.Format(time.RFC3339), nil
	}
	return "1800-01-01T00:00:00Z", nil
}

// OriginalTitle is the resolver for the originalTitle field.
func (r *episodeResolver) OriginalTitle(ctx context.Context, obj *model.Episode) (string, error) {
	return r.getTitleFromContext(ctx, obj, &[]string{"no", "en"})
}

// Title is the resolver for the title field.
func (r *episodeResolver) Title(ctx context.Context, obj *model.Episode) (string, error) {
	return r.getTitleFromContext(ctx, obj, nil)
}

// Image is the resolver for the image field.
func (r *episodeResolver) Image(ctx context.Context, obj *model.Episode, style *model.ImageStyle) (*string, error) {
	e, err := r.Loaders.EpisodeLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return nil, err
	}
	var fallbacks []common.Images
	if obj.Season != nil {
		s, err := r.Loaders.SeasonLoader.Get(ctx, utils.AsInt(obj.Season.ID))
		if err != nil {
			return nil, err
		}

		if s != nil && s.Images != nil {
			fallbacks = append(fallbacks, s.Images)
			sh, err := r.Loaders.ShowLoader.Get(ctx, s.ShowID)
			if err != nil {
				return nil, err
			}
			fallbacks = append(fallbacks, sh.Images)
		}
	}

	return imageOrFallback(ctx, e.Images, style, fallbacks...), nil
}

// Streams is the resolver for the streams field.
func (r *episodeResolver) Streams(ctx context.Context, obj *model.Episode) ([]*model.Stream, error) {
	var out []*model.Stream
	err := user.ValidateAccess(ctx, r.Loaders.EpisodePermissionLoader, utils.AsInt(obj.ID), user.CheckConditions{
		FromDate:    true,
		PublishDate: true,
	})
	if errors.Is(err, user.ErrPublishDateInFuture) {
		return out, nil
	} else if err != nil {
		return nil, err
	}

	e, err := r.GetLoaders().EpisodeLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return nil, err
	}

	r.GetLoaders().AssetStreamsLoader.LoadMany(ctx, lo.Values(e.Assets))

	if e.AssetID.Valid {
		r.GetLoaders().AssetStreamsLoader.Load(ctx, int(e.AssetID.Int64))

		streams, err := r.Resolver.Loaders.AssetStreamsLoader.Get(ctx, int(e.AssetID.Int64))
		if err != nil {
			return nil, err
		}

		for _, s := range streams {
			stream, err := model.StreamFrom(ctx, r.URLSigner, r.Resolver.APIConfig, s, false)
			if err != nil {
				return nil, err
			}

			out = append(out, stream)
		}
	}

	for lang, assetID := range e.Assets {
		languageKey := lang
		streams, err := r.GetLoaders().AssetStreamsLoader.Get(ctx, assetID)
		if err != nil {
			return nil, err
		}

		for _, s := range streams {
			stream, err := model.StreamFrom(ctx, r.URLSigner, r.Resolver.APIConfig, s, false)
			if err != nil {
				return nil, err
			}

			stream.VideoLanguage = &languageKey
			out = append(out, stream)
		}
	}

	out = lo.Filter(out, func(s *model.Stream, _ int) bool {
		return !strings.Contains(s.URL, common.IgnoreEpisodeAssetEndpoint)
	})

	return out, nil
}

// Files is the resolver for the files field.
func (r *episodeResolver) Files(ctx context.Context, obj *model.Episode) ([]*model.File, error) {
	err := user.ValidateAccess(ctx, r.Loaders.EpisodePermissionLoader, utils.AsInt(obj.ID), user.CheckConditions{FromDate: true, PublishDate: true, Download: true})
	if err != nil {
		return nil, nil
	}

	intID, err := strconv.ParseInt(obj.ID, 10, 32)
	if err != nil {
		return nil, err
	}

	files, err := r.Resolver.Loaders.FilesLoader.Get(ctx, int(intID))
	if err != nil {
		return nil, err
	}

	var out []*model.File
	for _, f := range files {
		out = append(out, model.FileFrom(ctx, r.URLSigner, r.Resolver.APIConfig.GetFilesCDNDomain(), f))
	}
	return out, nil
}

// Chapters is the resolver for the chapters field.
func (r *episodeResolver) Chapters(ctx context.Context, obj *model.Episode) ([]*model.Chapter, error) {
	i, err := r.Loaders.EpisodeLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil || !i.AssetID.Valid {
		return nil, err
	}
	metadataItems, err := r.Loaders.TimedMetadataLoader.GetMany(ctx, i.TimedMetadataIDs)
	if err != nil {
		return nil, err
	}
	metadataItems = lo.Filter(metadataItems, func(i *common.TimedMetadata, _ int) bool {
		return i.Type == "chapter"
	})
	ginCtx, _ := utils.GinCtx(ctx)
	languages := user.GetLanguagesFromCtx(ginCtx)

	r.Loaders.PhraseLoader.LoadMany(ctx, lo.Uniq(lo.Map(metadataItems, func(i *common.TimedMetadata, _ int) string {
		return i.ChapterType.Value
	})))

	return lo.Map(metadataItems, func(i *common.TimedMetadata, _ int) *model.Chapter {
		title := i.Title.Get(languages)
		phrase, _ := r.Loaders.PhraseLoader.Get(ctx, i.ChapterType.Value)
		emptyTitle := title == ""
		if emptyTitle && phrase != nil {
			title = phrase.Value.Get(languages)
		}

		switch i.ChapterType {
		case common.ChapterTypeSong, common.ChapterTypeSingAlong:
			if !i.SongID.Valid {
				break
			}
			song, _ := r.Loaders.SongLoader.Get(ctx, i.SongID.UUID)
			if song == nil {
				break
			}
			if emptyTitle {
				if phrase != nil {
					title = fmt.Sprintf("%s - %s", phrase.Value.Get(languages), song.Title.Get(languages))
				} else {
					title = song.Title.Get(languages)
				}
			} else {
				title = strings.Replace(title, "{{song.title}}", song.Title.Get(languages), -1)
			}
		case common.ChapterTypeSpeech, common.ChapterTypeAppeal, common.ChapterTypeTestimony:
			if len(i.PersonIDs) != 1 {
				break
			}
			personID := i.PersonIDs[0]
			person, _ := r.Loaders.PersonLoader.Get(ctx, personID)
			if person == nil {
				break
			}
			if emptyTitle {
				if phrase != nil {
					title = fmt.Sprintf("%s - %s", phrase.Value.Get(languages), person.Name)
				} else {
					title = person.Name
				}
			} else {
				title = strings.Replace(title, "{{person.name}}", person.Name, -1)
			}
		}

		return &model.Chapter{
			ID:          i.ID.String(),
			Title:       title,
			Description: i.Description.GetValueOrNil(languages),
			Start:       int(i.Timestamp),
		}
	}), nil
}

// Season is the resolver for the season field.
func (r *episodeResolver) Season(ctx context.Context, obj *model.Episode) (*model.Season, error) {
	if obj.Season != nil {
		return r.QueryRoot().Season(ctx, obj.Season.ID)
	}
	return nil, nil
}

// Progress is the resolver for the progress field.
func (r *episodeResolver) Progress(ctx context.Context, obj *model.Episode) (*int, error) {
	profileLoaders := r.ProfileLoaders(ctx)
	if profileLoaders == nil {
		return nil, nil
	}
	progress, err := profileLoaders.ProgressLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil || progress == nil {
		return nil, err
	}
	if progress.Progress <= 10 || (float64(progress.Progress)/float64(progress.Duration)) > 0.8 {
		return nil, nil
	}
	return &progress.Progress, nil
}

// Watched is the resolver for the watched field.
func (r *episodeResolver) Watched(ctx context.Context, obj *model.Episode) (bool, error) {
	profileLoaders := r.ProfileLoaders(ctx)
	if profileLoaders == nil {
		return false, nil
	}
	progress, err := profileLoaders.ProgressLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil || progress == nil {
		return false, err
	}
	return progress.Watched > 0, nil
}

// Context is the resolver for the context field.
func (r *episodeResolver) Context(ctx context.Context, obj *model.Episode) (model.EpisodeContextUnion, error) {
	var collectionId *int

	episodeContext, err := r.getEpisodeContext(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	if episodeContext.CollectionID.Valid {
		intId := int(episodeContext.CollectionID.Int64)
		collectionId = &intId
	}

	if collectionId != nil {
		col, err := r.Loaders.CollectionLoader.Get(ctx, *collectionId)
		if err != nil {
			return nil, err
		}
		ginCtx, _ := utils.GinCtx(ctx)
		languages := user.GetLanguagesFromCtx(ginCtx)

		strID := strconv.Itoa(*collectionId)
		return &model.ContextCollection{
			ID:   strID,
			Slug: col.Slugs.GetValueOrNil(languages),
		}, nil
	}
	if obj.Season != nil {
		return r.QueryRoot().Season(ctx, obj.Season.ID)
	}

	return nil, nil
}

// RelatedItems is the resolver for the relatedItems field.
func (r *episodeResolver) RelatedItems(ctx context.Context, obj *model.Episode, first *int, offset *int) (*model.SectionItemPagination, error) {
	var collectionId *int
	if obj.Type == model.EpisodeTypeStandalone {
		ginCtx, err := utils.GinCtx(ctx)
		if err != nil {
			return nil, err
		}
		app, err := applications.GetFromCtx(ginCtx)
		if err != nil {
			return nil, err
		}
		if app.RelatedCollectionID.Valid {
			intID := int(app.RelatedCollectionID.Int64)
			collectionId = &intID
		}
	}

	if collectionId != nil {
		page, err := r.sectionCollectionEntryResolver(ctx, &common.Section{
			CollectionID: null.IntFrom(int64(*collectionId)),
			Style:        "default",
		}, first, offset)
		if err != nil {
			return nil, err
		}
		return &model.SectionItemPagination{
			Total:  page.Total,
			First:  page.First,
			Offset: page.Offset,
			Items:  page.Items,
		}, nil
	}
	return nil, nil
}

// Lessons is the resolver for the lessons field.
func (r *episodeResolver) Lessons(ctx context.Context, obj *model.Episode, first *int, offset *int) (*model.LessonPagination, error) {
	ids, err := r.GetFilteredLoaders(ctx).EpisodeStudyLessonsLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return nil, err
	}
	page := utils.Paginate(ids, first, offset, nil)

	lessons, err := r.Loaders.StudyLessonLoader.GetMany(ctx, utils.PointerArrayToArray(page.Items))
	if err != nil {
		return nil, err
	}

	return &model.LessonPagination{
		Items:  utils.MapWithCtx(ctx, lessons, model.LessonFrom),
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
	}, nil
}

// ShareRestriction is the resolver for the shareCode field.
func (r *episodeResolver) ShareRestriction(ctx context.Context, obj *model.Episode) (model.ShareRestriction, error) {
	perms, err := r.Loaders.EpisodePermissionLoader.Get(ctx, utils.AsInt(obj.ID))
	if err != nil {
		return model.ShareRestrictionPublic, err
	}
	if lo.Contains(perms.Roles.Access, user.RolePublic) {
		return model.ShareRestrictionPublic, nil
	}
	if lo.Contains(perms.Roles.Access, user.RoleBCCMember) {
		return model.ShareRestrictionMembers, nil
	}
	if lo.Contains(perms.Roles.Access, user.RoleRegistered) {
		return model.ShareRestrictionRegistered, nil
	}
	return model.ShareRestrictionPublic, nil
}

// InMyList is the resolver for the inMyList field.
func (r *episodeResolver) InMyList(ctx context.Context, obj *model.Episode) (bool, error) {
	return r.isInMyList(ctx, utils.AsUuid(obj.UUID))
}

// Next is the resolver for the next field.
func (r *episodeResolver) Next(ctx context.Context, obj *model.Episode, limit *int) ([]*model.Episode, error) {
	next, err := r.getNextEpisodes(ctx, obj.ID, limit)
	if err != nil {
		return nil, err
	}
	var episodes []*model.Episode
	for _, id := range next {
		ep, err := r.QueryRoot().Episode(ctx, strconv.Itoa(id), nil)
		if err == nil && ep != nil {
			episodes = append(episodes, ep)
		}
	}
	return episodes, nil
}

// Cursor is the resolver for the cursor field.
func (r *episodeResolver) Cursor(ctx context.Context, obj *model.Episode) (string, error) {
	cursor, err := r.getEpisodeCursor(ctx, obj.ID)
	if err != nil {
		return "", err
	}
	return cursor.Encode()
}

// Episode returns generated.EpisodeResolver implementation.
func (r *Resolver) Episode() generated.EpisodeResolver { return &episodeResolver{r} }

type episodeResolver struct{ *Resolver }
