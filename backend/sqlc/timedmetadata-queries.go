package sqlc

import (
	"context"

	"github.com/bcc-code/bcc-media-platform/backend/common"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

// GetTimedMetadataIdsWithRoles returns metadata ids for the specified roles
func (rq *RoleQueries) GetTimedMetadataIdsWithRoles(ctx context.Context, ids []uuid.UUID) ([]uuid.UUID, error) {
	rows, err := rq.queries.getTimedMetadataIdsWithRoles(ctx, getTimedMetadataIdsWithRolesParams{
		Column1: ids,
		Column2: rq.roles,
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// GetTimedMetadata returns metadata items for the specified ids
func (q *Queries) GetTimedMetadata(ctx context.Context, ids []uuid.UUID) ([]common.TimedMetadata, error) {
	rows, err := q.getTimedMetadata(ctx, ids)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(i getTimedMetadataRow, _ int) common.TimedMetadata {
		title := toLocaleString(i.Title, i.OriginalTitle.String)
		description := toLocaleString(i.Description, i.OriginalDescription.String)
		chapterType := common.ChapterTypes.Parse(i.ChapterType.String)
		if chapterType == nil {
			chapterType = &common.ChapterTypeSpeech
		}
		return common.TimedMetadata{
			ID:          i.ID,
			Type:        i.Type,
			ChapterType: *chapterType,
			PersonIDs:   i.PersonIds,
			SongID:      i.SongID,
			Timestamp:   float64(i.Seconds),
			Title:       title,
			Description: description,
		}
	}), nil
}
