package model

import (
	"context"
	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/user"
	"github.com/bcc-code/brunstadtv/backend/utils"
)

// StudyTopicFrom a studytopic
func StudyTopicFrom(ctx context.Context, topic *common.StudyTopic) *StudyTopic {
	ginCtx, _ := utils.GinCtx(ctx)
	languages := user.GetLanguagesFromCtx(ginCtx)

	return &StudyTopic{
		ID:          topic.ID.String(),
		Title:       topic.Title.Get(languages),
		Description: topic.Description.Get(languages),
	}
}

// StudyTopicSectionItemFrom a studytopic
func StudyTopicSectionItemFrom(ctx context.Context, topic *common.StudyTopic, sort int, imageStyle common.ImageStyle) *SectionItem {
	item := StudyTopicFrom(ctx, topic)

	ginCtx, _ := utils.GinCtx(ctx)
	languages := user.GetLanguagesFromCtx(ginCtx)

	return &SectionItem{
		ID:          topic.ID.String(),
		Item:        item,
		Title:       item.Title,
		Description: item.Description,
		Sort:        sort,
		Image:       topic.Images.GetDefault(languages, imageStyle),
	}
}
