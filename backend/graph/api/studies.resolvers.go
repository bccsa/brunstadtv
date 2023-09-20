package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"errors"
	"strconv"

	"github.com/bcc-code/bcc-media-platform/backend/common"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/generated"
	"github.com/bcc-code/bcc-media-platform/backend/graph/api/model"
	"github.com/bcc-code/bcc-media-platform/backend/user"
	"github.com/bcc-code/bcc-media-platform/backend/utils"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

// Completed is the resolver for the completed field.
func (r *alternativesTaskResolver) Completed(ctx context.Context, obj *model.AlternativesTask) (bool, error) {
	id, err := r.GetProfileLoaders(ctx).TaskCompletedLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return false, err
	}
	return id != nil, nil
}

// Alternatives is the resolver for the alternatives field.
func (r *alternativesTaskResolver) Alternatives(ctx context.Context, obj *model.AlternativesTask) ([]*model.Alternative, error) {
	alts, err := r.Loaders.StudyQuestionAlternativesLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	ginCtx, _ := utils.GinCtx(ctx)
	languages := user.GetLanguagesFromCtx(ginCtx)

	selectedRow, err := r.GetProfileLoaders(ctx).GetSelectedAlternativesLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}

	selectedIDs := []uuid.UUID{}
	if selectedRow != nil {
		selectedIDs = selectedRow.Selected
	}

	return lo.Map(alts, func(alt *common.QuestionAlternative, _ int) *model.Alternative {
		var correct *bool
		if !obj.CompetitionMode {
			correct = &alt.IsCorrect
		}
		return &model.Alternative{
			ID:        alt.ID.String(),
			Title:     alt.Title.Get(languages),
			IsCorrect: correct,
			Selected:  lo.Contains(selectedIDs, alt.ID),
		}
	}), nil
}

// Locked is the resolver for the locked field.
func (r *alternativesTaskResolver) Locked(ctx context.Context, obj *model.AlternativesTask) (bool, error) {
	selectedRow, err := r.GetProfileLoaders(ctx).GetSelectedAlternativesLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil || selectedRow == nil {
		return false, err
	}

	return selectedRow.Locked, err
}

// Image is the resolver for the image field.
func (r *lessonResolver) Image(ctx context.Context, obj *model.Lesson, style *model.ImageStyle) (*string, error) {
	e, err := r.Loaders.StudyLessonLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	t, err := r.Loaders.StudyTopicLoader.Get(ctx, e.TopicID)
	if err != nil {
		return nil, err
	}
	return imageOrFallback(ctx, e.Images, style, t.Images), nil
}

// Tasks is the resolver for the tasks field.
func (r *lessonResolver) Tasks(ctx context.Context, obj *model.Lesson, first *int, offset *int) (*model.TaskPagination, error) {
	ids, err := r.FilteredLoaders(ctx).StudyTasksLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	page := utils.Paginate(ids, first, offset, nil)

	tasks, err := r.Loaders.StudyTaskLoader.GetMany(ctx, utils.PointerArrayToArray(page.Items))
	if err != nil {
		return nil, err
	}

	return &model.TaskPagination{
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
		Items:  utils.MapWithCtx(ctx, tasks, model.TaskFrom),
	}, nil
}

// Topic is the resolver for the topic field.
func (r *lessonResolver) Topic(ctx context.Context, obj *model.Lesson) (*model.StudyTopic, error) {
	return r.QueryRoot().StudyTopic(ctx, obj.Topic.ID)
}

// DefaultEpisode is the resolver for the defaultEpisode field.
func (r *lessonResolver) DefaultEpisode(ctx context.Context, obj *model.Lesson) (*model.Episode, error) {
	episodeIDs, err := r.GetFilteredLoaders(ctx).StudyLessonEpisodesLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	if len(episodeIDs) == 0 {
		return nil, nil
	}
	if episodeIDs[0] == nil {
		return nil, nil
	}
	episode, err := r.QueryRoot().Episode(ctx, strconv.Itoa(*episodeIDs[0]), nil)
	// Permission based errors that shouldn't trigger a failed response
	if errors.Is(err, common.ErrItemNotPublished) ||
		errors.Is(err, common.ErrItemNotFound) ||
		errors.Is(err, common.ErrItemNoAccess) {
		return nil, nil
	}
	return episode, err
}

// Episodes is the resolver for the episodes field.
func (r *lessonResolver) Episodes(ctx context.Context, obj *model.Lesson, first *int, offset *int) (*model.EpisodePagination, error) {
	ids, err := r.GetFilteredLoaders(ctx).StudyLessonEpisodesLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	page := utils.Paginate(ids, first, offset, nil)

	episodes, err := r.Loaders.EpisodeLoader.GetMany(ctx, utils.PointerArrayToArray(page.Items))
	if err != nil {
		return nil, err
	}

	return &model.EpisodePagination{
		Items:  utils.MapWithCtx(ctx, episodes, model.EpisodeFrom),
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
	}, nil
}

// Links is the resolver for the links field.
func (r *lessonResolver) Links(ctx context.Context, obj *model.Lesson, first *int, offset *int) (*model.LinkPagination, error) {
	ids, err := r.GetFilteredLoaders(ctx).StudyLessonLinksLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	page := utils.Paginate(ids, first, offset, nil)

	links, err := r.Loaders.LinkLoader.GetMany(ctx, utils.PointerArrayToArray(page.Items))
	if err != nil {
		return nil, err
	}
	links = lo.Filter(links, func(l *common.Link, _ int) bool {
		return l != nil
	})

	return &model.LinkPagination{
		Items: utils.MapWithCtx(ctx, links, func(ctx context.Context, link *common.Link) *model.Link {
			ginCtx, _ := utils.GinCtx(ctx)
			languages := user.GetLanguagesFromCtx(ginCtx)
			return &model.Link{
				ID:          strconv.Itoa(link.ID),
				URL:         link.URL,
				Title:       link.Title.Get(languages),
				Description: link.Description.GetValueOrNil(languages),
				Image:       link.Images.GetDefault(languages, common.ImageStyleDefault),
			}
		}),
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
	}, nil
}

// Progress is the resolver for the progress field.
func (r *lessonResolver) Progress(ctx context.Context, obj *model.Lesson) (*model.TasksProgress, error) {
	ids, err := r.GetFilteredLoaders(ctx).StudyTasksLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	completed, err := r.GetProfileLoaders(ctx).TaskCompletedLoader.GetMany(ctx, utils.PointerArrayToArray(ids))
	if err != nil {
		return nil, err
	}
	return &model.TasksProgress{
		Total: len(ids),
		Completed: len(lo.Filter(completed, func(i *uuid.UUID, _ int) bool {
			return i != nil
		})),
	}, nil
}

// Completed is the resolver for the completed field.
func (r *lessonResolver) Completed(ctx context.Context, obj *model.Lesson) (bool, error) {
	p, err := getProfile(ctx)
	if err != nil {
		return false, err
	}
	ids, err := r.Loaders.CompletedLessonsLoader.Get(ctx, p.ID)
	if err != nil {
		return false, err
	}
	for _, id := range ids {
		if id != nil && *id == utils.AsUuid(obj.ID) {
			return true, nil
		}
	}
	return false, nil
}

// Locked is the resolver for the locked field.
func (r *lessonResolver) Locked(ctx context.Context, obj *model.Lesson) (bool, error) {
	lockedByPrevious, err := isLessonLockedByPrevious(ctx, r, obj)
	if err != nil || lockedByPrevious {
		return lockedByPrevious, err
	}
	return isLessonLockedByEpisode(ctx, r, obj)
}

// Previous is the resolver for the previous field.
func (r *lessonResolver) Previous(ctx context.Context, obj *model.Lesson) (*model.Lesson, error) {
	l, err := r.Loaders.StudyLessonLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	siblings, err := r.GetFilteredLoaders(ctx).StudyLessonsLoader.Get(ctx, l.TopicID)
	if err != nil {
		return nil, err
	}
	index := -1
	for i, s := range siblings {
		if s != nil && *s == l.ID {
			index = i
			break
		}
	}
	if index <= 0 {
		return nil, nil
	}
	return r.QueryRoot().StudyLesson(ctx, siblings[index-1].String())
}

// Next is the resolver for the next field.
func (r *lessonResolver) Next(ctx context.Context, obj *model.Lesson) (*model.Lesson, error) {
	l, err := r.Loaders.StudyLessonLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	siblings, err := r.GetFilteredLoaders(ctx).StudyLessonsLoader.Get(ctx, l.TopicID)
	if err != nil {
		return nil, err
	}
	index := -1
	for i, s := range siblings {
		if s != nil && *s == l.ID {
			index = i
			break
		}
	}
	if index < 0 || index >= (len(siblings)-1) {
		return nil, nil
	}
	return r.QueryRoot().StudyLesson(ctx, siblings[index+1].String())
}

// Completed is the resolver for the completed field.
func (r *linkTaskResolver) Completed(ctx context.Context, obj *model.LinkTask) (bool, error) {
	id, err := r.GetProfileLoaders(ctx).TaskCompletedLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return false, err
	}
	return id != nil, nil
}

// Link is the resolver for the link field.
func (r *linkTaskResolver) Link(ctx context.Context, obj *model.LinkTask) (*model.Link, error) {
	link, err := r.Loaders.LinkLoader.Get(ctx, utils.AsInt(obj.Link.ID))

	if err != nil {
		return nil, err
	}
	ginCtx, _ := utils.GinCtx(ctx)
	languages := user.GetLanguagesFromCtx(ginCtx)

	t := model.LinkType(link.Type)
	if !t.IsValid() {
		t = model.LinkTypeOther
	}

	return &model.Link{
		ID:          strconv.Itoa(link.ID),
		URL:         link.URL,
		Title:       link.Title.Get(languages),
		Type:        t,
		Description: link.Description.GetValueOrNil(languages),
	}, nil
}

// Completed is the resolver for the completed field.
func (r *posterTaskResolver) Completed(ctx context.Context, obj *model.PosterTask) (bool, error) {
	id, err := r.GetProfileLoaders(ctx).TaskCompletedLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return false, err
	}
	return id != nil, nil
}

// Completed is the resolver for the completed field.
func (r *quoteTaskResolver) Completed(ctx context.Context, obj *model.QuoteTask) (bool, error) {
	id, err := r.GetProfileLoaders(ctx).TaskCompletedLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return false, err
	}
	return id != nil, nil
}

// Image is the resolver for the image field.
func (r *studyTopicResolver) Image(ctx context.Context, obj *model.StudyTopic, style *model.ImageStyle) (*string, error) {
	e, err := r.Loaders.StudyTopicLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	return imageOrFallback(ctx, e.Images, style), nil
}

// DefaultLesson is the resolver for the defaultLesson field.
func (r *studyTopicResolver) DefaultLesson(ctx context.Context, obj *model.StudyTopic) (*model.Lesson, error) {
	uid := utils.AsUuid(obj.ID)
	lessonID, err := r.GetProfileLoaders(ctx).TopicDefaultLessonLoader.Get(ctx, uid)
	if err != nil {
		return nil, err
	}
	if lessonID == nil {
		lessonIDs, err := r.GetFilteredLoaders(ctx).StudyLessonsLoader.Get(ctx, uid)
		if err != nil {
			return nil, err
		}
		if len(lessonIDs) == 0 {
			return nil, ErrItemNotFound
		}
		lessonID = lessonIDs[0]
	}
	if lessonID == nil {
		return nil, ErrItemNotFound
	}

	// Retrieve the first non-completed lesson
	lesson, err := r.QueryRoot().StudyLesson(ctx, lessonID.String())
	if err != nil {
		return nil, err
	}
	locked, err := r.Lesson().Locked(ctx, lesson)
	if err != nil {
		return nil, err
	}

	// If this lesson is locked, check the previous one, with a maximum depth of 5
	for i := 0; locked && i < 5; i++ {
		lesson, err = r.Lesson().Previous(ctx, lesson)
		if err != nil {
			return nil, err
		}
		if lesson == nil {
			return nil, ErrItemNotFound
		}
		locked, err = r.Lesson().Locked(ctx, lesson)
	}
	return lesson, nil
}

// Lessons is the resolver for the lessons field.
func (r *studyTopicResolver) Lessons(ctx context.Context, obj *model.StudyTopic, first *int, offset *int) (*model.LessonPagination, error) {
	ids, err := r.FilteredLoaders(ctx).StudyLessonsLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	page := utils.Paginate(ids, first, offset, nil)

	lessons, err := r.Loaders.StudyLessonLoader.GetMany(ctx, utils.PointerArrayToArray(page.Items))
	if err != nil {
		return nil, err
	}

	return &model.LessonPagination{
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
		Items:  utils.MapWithCtx(ctx, lessons, model.LessonFrom),
	}, nil
}

// Progress is the resolver for the progress field.
func (r *studyTopicResolver) Progress(ctx context.Context, obj *model.StudyTopic) (*model.LessonsProgress, error) {
	p, err := getProfile(ctx)
	if err != nil {
		return nil, err
	}
	ids, err := r.GetFilteredLoaders(ctx).StudyLessonsLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return nil, err
	}
	noPointerIds := utils.PointerArrayToArray(ids)
	completedLessonIDs, err := r.Loaders.CompletedLessonsLoader.Get(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	return &model.LessonsProgress{
		Total: len(noPointerIds),
		Completed: len(lo.Filter(completedLessonIDs, func(i *uuid.UUID, _ int) bool {
			return i != nil && lo.Contains(noPointerIds, *i)
		})),
	}, nil
}

// Completed is the resolver for the completed field.
func (r *textTaskResolver) Completed(ctx context.Context, obj *model.TextTask) (bool, error) {
	id, err := r.GetProfileLoaders(ctx).TaskCompletedLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return false, err
	}
	return id != nil, nil
}

// Completed is the resolver for the completed field.
func (r *videoTaskResolver) Completed(ctx context.Context, obj *model.VideoTask) (bool, error) {
	id, err := r.GetProfileLoaders(ctx).TaskCompletedLoader.Get(ctx, utils.AsUuid(obj.ID))
	if err != nil {
		return false, err
	}
	return id != nil, nil
}

// Episode is the resolver for the episode field.
func (r *videoTaskResolver) Episode(ctx context.Context, obj *model.VideoTask) (*model.Episode, error) {
	return r.QueryRoot().Episode(ctx, obj.Episode.ID, nil)
}

// AlternativesTask returns generated.AlternativesTaskResolver implementation.
func (r *Resolver) AlternativesTask() generated.AlternativesTaskResolver {
	return &alternativesTaskResolver{r}
}

// Lesson returns generated.LessonResolver implementation.
func (r *Resolver) Lesson() generated.LessonResolver { return &lessonResolver{r} }

// LinkTask returns generated.LinkTaskResolver implementation.
func (r *Resolver) LinkTask() generated.LinkTaskResolver { return &linkTaskResolver{r} }

// PosterTask returns generated.PosterTaskResolver implementation.
func (r *Resolver) PosterTask() generated.PosterTaskResolver { return &posterTaskResolver{r} }

// QuoteTask returns generated.QuoteTaskResolver implementation.
func (r *Resolver) QuoteTask() generated.QuoteTaskResolver { return &quoteTaskResolver{r} }

// StudyTopic returns generated.StudyTopicResolver implementation.
func (r *Resolver) StudyTopic() generated.StudyTopicResolver { return &studyTopicResolver{r} }

// TextTask returns generated.TextTaskResolver implementation.
func (r *Resolver) TextTask() generated.TextTaskResolver { return &textTaskResolver{r} }

// VideoTask returns generated.VideoTaskResolver implementation.
func (r *Resolver) VideoTask() generated.VideoTaskResolver { return &videoTaskResolver{r} }

type alternativesTaskResolver struct{ *Resolver }
type lessonResolver struct{ *Resolver }
type linkTaskResolver struct{ *Resolver }
type posterTaskResolver struct{ *Resolver }
type quoteTaskResolver struct{ *Resolver }
type studyTopicResolver struct{ *Resolver }
type textTaskResolver struct{ *Resolver }
type videoTaskResolver struct{ *Resolver }
