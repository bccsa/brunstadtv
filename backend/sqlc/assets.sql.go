// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: assets.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getFilesForAssets = `-- name: getFilesForAssets :many
SELECT 0::int as episodes_id, f.asset_id, f.audio_language_id, f.date_created, f.date_updated, f.extra_metadata, f.id, f.mime_type, f.path, f.storage, f.subtitle_language_id, f.type, f.user_created, f.user_updated
FROM assets a
         JOIN assetfiles f ON a.id = f.asset_id
WHERE a.id = ANY ($1::int[])
`

type getFilesForAssetsRow struct {
	EpisodesID         int32                 `db:"episodes_id" json:"episodesID"`
	AssetID            int32                 `db:"asset_id" json:"assetID"`
	AudioLanguageID    null_v4.String        `db:"audio_language_id" json:"audioLanguageID"`
	DateCreated        time.Time             `db:"date_created" json:"dateCreated"`
	DateUpdated        time.Time             `db:"date_updated" json:"dateUpdated"`
	ExtraMetadata      pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	ID                 int32                 `db:"id" json:"id"`
	MimeType           string                `db:"mime_type" json:"mimeType"`
	Path               string                `db:"path" json:"path"`
	Storage            string                `db:"storage" json:"storage"`
	SubtitleLanguageID null_v4.String        `db:"subtitle_language_id" json:"subtitleLanguageID"`
	Type               string                `db:"type" json:"type"`
	UserCreated        uuid.NullUUID         `db:"user_created" json:"userCreated"`
	UserUpdated        uuid.NullUUID         `db:"user_updated" json:"userUpdated"`
}

func (q *Queries) getFilesForAssets(ctx context.Context, dollar_1 []int32) ([]getFilesForAssetsRow, error) {
	rows, err := q.db.QueryContext(ctx, getFilesForAssets, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getFilesForAssetsRow
	for rows.Next() {
		var i getFilesForAssetsRow
		if err := rows.Scan(
			&i.EpisodesID,
			&i.AssetID,
			&i.AudioLanguageID,
			&i.DateCreated,
			&i.DateUpdated,
			&i.ExtraMetadata,
			&i.ID,
			&i.MimeType,
			&i.Path,
			&i.Storage,
			&i.SubtitleLanguageID,
			&i.Type,
			&i.UserCreated,
			&i.UserUpdated,
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

const getFilesForEpisodes = `-- name: getFilesForEpisodes :many
SELECT e.id AS episodes_id, f.asset_id, f.audio_language_id, f.date_created, f.date_updated, f.extra_metadata, f.id, f.mime_type, f.path, f.storage, f.subtitle_language_id, f.type, f.user_created, f.user_updated
FROM episodes e
         JOIN assets a ON e.asset_id = a.id
         JOIN assetfiles f ON a.id = f.asset_id
WHERE e.id = ANY ($1::int[])
`

type getFilesForEpisodesRow struct {
	EpisodesID         int32                 `db:"episodes_id" json:"episodesID"`
	AssetID            int32                 `db:"asset_id" json:"assetID"`
	AudioLanguageID    null_v4.String        `db:"audio_language_id" json:"audioLanguageID"`
	DateCreated        time.Time             `db:"date_created" json:"dateCreated"`
	DateUpdated        time.Time             `db:"date_updated" json:"dateUpdated"`
	ExtraMetadata      pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	ID                 int32                 `db:"id" json:"id"`
	MimeType           string                `db:"mime_type" json:"mimeType"`
	Path               string                `db:"path" json:"path"`
	Storage            string                `db:"storage" json:"storage"`
	SubtitleLanguageID null_v4.String        `db:"subtitle_language_id" json:"subtitleLanguageID"`
	Type               string                `db:"type" json:"type"`
	UserCreated        uuid.NullUUID         `db:"user_created" json:"userCreated"`
	UserUpdated        uuid.NullUUID         `db:"user_updated" json:"userUpdated"`
}

func (q *Queries) getFilesForEpisodes(ctx context.Context, dollar_1 []int32) ([]getFilesForEpisodesRow, error) {
	rows, err := q.db.QueryContext(ctx, getFilesForEpisodes, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getFilesForEpisodesRow
	for rows.Next() {
		var i getFilesForEpisodesRow
		if err := rows.Scan(
			&i.EpisodesID,
			&i.AssetID,
			&i.AudioLanguageID,
			&i.DateCreated,
			&i.DateUpdated,
			&i.ExtraMetadata,
			&i.ID,
			&i.MimeType,
			&i.Path,
			&i.Storage,
			&i.SubtitleLanguageID,
			&i.Type,
			&i.UserCreated,
			&i.UserUpdated,
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

const getStreamsForAssets = `-- name: getStreamsForAssets :many
WITH audiolang AS (SELECT s.id, array_agg(al.languages_code) langs
                   FROM assets a
                            LEFT JOIN assetstreams s ON a.id = s.asset_id
                            LEFT JOIN assetstreams_audio_languages al ON al.assetstreams_id = s.id
                   WHERE al.languages_code IS NOT NULL
                   GROUP BY s.id),
     sublang AS (SELECT s.id, array_agg(al.languages_code) langs
                 FROM assets a
                          LEFT JOIN assetstreams s ON a.id = s.asset_id
                          LEFT JOIN assetstreams_subtitle_languages al ON al.assetstreams_id = s.id
                 WHERE al.languages_code IS NOT NULL
                 GROUP BY s.id)
SELECT 0::int as episodes_id, s.asset_id, s.date_created, s.date_updated, s.encryption_key_id, s.extra_metadata, s.id, s.legacy_videourl_id, s.path, s.service, s.status, s.type, s.url, s.user_created, s.user_updated, COALESCE(al.langs, '{}')::text[] audio_languages, COALESCE(sl.langs, '{}')::text[] subtitle_languages
FROM assets a
         JOIN assetstreams s ON a.id = s.asset_id
         LEFT JOIN audiolang al ON al.id = s.id
         LEFT JOIN sublang sl ON sl.id = s.id
WHERE a.id = ANY ($1::int[])
`

type getStreamsForAssetsRow struct {
	EpisodesID        int32                 `db:"episodes_id" json:"episodesID"`
	AssetID           int32                 `db:"asset_id" json:"assetID"`
	DateCreated       time.Time             `db:"date_created" json:"dateCreated"`
	DateUpdated       time.Time             `db:"date_updated" json:"dateUpdated"`
	EncryptionKeyID   null_v4.String        `db:"encryption_key_id" json:"encryptionKeyID"`
	ExtraMetadata     pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	ID                int32                 `db:"id" json:"id"`
	LegacyVideourlID  null_v4.Int           `db:"legacy_videourl_id" json:"legacyVideourlID"`
	Path              string                `db:"path" json:"path"`
	Service           string                `db:"service" json:"service"`
	Status            string                `db:"status" json:"status"`
	Type              string                `db:"type" json:"type"`
	Url               string                `db:"url" json:"url"`
	UserCreated       uuid.NullUUID         `db:"user_created" json:"userCreated"`
	UserUpdated       uuid.NullUUID         `db:"user_updated" json:"userUpdated"`
	AudioLanguages    []string              `db:"audio_languages" json:"audioLanguages"`
	SubtitleLanguages []string              `db:"subtitle_languages" json:"subtitleLanguages"`
}

func (q *Queries) getStreamsForAssets(ctx context.Context, dollar_1 []int32) ([]getStreamsForAssetsRow, error) {
	rows, err := q.db.QueryContext(ctx, getStreamsForAssets, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getStreamsForAssetsRow
	for rows.Next() {
		var i getStreamsForAssetsRow
		if err := rows.Scan(
			&i.EpisodesID,
			&i.AssetID,
			&i.DateCreated,
			&i.DateUpdated,
			&i.EncryptionKeyID,
			&i.ExtraMetadata,
			&i.ID,
			&i.LegacyVideourlID,
			&i.Path,
			&i.Service,
			&i.Status,
			&i.Type,
			&i.Url,
			&i.UserCreated,
			&i.UserUpdated,
			pq.Array(&i.AudioLanguages),
			pq.Array(&i.SubtitleLanguages),
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

const getStreamsForEpisodes = `-- name: getStreamsForEpisodes :many
WITH audiolang AS (SELECT s.id, array_agg(al.languages_code) langs
                   FROM episodes e
                            JOIN assets a ON e.asset_id = a.id
                            LEFT JOIN assetstreams s ON a.id = s.asset_id
                            LEFT JOIN assetstreams_audio_languages al ON al.assetstreams_id = s.id
                   WHERE al.languages_code IS NOT NULL
                   GROUP BY s.id),
     sublang AS (SELECT s.id, array_agg(al.languages_code) langs
                 FROM episodes e
                          JOIN assets a ON e.asset_id = a.id
                          LEFT JOIN assetstreams s ON a.id = s.asset_id
                          LEFT JOIN assetstreams_subtitle_languages al ON al.assetstreams_id = s.id
                 WHERE al.languages_code IS NOT NULL
                 GROUP BY s.id)
SELECT e.id AS episodes_id, s.asset_id, s.date_created, s.date_updated, s.encryption_key_id, s.extra_metadata, s.id, s.legacy_videourl_id, s.path, s.service, s.status, s.type, s.url, s.user_created, s.user_updated, COALESCE(al.langs, '{}')::text[] audio_languages, COALESCE(sl.langs, '{}')::text[] subtitle_languages
FROM episodes e
         JOIN assets a ON e.asset_id = a.id
         JOIN assetstreams s ON a.id = s.asset_id
         LEFT JOIN audiolang al ON al.id = s.id
         LEFT JOIN sublang sl ON sl.id = s.id
WHERE e.id = ANY($1::int[])
`

type getStreamsForEpisodesRow struct {
	EpisodesID        int32                 `db:"episodes_id" json:"episodesID"`
	AssetID           int32                 `db:"asset_id" json:"assetID"`
	DateCreated       time.Time             `db:"date_created" json:"dateCreated"`
	DateUpdated       time.Time             `db:"date_updated" json:"dateUpdated"`
	EncryptionKeyID   null_v4.String        `db:"encryption_key_id" json:"encryptionKeyID"`
	ExtraMetadata     pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	ID                int32                 `db:"id" json:"id"`
	LegacyVideourlID  null_v4.Int           `db:"legacy_videourl_id" json:"legacyVideourlID"`
	Path              string                `db:"path" json:"path"`
	Service           string                `db:"service" json:"service"`
	Status            string                `db:"status" json:"status"`
	Type              string                `db:"type" json:"type"`
	Url               string                `db:"url" json:"url"`
	UserCreated       uuid.NullUUID         `db:"user_created" json:"userCreated"`
	UserUpdated       uuid.NullUUID         `db:"user_updated" json:"userUpdated"`
	AudioLanguages    []string              `db:"audio_languages" json:"audioLanguages"`
	SubtitleLanguages []string              `db:"subtitle_languages" json:"subtitleLanguages"`
}

func (q *Queries) getStreamsForEpisodes(ctx context.Context, dollar_1 []int32) ([]getStreamsForEpisodesRow, error) {
	rows, err := q.db.QueryContext(ctx, getStreamsForEpisodes, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getStreamsForEpisodesRow
	for rows.Next() {
		var i getStreamsForEpisodesRow
		if err := rows.Scan(
			&i.EpisodesID,
			&i.AssetID,
			&i.DateCreated,
			&i.DateUpdated,
			&i.EncryptionKeyID,
			&i.ExtraMetadata,
			&i.ID,
			&i.LegacyVideourlID,
			&i.Path,
			&i.Service,
			&i.Status,
			&i.Type,
			&i.Url,
			&i.UserCreated,
			&i.UserUpdated,
			pq.Array(&i.AudioLanguages),
			pq.Array(&i.SubtitleLanguages),
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
