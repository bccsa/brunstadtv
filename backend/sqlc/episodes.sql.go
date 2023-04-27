// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: episodes.sql

package sqlc

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getEpisodeIDsForLegacyIDs = `-- name: getEpisodeIDsForLegacyIDs :many
SELECT e.id, e.legacy_id as legacy_id
FROM episodes e
WHERE e.legacy_id = ANY ($1::int[])
`

type getEpisodeIDsForLegacyIDsRow struct {
	ID       int32       `db:"id" json:"id"`
	LegacyID null_v4.Int `db:"legacy_id" json:"legacyID"`
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
	LegacyID null_v4.Int `db:"legacy_id" json:"legacyID"`
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
	SeasonID null_v4.Int `db:"season_id" json:"seasonID"`
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
	SeasonID null_v4.Int `db:"season_id" json:"seasonID"`
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
                   json_object_agg(languages_code, title)             AS title,
                   json_object_agg(languages_code, description)       AS description,
                   json_object_agg(languages_code, extra_description) AS extra_description
            FROM episodes_translations
            WHERE episodes_id = ANY ($1::int[])
            GROUP BY episodes_id),
     tags AS (SELECT episodes_id,
                     array_agg(tags_id) AS tags
              FROM episodes_tags
              GROUP BY episodes_id),
     images AS (WITH images AS (SELECT episode_id, style, language, filename_disk
                                FROM images img
                                         JOIN directus_files df on img.file = df.id
                                WHERE episode_id = ANY ($1::int[]))
                SELECT episode_id, json_agg(images) as json
                FROM images
                GROUP BY episode_id)
SELECT e.id,
       e.uuid,
       e.status,
       e.legacy_id,
       e.legacy_program_id,
       e.asset_id,
       e.episode_number,
       e.publish_date,
       e.production_date,
       e.public_title,
       s.episode_number_in_title                         AS number_in_title,
       COALESCE(e.prevent_public_indexing, false)::bool  as prevent_public_indexing,
       ea.available_from::timestamp without time zone    AS available_from,
       ea.available_to::timestamp without time zone      AS available_to,
       COALESCE(e.publish_date_in_title, sh.publish_date_in_title, sh.type = 'event',
                false)::bool                             AS publish_date_in_title,
       fs.filename_disk                                  as image_file_name,
       e.season_id,
       e.type,
       COALESCE(img.json, '[]')                          as images,
       ts.title,
       ts.description,
       ts.extra_description,
       tags.tags::int[]                                  AS tag_ids,
       assets.duration                                   as duration,
       COALESCE(e.agerating_code, s.agerating_code, 'A') as agerating
FROM episodes e
         LEFT JOIN ts ON e.id = ts.episodes_id
         LEFT JOIN tags ON tags.episodes_id = e.id
         LEFT JOIN images img ON img.episode_id = e.id
         LEFT JOIN assets ON e.asset_id = assets.id
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
	Status                string                `db:"status" json:"status"`
	LegacyID              null_v4.Int           `db:"legacy_id" json:"legacyID"`
	LegacyProgramID       null_v4.Int           `db:"legacy_program_id" json:"legacyProgramID"`
	AssetID               null_v4.Int           `db:"asset_id" json:"assetID"`
	EpisodeNumber         null_v4.Int           `db:"episode_number" json:"episodeNumber"`
	PublishDate           time.Time             `db:"publish_date" json:"publishDate"`
	ProductionDate        time.Time             `db:"production_date" json:"productionDate"`
	PublicTitle           null_v4.String        `db:"public_title" json:"publicTitle"`
	NumberInTitle         sql.NullBool          `db:"number_in_title" json:"numberInTitle"`
	PreventPublicIndexing bool                  `db:"prevent_public_indexing" json:"preventPublicIndexing"`
	AvailableFrom         time.Time             `db:"available_from" json:"availableFrom"`
	AvailableTo           time.Time             `db:"available_to" json:"availableTo"`
	PublishDateInTitle    bool                  `db:"publish_date_in_title" json:"publishDateInTitle"`
	ImageFileName         null_v4.String        `db:"image_file_name" json:"imageFileName"`
	SeasonID              null_v4.Int           `db:"season_id" json:"seasonID"`
	Type                  string                `db:"type" json:"type"`
	Images                json.RawMessage       `db:"images" json:"images"`
	Title                 pqtype.NullRawMessage `db:"title" json:"title"`
	Description           pqtype.NullRawMessage `db:"description" json:"description"`
	ExtraDescription      pqtype.NullRawMessage `db:"extra_description" json:"extraDescription"`
	TagIds                []int32               `db:"tag_ids" json:"tagIds"`
	Duration              null_v4.Int           `db:"duration" json:"duration"`
	Agerating             string                `db:"agerating" json:"agerating"`
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
			&i.Status,
			&i.LegacyID,
			&i.LegacyProgramID,
			&i.AssetID,
			&i.EpisodeNumber,
			&i.PublishDate,
			&i.ProductionDate,
			&i.PublicTitle,
			&i.NumberInTitle,
			&i.PreventPublicIndexing,
			&i.AvailableFrom,
			&i.AvailableTo,
			&i.PublishDateInTitle,
			&i.ImageFileName,
			&i.SeasonID,
			&i.Type,
			&i.Images,
			&i.Title,
			&i.Description,
			&i.ExtraDescription,
			pq.Array(&i.TagIds),
			&i.Duration,
			&i.Agerating,
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
                   json_object_agg(languages_code, title)             AS title,
                   json_object_agg(languages_code, description)       AS description,
                   json_object_agg(languages_code, extra_description) AS extra_description
            FROM episodes_translations
            GROUP BY episodes_id),
     tags AS (SELECT episodes_id,
                     array_agg(tags_id) AS tags
              FROM episodes_tags
              GROUP BY episodes_id),
     images AS (WITH images AS (SELECT episode_id, style, language, filename_disk
                                FROM images img
                                         JOIN directus_files df on img.file = df.id)
                SELECT episode_id, json_agg(images) as json
                FROM images
                GROUP BY episode_id)
SELECT e.id,
       e.uuid,
       e.status,
       e.legacy_id,
       e.legacy_program_id,
       e.asset_id,
       e.episode_number,
       e.publish_date,
       e.production_date,
       e.public_title,
       s.episode_number_in_title                         AS number_in_title,
       COALESCE(e.prevent_public_indexing, false)::bool  as prevent_public_indexing,
       ea.available_from::timestamp without time zone    AS available_from,
       ea.available_to::timestamp without time zone      AS available_to,
       COALESCE(e.publish_date_in_title, sh.publish_date_in_title, sh.type = 'event',
                false)::bool                             AS publish_date_in_title,
       fs.filename_disk                                  as image_file_name,
       e.season_id,
       e.type,
       COALESCE(img.json, '[]')                          as images,
       ts.title,
       ts.description,
       ts.extra_description,
       tags.tags::int[]                                  AS tag_ids,
       assets.duration                                   as duration,
       COALESCE(e.agerating_code, s.agerating_code, 'A') as agerating
FROM episodes e
         LEFT JOIN ts ON e.id = ts.episodes_id
         LEFT JOIN tags ON tags.episodes_id = e.id
         LEFT JOIN images img ON img.episode_id = e.id
         LEFT JOIN assets ON e.asset_id = assets.id
         LEFT JOIN seasons s ON e.season_id = s.id
         LEFT JOIN shows sh ON s.show_id = sh.id
         LEFT JOIN directus_files fs ON fs.id = COALESCE(e.image_file_id, s.image_file_id, sh.image_file_id)
         LEFT JOIN episode_availability ea on e.id = ea.id
`

type listEpisodesRow struct {
	ID                    int32                 `db:"id" json:"id"`
	Uuid                  uuid.UUID             `db:"uuid" json:"uuid"`
	Status                string                `db:"status" json:"status"`
	LegacyID              null_v4.Int           `db:"legacy_id" json:"legacyID"`
	LegacyProgramID       null_v4.Int           `db:"legacy_program_id" json:"legacyProgramID"`
	AssetID               null_v4.Int           `db:"asset_id" json:"assetID"`
	EpisodeNumber         null_v4.Int           `db:"episode_number" json:"episodeNumber"`
	PublishDate           time.Time             `db:"publish_date" json:"publishDate"`
	ProductionDate        time.Time             `db:"production_date" json:"productionDate"`
	PublicTitle           null_v4.String        `db:"public_title" json:"publicTitle"`
	NumberInTitle         sql.NullBool          `db:"number_in_title" json:"numberInTitle"`
	PreventPublicIndexing bool                  `db:"prevent_public_indexing" json:"preventPublicIndexing"`
	AvailableFrom         time.Time             `db:"available_from" json:"availableFrom"`
	AvailableTo           time.Time             `db:"available_to" json:"availableTo"`
	PublishDateInTitle    bool                  `db:"publish_date_in_title" json:"publishDateInTitle"`
	ImageFileName         null_v4.String        `db:"image_file_name" json:"imageFileName"`
	SeasonID              null_v4.Int           `db:"season_id" json:"seasonID"`
	Type                  string                `db:"type" json:"type"`
	Images                json.RawMessage       `db:"images" json:"images"`
	Title                 pqtype.NullRawMessage `db:"title" json:"title"`
	Description           pqtype.NullRawMessage `db:"description" json:"description"`
	ExtraDescription      pqtype.NullRawMessage `db:"extra_description" json:"extraDescription"`
	TagIds                []int32               `db:"tag_ids" json:"tagIds"`
	Duration              null_v4.Int           `db:"duration" json:"duration"`
	Agerating             string                `db:"agerating" json:"agerating"`
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
			&i.Status,
			&i.LegacyID,
			&i.LegacyProgramID,
			&i.AssetID,
			&i.EpisodeNumber,
			&i.PublishDate,
			&i.ProductionDate,
			&i.PublicTitle,
			&i.NumberInTitle,
			&i.PreventPublicIndexing,
			&i.AvailableFrom,
			&i.AvailableTo,
			&i.PublishDateInTitle,
			&i.ImageFileName,
			&i.SeasonID,
			&i.Type,
			&i.Images,
			&i.Title,
			&i.Description,
			&i.ExtraDescription,
			pq.Array(&i.TagIds),
			&i.Duration,
			&i.Agerating,
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
