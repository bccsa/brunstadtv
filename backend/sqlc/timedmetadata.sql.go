// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: timedmetadata.sql

package sqlc

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

const clearEpisodeTimedMetadata = `-- name: ClearEpisodeTimedMetadata :exec
DELETE FROM timedmetadata WHERE episode_id = $1
`

func (q *Queries) ClearEpisodeTimedMetadata(ctx context.Context, episodeID null_v4.Int) error {
	_, err := q.db.ExecContext(ctx, clearEpisodeTimedMetadata, episodeID)
	return err
}

const getAssetTimedMetadata = `-- name: GetAssetTimedMetadata :many
SELECT t.id,
       status,
       user_created,
       date_created,
       user_updated,
       date_updated,
       label,
       type,
       highlight,
       title,
       asset_id,
       seconds,
       description,
       episode_id,
       chapter_type,
       song_id,
       (SELECT array_agg(p.persons_id) FROM "timedmetadata_persons" p WHERE p.timedmetadata_id = t.id)::uuid[]  AS person_ids
FROM timedmetadata t
WHERE asset_id = $1
ORDER BY seconds
`

type GetAssetTimedMetadataRow struct {
	ID          uuid.UUID      `db:"id" json:"id"`
	Status      string         `db:"status" json:"status"`
	UserCreated uuid.NullUUID  `db:"user_created" json:"userCreated"`
	DateCreated null_v4.Time   `db:"date_created" json:"dateCreated"`
	UserUpdated uuid.NullUUID  `db:"user_updated" json:"userUpdated"`
	DateUpdated null_v4.Time   `db:"date_updated" json:"dateUpdated"`
	Label       string         `db:"label" json:"label"`
	Type        string         `db:"type" json:"type"`
	Highlight   bool           `db:"highlight" json:"highlight"`
	Title       null_v4.String `db:"title" json:"title"`
	AssetID     null_v4.Int    `db:"asset_id" json:"assetId"`
	Seconds     float32        `db:"seconds" json:"seconds"`
	Description null_v4.String `db:"description" json:"description"`
	EpisodeID   null_v4.Int    `db:"episode_id" json:"episodeId"`
	ChapterType null_v4.String `db:"chapter_type" json:"chapterType"`
	SongID      uuid.NullUUID  `db:"song_id" json:"songId"`
	PersonIds   []uuid.UUID    `db:"person_ids" json:"personIds"`
}

func (q *Queries) GetAssetTimedMetadata(ctx context.Context, assetID null_v4.Int) ([]GetAssetTimedMetadataRow, error) {
	rows, err := q.db.QueryContext(ctx, getAssetTimedMetadata, assetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAssetTimedMetadataRow
	for rows.Next() {
		var i GetAssetTimedMetadataRow
		if err := rows.Scan(
			&i.ID,
			&i.Status,
			&i.UserCreated,
			&i.DateCreated,
			&i.UserUpdated,
			&i.DateUpdated,
			&i.Label,
			&i.Type,
			&i.Highlight,
			&i.Title,
			&i.AssetID,
			&i.Seconds,
			&i.Description,
			&i.EpisodeID,
			&i.ChapterType,
			&i.SongID,
			pq.Array(&i.PersonIds),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertTimedMetadata = `-- name: InsertTimedMetadata :exec
INSERT INTO timedmetadata (id, status, date_created, date_updated, label, type, highlight,
                           title, asset_id, seconds, description, episode_id, chapter_type, song_id)
VALUES ($1, $2, NOW(), NOW(), $3, $4, $5, $6::varchar,
        $7, $8::real, $9::varchar, $10, $11, $12)
`

type InsertTimedMetadataParams struct {
	ID          uuid.UUID      `db:"id" json:"id"`
	Status      string         `db:"status" json:"status"`
	Label       string         `db:"label" json:"label"`
	Type        string         `db:"type" json:"type"`
	Highlight   bool           `db:"highlight" json:"highlight"`
	Title       string         `db:"title" json:"title"`
	AssetID     null_v4.Int    `db:"asset_id" json:"assetId"`
	Seconds     float32        `db:"seconds" json:"seconds"`
	Description string         `db:"description" json:"description"`
	EpisodeID   null_v4.Int    `db:"episode_id" json:"episodeId"`
	ChapterType null_v4.String `db:"chapter_type" json:"chapterType"`
	SongID      uuid.NullUUID  `db:"song_id" json:"songId"`
}

func (q *Queries) InsertTimedMetadata(ctx context.Context, arg InsertTimedMetadataParams) error {
	_, err := q.db.ExecContext(ctx, insertTimedMetadata,
		arg.ID,
		arg.Status,
		arg.Label,
		arg.Type,
		arg.Highlight,
		arg.Title,
		arg.AssetID,
		arg.Seconds,
		arg.Description,
		arg.EpisodeID,
		arg.ChapterType,
		arg.SongID,
	)
	return err
}

const getTimedMetadata = `-- name: getTimedMetadata :many
SELECT md.id,
       md.type,
       md.chapter_type,
       md.song_id,
       (SELECT array_agg(p.persons_id) FROM "timedmetadata_persons" p WHERE p.timedmetadata_id = md.id)::uuid[] AS person_ids,
       md.title                                                  AS original_title,
       md.description                                            AS original_description,
       COALESCE((SELECT json_object_agg(ts.languages_code, ts.title)
                 FROM timedmetadata_translations ts
                 WHERE ts.timedmetadata_id = md.id), '{}')::json AS title,
       COALESCE((SELECT json_object_agg(ts.languages_code, ts.description)
                 FROM timedmetadata_translations ts
                 WHERE ts.timedmetadata_id = md.id), '{}')::json AS description,
       md.seconds,
       md.highlight
FROM timedmetadata md
WHERE md.id = ANY ($1::uuid[])
`

type getTimedMetadataRow struct {
	ID                  uuid.UUID       `db:"id" json:"id"`
	Type                string          `db:"type" json:"type"`
	ChapterType         null_v4.String  `db:"chapter_type" json:"chapterType"`
	SongID              uuid.NullUUID   `db:"song_id" json:"songId"`
	PersonIds           []uuid.UUID     `db:"person_ids" json:"personIds"`
	OriginalTitle       null_v4.String  `db:"original_title" json:"originalTitle"`
	OriginalDescription null_v4.String  `db:"original_description" json:"originalDescription"`
	Title               json.RawMessage `db:"title" json:"title"`
	Description         json.RawMessage `db:"description" json:"description"`
	Seconds             float32         `db:"seconds" json:"seconds"`
	Highlight           bool            `db:"highlight" json:"highlight"`
}

func (q *Queries) getTimedMetadata(ctx context.Context, ids []uuid.UUID) ([]getTimedMetadataRow, error) {
	rows, err := q.db.QueryContext(ctx, getTimedMetadata, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getTimedMetadataRow
	for rows.Next() {
		var i getTimedMetadataRow
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.ChapterType,
			&i.SongID,
			pq.Array(&i.PersonIds),
			&i.OriginalTitle,
			&i.OriginalDescription,
			&i.Title,
			&i.Description,
			&i.Seconds,
			&i.Highlight,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
