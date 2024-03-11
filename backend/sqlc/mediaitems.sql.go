// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: mediaitems.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const getMediaItemByID = `-- name: GetMediaItemByID :one
SELECT id, user_created, date_created, user_updated, date_updated, label, title, description, type, asset_id, parent_episode_id, parent_starts_at, parent_ends_at, published_at, production_date, parent_id, content_type, audience, agerating_code, translations_required, timedmetadata_from_asset FROM mediaitems WHERE id = $1::uuid
`

func (q *Queries) GetMediaItemByID(ctx context.Context, id uuid.UUID) (Mediaitem, error) {
	row := q.db.QueryRowContext(ctx, getMediaItemByID, id)
	var i Mediaitem
	err := row.Scan(
		&i.ID,
		&i.UserCreated,
		&i.DateCreated,
		&i.UserUpdated,
		&i.DateUpdated,
		&i.Label,
		&i.Title,
		&i.Description,
		&i.Type,
		&i.AssetID,
		&i.ParentEpisodeID,
		&i.ParentStartsAt,
		&i.ParentEndsAt,
		&i.PublishedAt,
		&i.ProductionDate,
		&i.ParentID,
		&i.ContentType,
		&i.Audience,
		&i.AgeratingCode,
		&i.TranslationsRequired,
		&i.TimedmetadataFromAsset,
	)
	return i, err
}
