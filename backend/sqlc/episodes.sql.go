// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: episodes.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/sqlc-dev/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getEpisodeIDsForLegacyIDs = `-- name: getEpisodeIDsForLegacyIDs :many
SELECT e.id, e.legacy_id as legacy_id
FROM episodes e
WHERE e.legacy_id = ANY ($1::int[])
`

type getEpisodeIDsForLegacyIDsRow struct {
	ID       int32       `db:"id" json:"id"`
	LegacyID null_v4.Int `db:"legacy_id" json:"legacyId"`
}

func (q *Queries) getEpisodeIDsForLegacyIDs(ctx context.Context, dollar_1 []int32) ([]getEpisodeIDsForLegacyIDsRow, error) {
	rows, err := q.db.QueryContext(ctx, getEpisodeIDsForLegacyIDs, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getEpisodeIDsForLegacyIDsRow
	for rows.Next() {
		var i getEpisodeIDsForLegacyIDsRow
		if err := rows.Scan(&i.ID, &i.LegacyID); err != nil {
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

const getEpisodeIDsForLegacyProgramIDs = `-- name: getEpisodeIDsForLegacyProgramIDs :many
SELECT e.id, e.legacy_program_id as legacy_id
FROM episodes e
WHERE e.legacy_program_id = ANY ($1::int[])
`

type getEpisodeIDsForLegacyProgramIDsRow struct {
	ID       int32       `db:"id" json:"id"`
	LegacyID null_v4.Int `db:"legacy_id" json:"legacyId"`
}

func (q *Queries) getEpisodeIDsForLegacyProgramIDs(ctx context.Context, dollar_1 []int32) ([]getEpisodeIDsForLegacyProgramIDsRow, error) {
	rows, err := q.db.QueryContext(ctx, getEpisodeIDsForLegacyProgramIDs, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getEpisodeIDsForLegacyProgramIDsRow
	for rows.Next() {
		var i getEpisodeIDsForLegacyProgramIDsRow
		if err := rows.Scan(&i.ID, &i.LegacyID); err != nil {
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

const getEpisodeIDsForSeasons = `-- name: getEpisodeIDsForSeasons :many
SELECT e.id,
       e.season_id
FROM episodes e
WHERE e.season_id = ANY ($1::int[])
ORDER BY e.episode_number
`

type getEpisodeIDsForSeasonsRow struct {
	ID       int32       `db:"id" json:"id"`
	SeasonID null_v4.Int `db:"season_id" json:"seasonId"`
}

func (q *Queries) getEpisodeIDsForSeasons(ctx context.Context, dollar_1 []int32) ([]getEpisodeIDsForSeasonsRow, error) {
	rows, err := q.db.QueryContext(ctx, getEpisodeIDsForSeasons, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getEpisodeIDsForSeasonsRow
	for rows.Next() {
		var i getEpisodeIDsForSeasonsRow
		if err := rows.Scan(&i.ID, &i.SeasonID); err != nil {
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

const getEpisodeIDsForSeasonsWithRoles = `-- name: getEpisodeIDsForSeasonsWithRoles :many
SELECT e.id,
       e.season_id
FROM episodes e
         LEFT JOIN episode_availability access ON access.id = e.id
         LEFT JOIN episode_roles roles ON roles.id = e.id
WHERE season_id = ANY ($1::int[])
  AND access.published
  AND access.available_to > now()
  AND (
    (roles.roles && $2::varchar[] AND access.available_from < now()) OR
    (roles.roles_earlyaccess && $2::varchar[])
    )
ORDER BY e.episode_number
`

type getEpisodeIDsForSeasonsWithRolesParams struct {
	Column1 []int32  `db:"column_1" json:"column1"`
	Column2 []string `db:"column_2" json:"column2"`
}

type getEpisodeIDsForSeasonsWithRolesRow struct {
	ID       int32       `db:"id" json:"id"`
	SeasonID null_v4.Int `db:"season_id" json:"seasonId"`
}

func (q *Queries) getEpisodeIDsForSeasonsWithRoles(ctx context.Context, arg getEpisodeIDsForSeasonsWithRolesParams) ([]getEpisodeIDsForSeasonsWithRolesRow, error) {
	rows, err := q.db.QueryContext(ctx, getEpisodeIDsForSeasonsWithRoles, pq.Array(arg.Column1), pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getEpisodeIDsForSeasonsWithRolesRow
	for rows.Next() {
		var i getEpisodeIDsForSeasonsWithRolesRow
		if err := rows.Scan(&i.ID, &i.SeasonID); err != nil {
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

const getEpisodeIDsForUuids = `-- name: getEpisodeIDsForUuids :many
SELECT e.id as result, e.uuid as original
FROM episodes e
WHERE e.uuid = ANY ($1::uuid[])
`

type getEpisodeIDsForUuidsRow struct {
	Result   int32     `db:"result" json:"result"`
	Original uuid.UUID `db:"original" json:"original"`
}

func (q *Queries) getEpisodeIDsForUuids(ctx context.Context, ids []uuid.UUID) ([]getEpisodeIDsForUuidsRow, error) {
	rows, err := q.db.QueryContext(ctx, getEpisodeIDsForUuids, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getEpisodeIDsForUuidsRow
	for rows.Next() {
		var i getEpisodeIDsForUuidsRow
		if err := rows.Scan(&i.Result, &i.Original); err != nil {
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

const getEpisodeIDsWithRoles = `-- name: getEpisodeIDsWithRoles :many
SELECT e.id
FROM episodes e
         LEFT JOIN episode_availability access ON access.id = e.id
         LEFT JOIN episode_roles roles ON roles.id = e.id
WHERE e.id = ANY ($1::int[])
  AND access.published
  AND access.available_to > now()
  AND (
    (roles.roles && $2::varchar[] AND access.available_from < now()) OR
    (roles.roles_earlyaccess && $2::varchar[])
    )
`

type getEpisodeIDsWithRolesParams struct {
	Column1 []int32  `db:"column_1" json:"column1"`
	Column2 []string `db:"column_2" json:"column2"`
}

func (q *Queries) getEpisodeIDsWithRoles(ctx context.Context, arg getEpisodeIDsWithRolesParams) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, getEpisodeIDsWithRoles, pq.Array(arg.Column1), pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEpisodeIDsWithTagIDs = `-- name: getEpisodeIDsWithTagIDs :many
SELECT t.episodes_id AS id, t.tags_id AS parent_id
FROM episodes_tags t
         LEFT JOIN episode_availability access ON access.id = t.episodes_id
         LEFT JOIN episode_roles roles ON roles.id = t.episodes_id
WHERE t.tags_id = ANY ($1::int[])
  AND access.published
  AND access.available_to > now()
  AND (
    (roles.roles && $2::varchar[] AND access.available_from < now()) OR
    (roles.roles_earlyaccess && $2::varchar[])
    )
`

type getEpisodeIDsWithTagIDsParams struct {
	TagIds []int32  `db:"tag_ids" json:"tagIds"`
	Roles  []string `db:"roles" json:"roles"`
}

type getEpisodeIDsWithTagIDsRow struct {
	ID       int32 `db:"id" json:"id"`
	ParentID int32 `db:"parent_id" json:"parentId"`
}

func (q *Queries) getEpisodeIDsWithTagIDs(ctx context.Context, arg getEpisodeIDsWithTagIDsParams) ([]getEpisodeIDsWithTagIDsRow, error) {
	rows, err := q.db.QueryContext(ctx, getEpisodeIDsWithTagIDs, pq.Array(arg.TagIds), pq.Array(arg.Roles))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getEpisodeIDsWithTagIDsRow
	for rows.Next() {
		var i getEpisodeIDsWithTagIDsRow
		if err := rows.Scan(&i.ID, &i.ParentID); err != nil {
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

const getEpisodeUUIDsWithRoles = `-- name: getEpisodeUUIDsWithRoles :many
SELECT e.uuid
FROM episodes e
         LEFT JOIN episode_availability access ON access.id = e.id
         LEFT JOIN episode_roles roles ON roles.id = e.id
WHERE e.uuid = ANY ($1::uuid[])
  AND access.published
  AND access.available_to > now()
  AND (
    (roles.roles && $2::varchar[] AND access.available_from < now()) OR
    (roles.roles_earlyaccess && $2::varchar[])
    )
`

type getEpisodeUUIDsWithRolesParams struct {
	Column1 []uuid.UUID `db:"column_1" json:"column1"`
	Column2 []string    `db:"column_2" json:"column2"`
}

func (q *Queries) getEpisodeUUIDsWithRoles(ctx context.Context, arg getEpisodeUUIDsWithRolesParams) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getEpisodeUUIDsWithRoles, pq.Array(arg.Column1), pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uuid.UUID
	for rows.Next() {
		var uuid uuid.UUID
		if err := rows.Scan(&uuid); err != nil {
			return nil, err
		}
		items = append(items, uuid)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEpisodes = `-- name: getEpisodes :many
WITH ts AS (SELECT episodes_id,
                   json_object_agg(languages_code, extra_description) AS extra_description
            FROM episodes_translations
            WHERE episodes_id = ANY ($1::int[])
            GROUP BY episodes_id)
SELECT e.id,
       e.uuid,
       e.mediaitem_id,
       e.status,
       e.legacy_id,
       e.legacy_program_id,
       fs.filename_disk                                                     as image_file_name,
       e.season_id,
       e.type,
       e.episode_number,
       e.public_title,
       s.episode_number_in_title                                            AS number_in_title,
       COALESCE(e.prevent_public_indexing, false)::bool                     as prevent_public_indexing,
       COALESCE(e.publish_date_in_title, false)::bool                       AS publish_date_in_title,
       ea.available_from::timestamp without time zone                       AS available_from,
       ea.available_to::timestamp without time zone                         AS available_to,
       mi.asset_id,
       mi.assets,
       mi.published_at,
       mi.production_date,
       mi.images,
       mi.original_title,
       mi.original_description,
       mi.title,
       mi.description,
       ts.extra_description,
       mi.tag_ids,
       mi.duration,
       mi.asset_date_updated,
       COALESCE(mi.agerating_code, e.agerating_code, s.agerating_code, 'A') as agerating,
       mi.audience,
       mi.content_type,
       mi.timedmetadata_ids
FROM episodes e
         LEFT JOIN mediaitems_by_episodes($1::int[]) mi ON mi.id = e.mediaitem_id
         LEFT JOIN ts ON e.id = ts.episodes_id
         LEFT JOIN seasons s ON e.season_id = s.id
         LEFT JOIN shows sh ON s.show_id = sh.id
         LEFT JOIN directus_files fs ON fs.id = COALESCE(e.image_file_id, s.image_file_id, sh.image_file_id)
         LEFT JOIN episode_availability ea on e.id = ea.id
WHERE e.id = ANY ($1::int[])
ORDER BY e.episode_number
`

type getEpisodesRow struct {
	ID                    int32                 `db:"id" json:"id"`
	Uuid                  uuid.UUID             `db:"uuid" json:"uuid"`
	MediaitemID           uuid.NullUUID         `db:"mediaitem_id" json:"mediaitemId"`
	Status                string                `db:"status" json:"status"`
	LegacyID              null_v4.Int           `db:"legacy_id" json:"legacyId"`
	LegacyProgramID       null_v4.Int           `db:"legacy_program_id" json:"legacyProgramId"`
	ImageFileName         null_v4.String        `db:"image_file_name" json:"imageFileName"`
	SeasonID              null_v4.Int           `db:"season_id" json:"seasonId"`
	Type                  string                `db:"type" json:"type"`
	EpisodeNumber         null_v4.Int           `db:"episode_number" json:"episodeNumber"`
	PublicTitle           null_v4.String        `db:"public_title" json:"publicTitle"`
	NumberInTitle         sql.NullBool          `db:"number_in_title" json:"numberInTitle"`
	PreventPublicIndexing sql.NullBool          `db:"prevent_public_indexing" json:"preventPublicIndexing"`
	PublishDateInTitle    sql.NullBool          `db:"publish_date_in_title" json:"publishDateInTitle"`
	AvailableFrom         null_v4.Time          `db:"available_from" json:"availableFrom"`
	AvailableTo           null_v4.Time          `db:"available_to" json:"availableTo"`
	AssetID               null_v4.Int           `db:"asset_id" json:"assetId"`
	Assets                pqtype.NullRawMessage `db:"assets" json:"assets"`
	PublishedAt           null_v4.Time          `db:"published_at" json:"publishedAt"`
	ProductionDate        null_v4.Time          `db:"production_date" json:"productionDate"`
	Images                pqtype.NullRawMessage `db:"images" json:"images"`
	OriginalTitle         null_v4.String        `db:"original_title" json:"originalTitle"`
	OriginalDescription   null_v4.String        `db:"original_description" json:"originalDescription"`
	Title                 pqtype.NullRawMessage `db:"title" json:"title"`
	Description           pqtype.NullRawMessage `db:"description" json:"description"`
	ExtraDescription      pqtype.NullRawMessage `db:"extra_description" json:"extraDescription"`
	TagIds                []int32               `db:"tag_ids" json:"tagIds"`
	Duration              null_v4.Int           `db:"duration" json:"duration"`
	AssetDateUpdated      null_v4.Time          `db:"asset_date_updated" json:"assetDateUpdated"`
	Agerating             null_v4.String        `db:"agerating" json:"agerating"`
	Audience              null_v4.String        `db:"audience" json:"audience"`
	ContentType           null_v4.String        `db:"content_type" json:"contentType"`
	TimedmetadataIds      []uuid.UUID           `db:"timedmetadata_ids" json:"timedmetadataIds"`
}

func (q *Queries) getEpisodes(ctx context.Context, dollar_1 []int32) ([]getEpisodesRow, error) {
	rows, err := q.db.QueryContext(ctx, getEpisodes, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getEpisodesRow
	for rows.Next() {
		var i getEpisodesRow
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.MediaitemID,
			&i.Status,
			&i.LegacyID,
			&i.LegacyProgramID,
			&i.ImageFileName,
			&i.SeasonID,
			&i.Type,
			&i.EpisodeNumber,
			&i.PublicTitle,
			&i.NumberInTitle,
			&i.PreventPublicIndexing,
			&i.PublishDateInTitle,
			&i.AvailableFrom,
			&i.AvailableTo,
			&i.AssetID,
			&i.Assets,
			&i.PublishedAt,
			&i.ProductionDate,
			&i.Images,
			&i.OriginalTitle,
			&i.OriginalDescription,
			&i.Title,
			&i.Description,
			&i.ExtraDescription,
			pq.Array(&i.TagIds),
			&i.Duration,
			&i.AssetDateUpdated,
			&i.Agerating,
			&i.Audience,
			&i.ContentType,
			pq.Array(&i.TimedmetadataIds),
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

const getPermissionsForEpisodes = `-- name: getPermissionsForEpisodes :many
SELECT e.id,
       e.status = 'unlisted'              AS unlisted,
       access.published::bool             AS published,
       access.available_from::timestamp   AS available_from,
       access.available_to::timestamp     AS available_to,
       access.published_on::timestamp     AS published_on,
       roles.roles::varchar[]             AS usergroups,
       roles.roles_download::varchar[]    AS usergroups_downloads,
       roles.roles_earlyaccess::varchar[] AS usergroups_earlyaccess
FROM episodes e
         LEFT JOIN episode_availability access ON access.id = e.id
         LEFT JOIN episode_roles roles ON roles.id = e.id
WHERE e.id = ANY ($1::int[])
`

type getPermissionsForEpisodesRow struct {
	ID                    int32     `db:"id" json:"id"`
	Unlisted              bool      `db:"unlisted" json:"unlisted"`
	Published             bool      `db:"published" json:"published"`
	AvailableFrom         time.Time `db:"available_from" json:"availableFrom"`
	AvailableTo           time.Time `db:"available_to" json:"availableTo"`
	PublishedOn           time.Time `db:"published_on" json:"publishedOn"`
	Usergroups            []string  `db:"usergroups" json:"usergroups"`
	UsergroupsDownloads   []string  `db:"usergroups_downloads" json:"usergroupsDownloads"`
	UsergroupsEarlyaccess []string  `db:"usergroups_earlyaccess" json:"usergroupsEarlyaccess"`
}

func (q *Queries) getPermissionsForEpisodes(ctx context.Context, dollar_1 []int32) ([]getPermissionsForEpisodesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPermissionsForEpisodes, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getPermissionsForEpisodesRow
	for rows.Next() {
		var i getPermissionsForEpisodesRow
		if err := rows.Scan(
			&i.ID,
			&i.Unlisted,
			&i.Published,
			&i.AvailableFrom,
			&i.AvailableTo,
			&i.PublishedOn,
			pq.Array(&i.Usergroups),
			pq.Array(&i.UsergroupsDownloads),
			pq.Array(&i.UsergroupsEarlyaccess),
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

const listEpisodes = `-- name: listEpisodes :many
WITH ts AS (SELECT episodes_id,
                   json_object_agg(languages_code, extra_description) AS extra_description
            FROM episodes_translations
            GROUP BY episodes_id)
SELECT e.id,
       e.uuid,
       e.mediaitem_id,
       e.status,
       e.legacy_id,
       e.legacy_program_id,
       fs.filename_disk                                                     as image_file_name,
       e.season_id,
       e.type,
       e.episode_number,
       e.public_title,
       s.episode_number_in_title                                            AS number_in_title,
       COALESCE(e.prevent_public_indexing, false)::bool                     as prevent_public_indexing,
       COALESCE(e.publish_date_in_title, false)::bool                       AS publish_date_in_title,
       ea.available_from::timestamp without time zone                       AS available_from,
       ea.available_to::timestamp without time zone                         AS available_to,
       mi.asset_id,
       mi.assets,
       mi.published_at,
       mi.production_date,
       mi.images,
       mi.original_title,
       mi.original_description,
       mi.title,
       mi.description,
       ts.extra_description,
       mi.tag_ids,
       mi.duration,
       mi.asset_date_updated,
       COALESCE(mi.agerating_code, e.agerating_code, s.agerating_code, 'A') as agerating,
       mi.audience,
       mi.content_type,
       mi.timedmetadata_ids
FROM episodes e
         LEFT JOIN mediaitems_view mi ON mi.id = e.mediaitem_id
         LEFT JOIN ts ON e.id = ts.episodes_id
         LEFT JOIN seasons s ON e.season_id = s.id
         LEFT JOIN shows sh ON s.show_id = sh.id
         LEFT JOIN directus_files fs ON fs.id = COALESCE(e.image_file_id, s.image_file_id, sh.image_file_id)
         LEFT JOIN episode_availability ea on e.id = ea.id
`

type listEpisodesRow struct {
	ID                    int32                 `db:"id" json:"id"`
	Uuid                  uuid.UUID             `db:"uuid" json:"uuid"`
	MediaitemID           uuid.NullUUID         `db:"mediaitem_id" json:"mediaitemId"`
	Status                string                `db:"status" json:"status"`
	LegacyID              null_v4.Int           `db:"legacy_id" json:"legacyId"`
	LegacyProgramID       null_v4.Int           `db:"legacy_program_id" json:"legacyProgramId"`
	ImageFileName         null_v4.String        `db:"image_file_name" json:"imageFileName"`
	SeasonID              null_v4.Int           `db:"season_id" json:"seasonId"`
	Type                  string                `db:"type" json:"type"`
	EpisodeNumber         null_v4.Int           `db:"episode_number" json:"episodeNumber"`
	PublicTitle           null_v4.String        `db:"public_title" json:"publicTitle"`
	NumberInTitle         sql.NullBool          `db:"number_in_title" json:"numberInTitle"`
	PreventPublicIndexing bool                  `db:"prevent_public_indexing" json:"preventPublicIndexing"`
	PublishDateInTitle    bool                  `db:"publish_date_in_title" json:"publishDateInTitle"`
	AvailableFrom         time.Time             `db:"available_from" json:"availableFrom"`
	AvailableTo           time.Time             `db:"available_to" json:"availableTo"`
	AssetID               null_v4.Int           `db:"asset_id" json:"assetId"`
	Assets                pqtype.NullRawMessage `db:"assets" json:"assets"`
	PublishedAt           null_v4.Time          `db:"published_at" json:"publishedAt"`
	ProductionDate        null_v4.Time          `db:"production_date" json:"productionDate"`
	Images                pqtype.NullRawMessage `db:"images" json:"images"`
	OriginalTitle         null_v4.String        `db:"original_title" json:"originalTitle"`
	OriginalDescription   null_v4.String        `db:"original_description" json:"originalDescription"`
	Title                 pqtype.NullRawMessage `db:"title" json:"title"`
	Description           pqtype.NullRawMessage `db:"description" json:"description"`
	ExtraDescription      pqtype.NullRawMessage `db:"extra_description" json:"extraDescription"`
	TagIds                []int32               `db:"tag_ids" json:"tagIds"`
	Duration              null_v4.Int           `db:"duration" json:"duration"`
	AssetDateUpdated      null_v4.Time          `db:"asset_date_updated" json:"assetDateUpdated"`
	Agerating             string                `db:"agerating" json:"agerating"`
	Audience              null_v4.String        `db:"audience" json:"audience"`
	ContentType           null_v4.String        `db:"content_type" json:"contentType"`
	TimedmetadataIds      []uuid.UUID           `db:"timedmetadata_ids" json:"timedmetadataIds"`
}

func (q *Queries) listEpisodes(ctx context.Context) ([]listEpisodesRow, error) {
	rows, err := q.db.QueryContext(ctx, listEpisodes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listEpisodesRow
	for rows.Next() {
		var i listEpisodesRow
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.MediaitemID,
			&i.Status,
			&i.LegacyID,
			&i.LegacyProgramID,
			&i.ImageFileName,
			&i.SeasonID,
			&i.Type,
			&i.EpisodeNumber,
			&i.PublicTitle,
			&i.NumberInTitle,
			&i.PreventPublicIndexing,
			&i.PublishDateInTitle,
			&i.AvailableFrom,
			&i.AvailableTo,
			&i.AssetID,
			&i.Assets,
			&i.PublishedAt,
			&i.ProductionDate,
			&i.Images,
			&i.OriginalTitle,
			&i.OriginalDescription,
			&i.Title,
			&i.Description,
			&i.ExtraDescription,
			pq.Array(&i.TagIds),
			&i.Duration,
			&i.AssetDateUpdated,
			&i.Agerating,
			&i.Audience,
			&i.ContentType,
			pq.Array(&i.TimedmetadataIds),
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
