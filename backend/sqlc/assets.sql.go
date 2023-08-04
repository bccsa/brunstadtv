// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: assets.sql

package sqlc

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/sqlc-dev/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const assetIDByARN = `-- name: AssetIDByARN :one
SELECT id
FROM assets
WHERE aws_arn = $1::varchar
LIMIT 1
`

func (q *Queries) AssetIDByARN(ctx context.Context, awsArn string) (int32, error) {
	row := q.db.QueryRowContext(ctx, assetIDByARN, awsArn)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deletePath = `-- name: DeletePath :exec
DELETE
FROM assets
WHERE main_storage_path = $1
`

func (q *Queries) DeletePath(ctx context.Context, path null_v4.String) error {
	_, err := q.db.ExecContext(ctx, deletePath, path)
	return err
}

const insertAsset = `-- name: InsertAsset :one
INSERT INTO assets (duration, encoding_version, legacy_id, main_storage_path,
                    mediabanken_id, name, status, aws_arn, date_updated, date_created)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(),
        NOW())
RETURNING id
`

type InsertAssetParams struct {
	Duration        int32          `db:"duration" json:"duration"`
	EncodingVersion null_v4.String `db:"encoding_version" json:"encodingVersion"`
	LegacyID        null_v4.Int    `db:"legacy_id" json:"legacyId"`
	MainStoragePath null_v4.String `db:"main_storage_path" json:"mainStoragePath"`
	MediabankenID   null_v4.String `db:"mediabanken_id" json:"mediabankenId"`
	Name            string         `db:"name" json:"name"`
	Status          null_v4.String `db:"status" json:"status"`
	AwsArn          null_v4.String `db:"aws_arn" json:"awsArn"`
}

func (q *Queries) InsertAsset(ctx context.Context, arg InsertAssetParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertAsset,
		arg.Duration,
		arg.EncodingVersion,
		arg.LegacyID,
		arg.MainStoragePath,
		arg.MediabankenID,
		arg.Name,
		arg.Status,
		arg.AwsArn,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertAssetFile = `-- name: InsertAssetFile :one
INSERT INTO assetfiles (asset_id, audio_language_id, subtitle_language_id, size, path, resolution, mime_type, type,
                        storage, date_updated, date_created)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,
        NOW(), NOW())
RETURNING id
`

type InsertAssetFileParams struct {
	AssetID            int32          `db:"asset_id" json:"assetId"`
	AudioLanguageID    null_v4.String `db:"audio_language_id" json:"audioLanguageId"`
	SubtitleLanguageID null_v4.String `db:"subtitle_language_id" json:"subtitleLanguageId"`
	Size               int64          `db:"size" json:"size"`
	Path               string         `db:"path" json:"path"`
	Resolution         null_v4.String `db:"resolution" json:"resolution"`
	MimeType           string         `db:"mime_type" json:"mimeType"`
	Type               string         `db:"type" json:"type"`
	Storage            string         `db:"storage" json:"storage"`
}

func (q *Queries) InsertAssetFile(ctx context.Context, arg InsertAssetFileParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertAssetFile,
		arg.AssetID,
		arg.AudioLanguageID,
		arg.SubtitleLanguageID,
		arg.Size,
		arg.Path,
		arg.Resolution,
		arg.MimeType,
		arg.Type,
		arg.Storage,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertAssetStream = `-- name: InsertAssetStream :one
INSERT INTO assetstreams (asset_id, encryption_key_id, extra_metadata, legacy_videourl_id, path, service, status, type,
                          url, date_updated, date_created)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,
        NOW(), NOW())
RETURNING id
`

type InsertAssetStreamParams struct {
	AssetID          int32                 `db:"asset_id" json:"assetId"`
	EncryptionKeyID  null_v4.String        `db:"encryption_key_id" json:"encryptionKeyId"`
	ExtraMetadata    pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	LegacyVideourlID null_v4.Int           `db:"legacy_videourl_id" json:"legacyVideourlId"`
	Path             string                `db:"path" json:"path"`
	Service          string                `db:"service" json:"service"`
	Status           string                `db:"status" json:"status"`
	Type             string                `db:"type" json:"type"`
	Url              string                `db:"url" json:"url"`
}

func (q *Queries) InsertAssetStream(ctx context.Context, arg InsertAssetStreamParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertAssetStream,
		arg.AssetID,
		arg.EncryptionKeyID,
		arg.ExtraMetadata,
		arg.LegacyVideourlID,
		arg.Path,
		arg.Service,
		arg.Status,
		arg.Type,
		arg.Url,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertAssetStreamAudioLanguage = `-- name: InsertAssetStreamAudioLanguage :one
INSERT INTO assetstreams_audio_languages (assetstreams_id, languages_code)
VALUES ($1, $2)
RETURNING id
`

type InsertAssetStreamAudioLanguageParams struct {
	AssetstreamsID null_v4.Int    `db:"assetstreams_id" json:"assetstreamsId"`
	LanguagesCode  null_v4.String `db:"languages_code" json:"languagesCode"`
}

func (q *Queries) InsertAssetStreamAudioLanguage(ctx context.Context, arg InsertAssetStreamAudioLanguageParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertAssetStreamAudioLanguage, arg.AssetstreamsID, arg.LanguagesCode)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertAssetStreamSubtitleLanguage = `-- name: InsertAssetStreamSubtitleLanguage :one
INSERT INTO assetstreams_subtitle_languages (assetstreams_id, languages_code)
VALUES ($1, $2)
RETURNING id
`

type InsertAssetStreamSubtitleLanguageParams struct {
	AssetstreamsID null_v4.Int    `db:"assetstreams_id" json:"assetstreamsId"`
	LanguagesCode  null_v4.String `db:"languages_code" json:"languagesCode"`
}

func (q *Queries) InsertAssetStreamSubtitleLanguage(ctx context.Context, arg InsertAssetStreamSubtitleLanguageParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertAssetStreamSubtitleLanguage, arg.AssetstreamsID, arg.LanguagesCode)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const listAssets = `-- name: ListAssets :many
SELECT date_created, date_updated, duration, encoding_version, id, legacy_id, main_storage_path, mediabanken_id, name, status, user_created, user_updated, aws_arn
FROM assets
`

func (q *Queries) ListAssets(ctx context.Context) ([]Asset, error) {
	rows, err := q.db.QueryContext(ctx, listAssets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Asset
	for rows.Next() {
		var i Asset
		if err := rows.Scan(
			&i.DateCreated,
			&i.DateUpdated,
			&i.Duration,
			&i.EncodingVersion,
			&i.ID,
			&i.LegacyID,
			&i.MainStoragePath,
			&i.MediabankenID,
			&i.Name,
			&i.Status,
			&i.UserCreated,
			&i.UserUpdated,
			&i.AwsArn,
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

const newestPreviousAsset = `-- name: NewestPreviousAsset :one
SELECT date_created, date_updated, duration, encoding_version, id, legacy_id, main_storage_path, mediabanken_id, name, status, user_created, user_updated, aws_arn
FROM assets
WHERE mediabanken_id = $1::varchar
ORDER BY date_created DESC
LIMIT 1
`

func (q *Queries) NewestPreviousAsset(ctx context.Context, mediabankenID string) (Asset, error) {
	row := q.db.QueryRowContext(ctx, newestPreviousAsset, mediabankenID)
	var i Asset
	err := row.Scan(
		&i.DateCreated,
		&i.DateUpdated,
		&i.Duration,
		&i.EncodingVersion,
		&i.ID,
		&i.LegacyID,
		&i.MainStoragePath,
		&i.MediabankenID,
		&i.Name,
		&i.Status,
		&i.UserCreated,
		&i.UserUpdated,
		&i.AwsArn,
	)
	return i, err
}

const updateAssetArn = `-- name: UpdateAssetArn :exec
UPDATE assets
SET aws_arn = $1
WHERE id = $2
`

type UpdateAssetArnParams struct {
	AwsArn null_v4.String `db:"aws_arn" json:"awsArn"`
	ID     int32          `db:"id" json:"id"`
}

func (q *Queries) UpdateAssetArn(ctx context.Context, arg UpdateAssetArnParams) error {
	_, err := q.db.ExecContext(ctx, updateAssetArn, arg.AwsArn, arg.ID)
	return err
}

const updateAssetStatus = `-- name: UpdateAssetStatus :exec
UPDATE assets
SET status = $1::varchar
WHERE id = $2
`

type UpdateAssetStatusParams struct {
	Status string `db:"status" json:"status"`
	ID     int32  `db:"id" json:"id"`
}

func (q *Queries) UpdateAssetStatus(ctx context.Context, arg UpdateAssetStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateAssetStatus, arg.Status, arg.ID)
	return err
}

const getFilesForAssets = `-- name: getFilesForAssets :many
SELECT 0::int as episodes_id, f.asset_id, f.audio_language_id, f.date_created, f.date_updated, f.extra_metadata, f.id, f.mime_type, f.path, f.storage, f.subtitle_language_id, f.type, f.user_created, f.user_updated, f.resolution, f.size
FROM assets a
         JOIN assetfiles f ON a.id = f.asset_id
WHERE a.id = ANY ($1::int[])
`

type getFilesForAssetsRow struct {
	EpisodesID         int32                 `db:"episodes_id" json:"episodesId"`
	AssetID            int32                 `db:"asset_id" json:"assetId"`
	AudioLanguageID    null_v4.String        `db:"audio_language_id" json:"audioLanguageId"`
	DateCreated        time.Time             `db:"date_created" json:"dateCreated"`
	DateUpdated        time.Time             `db:"date_updated" json:"dateUpdated"`
	ExtraMetadata      pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	ID                 int32                 `db:"id" json:"id"`
	MimeType           string                `db:"mime_type" json:"mimeType"`
	Path               string                `db:"path" json:"path"`
	Storage            string                `db:"storage" json:"storage"`
	SubtitleLanguageID null_v4.String        `db:"subtitle_language_id" json:"subtitleLanguageId"`
	Type               string                `db:"type" json:"type"`
	UserCreated        uuid.NullUUID         `db:"user_created" json:"userCreated"`
	UserUpdated        uuid.NullUUID         `db:"user_updated" json:"userUpdated"`
	Resolution         null_v4.String        `db:"resolution" json:"resolution"`
	Size               int64                 `db:"size" json:"size"`
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
			&i.Resolution,
			&i.Size,
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
SELECT e.id AS episodes_id, f.asset_id, f.audio_language_id, f.date_created, f.date_updated, f.extra_metadata, f.id, f.mime_type, f.path, f.storage, f.subtitle_language_id, f.type, f.user_created, f.user_updated, f.resolution, f.size
FROM episodes e
         JOIN assets a ON e.asset_id = a.id
         JOIN assetfiles f ON a.id = f.asset_id
WHERE e.id = ANY ($1::int[])
`

type getFilesForEpisodesRow struct {
	EpisodesID         int32                 `db:"episodes_id" json:"episodesId"`
	AssetID            int32                 `db:"asset_id" json:"assetId"`
	AudioLanguageID    null_v4.String        `db:"audio_language_id" json:"audioLanguageId"`
	DateCreated        time.Time             `db:"date_created" json:"dateCreated"`
	DateUpdated        time.Time             `db:"date_updated" json:"dateUpdated"`
	ExtraMetadata      pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	ID                 int32                 `db:"id" json:"id"`
	MimeType           string                `db:"mime_type" json:"mimeType"`
	Path               string                `db:"path" json:"path"`
	Storage            string                `db:"storage" json:"storage"`
	SubtitleLanguageID null_v4.String        `db:"subtitle_language_id" json:"subtitleLanguageId"`
	Type               string                `db:"type" json:"type"`
	UserCreated        uuid.NullUUID         `db:"user_created" json:"userCreated"`
	UserUpdated        uuid.NullUUID         `db:"user_updated" json:"userUpdated"`
	Resolution         null_v4.String        `db:"resolution" json:"resolution"`
	Size               int64                 `db:"size" json:"size"`
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
			&i.Resolution,
			&i.Size,
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
SELECT 0::int as                        episodes_id,
       s.asset_id, s.date_created, s.date_updated, s.encryption_key_id, s.extra_metadata, s.id, s.legacy_videourl_id, s.path, s.service, s.status, s.type, s.url, s.user_created, s.user_updated,
       COALESCE(al.langs, '{}')::text[] audio_languages,
       COALESCE(sl.langs, '{}')::text[] subtitle_languages
FROM assets a
         JOIN assetstreams s ON a.id = s.asset_id
         LEFT JOIN audiolang al ON al.id = s.id
         LEFT JOIN sublang sl ON sl.id = s.id
WHERE a.id = ANY ($1::int[])
`

type getStreamsForAssetsRow struct {
	EpisodesID        int32                 `db:"episodes_id" json:"episodesId"`
	AssetID           int32                 `db:"asset_id" json:"assetId"`
	DateCreated       time.Time             `db:"date_created" json:"dateCreated"`
	DateUpdated       time.Time             `db:"date_updated" json:"dateUpdated"`
	EncryptionKeyID   null_v4.String        `db:"encryption_key_id" json:"encryptionKeyId"`
	ExtraMetadata     pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	ID                int32                 `db:"id" json:"id"`
	LegacyVideourlID  null_v4.Int           `db:"legacy_videourl_id" json:"legacyVideourlId"`
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
SELECT e.id AS                          episodes_id,
       s.asset_id, s.date_created, s.date_updated, s.encryption_key_id, s.extra_metadata, s.id, s.legacy_videourl_id, s.path, s.service, s.status, s.type, s.url, s.user_created, s.user_updated,
       COALESCE(al.langs, '{}')::text[] audio_languages,
       COALESCE(sl.langs, '{}')::text[] subtitle_languages
FROM episodes e
         JOIN assets a ON e.asset_id = a.id
         JOIN assetstreams s ON a.id = s.asset_id
         LEFT JOIN audiolang al ON al.id = s.id
         LEFT JOIN sublang sl ON sl.id = s.id
WHERE e.id = ANY ($1::int[])
`

type getStreamsForEpisodesRow struct {
	EpisodesID        int32                 `db:"episodes_id" json:"episodesId"`
	AssetID           int32                 `db:"asset_id" json:"assetId"`
	DateCreated       time.Time             `db:"date_created" json:"dateCreated"`
	DateUpdated       time.Time             `db:"date_updated" json:"dateUpdated"`
	EncryptionKeyID   null_v4.String        `db:"encryption_key_id" json:"encryptionKeyId"`
	ExtraMetadata     pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	ID                int32                 `db:"id" json:"id"`
	LegacyVideourlID  null_v4.Int           `db:"legacy_videourl_id" json:"legacyVideourlId"`
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

const getTimedMetadataForAssets = `-- name: getTimedMetadataForAssets :many
SELECT md.id,
       md.asset_id,
       md.type,
       md.title                            AS original_title,
       md.description                      AS original_description,
       (SELECT json_object_agg(ts.languages_code, ts.title)
        FROM timedmetadata_translations ts
        WHERE ts.timedmetadata_id = md.id) AS title,
       (SELECT json_object_agg(ts.languages_code, ts.description)
        FROM timedmetadata_translations ts
        WHERE ts.timedmetadata_id = md.id) AS description,
       md.timestamp,
       md.highlight
FROM timedmetadata md
WHERE md.asset_id = ANY ($1::int[])
`

type getTimedMetadataForAssetsRow struct {
	ID                  uuid.UUID       `db:"id" json:"id"`
	AssetID             int32           `db:"asset_id" json:"assetId"`
	Type                string          `db:"type" json:"type"`
	OriginalTitle       string          `db:"original_title" json:"originalTitle"`
	OriginalDescription null_v4.String  `db:"original_description" json:"originalDescription"`
	Title               json.RawMessage `db:"title" json:"title"`
	Description         json.RawMessage `db:"description" json:"description"`
	Timestamp           time.Time       `db:"timestamp" json:"timestamp"`
	Highlight           bool            `db:"highlight" json:"highlight"`
}

func (q *Queries) getTimedMetadataForAssets(ctx context.Context, assetIds []int32) ([]getTimedMetadataForAssetsRow, error) {
	rows, err := q.db.QueryContext(ctx, getTimedMetadataForAssets, pq.Array(assetIds))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getTimedMetadataForAssetsRow
	for rows.Next() {
		var i getTimedMetadataForAssetsRow
		if err := rows.Scan(
			&i.ID,
			&i.AssetID,
			&i.Type,
			&i.OriginalTitle,
			&i.OriginalDescription,
			&i.Title,
			&i.Description,
			&i.Timestamp,
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
