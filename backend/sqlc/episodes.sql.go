// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: episodes.sql

package sqlc

import (
	"context"
	"encoding/json"
	"time"

	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

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

const getEpisodes = `-- name: getEpisodes :many
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
       e.legacy_id,
       e.legacy_program_id,
       e.asset_id,
       e.episode_number,
       e.publish_date,
       ea.available_from::timestamp without time zone              AS available_from,
       ea.available_to::timestamp without time zone                AS available_to,
       COALESCE(e.publish_date_in_title, sh.publish_date_in_title, sh.type = 'event', false) AS publish_date_in_title,
       fs.filename_disk                                            as image_file_name,
       e.season_id,
       e.type,
       COALESCE(img.json, '[]')                                    as images,
       ts.title,
       ts.description,
       ts.extra_description,
       tags.tags::int[]                                            AS tag_ids,
       assets.duration                                             as duration,
       COALESCE(e.agerating_code, s.agerating_code, 'A')           as agerating
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
	ID                 int32                 `db:"id" json:"id"`
	LegacyID           null_v4.Int           `db:"legacy_id" json:"legacyID"`
	LegacyProgramID    null_v4.Int           `db:"legacy_program_id" json:"legacyProgramID"`
	AssetID            null_v4.Int           `db:"asset_id" json:"assetID"`
	EpisodeNumber      null_v4.Int           `db:"episode_number" json:"episodeNumber"`
	PublishDate        time.Time             `db:"publish_date" json:"publishDate"`
	AvailableFrom      time.Time             `db:"available_from" json:"availableFrom"`
	AvailableTo        time.Time             `db:"available_to" json:"availableTo"`
	PublishDateInTitle bool                  `db:"publish_date_in_title" json:"publishDateInTitle"`
	ImageFileName      null_v4.String        `db:"image_file_name" json:"imageFileName"`
	SeasonID           null_v4.Int           `db:"season_id" json:"seasonID"`
	Type               string                `db:"type" json:"type"`
	Images             json.RawMessage       `db:"images" json:"images"`
	Title              pqtype.NullRawMessage `db:"title" json:"title"`
	Description        pqtype.NullRawMessage `db:"description" json:"description"`
	ExtraDescription   pqtype.NullRawMessage `db:"extra_description" json:"extraDescription"`
	TagIds             []int32               `db:"tag_ids" json:"tagIds"`
	Duration           null_v4.Int           `db:"duration" json:"duration"`
	Agerating          string                `db:"agerating" json:"agerating"`
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
			&i.LegacyID,
			&i.LegacyProgramID,
			&i.AssetID,
			&i.EpisodeNumber,
			&i.PublishDate,
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
       access.published::bool             AS published,
       access.available_from::timestamp   AS available_from,
       access.available_to::timestamp     AS available_to,
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
	Published             bool      `db:"published" json:"published"`
	AvailableFrom         time.Time `db:"available_from" json:"availableFrom"`
	AvailableTo           time.Time `db:"available_to" json:"availableTo"`
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
			&i.Published,
			&i.AvailableFrom,
			&i.AvailableTo,
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
       e.legacy_id,
       e.legacy_program_id,
       e.asset_id,
       e.episode_number,
       e.publish_date,
       ea.available_from::timestamp without time zone              AS available_from,
       ea.available_to::timestamp without time zone                AS available_to,
       COALESCE(e.publish_date_in_title, sh.publish_date_in_title, sh.type = 'event', false) AS publish_date_in_title,
       fs.filename_disk                                            as image_file_name,
       e.season_id,
       e.type,
       COALESCE(img.json, '[]')                                    as images,
       ts.title,
       ts.description,
       ts.extra_description,
       tags.tags::int[]                                            AS tag_ids,
       assets.duration                                             as duration,
       COALESCE(e.agerating_code, s.agerating_code, 'A')           as agerating
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
	ID                 int32                 `db:"id" json:"id"`
	LegacyID           null_v4.Int           `db:"legacy_id" json:"legacyID"`
	LegacyProgramID    null_v4.Int           `db:"legacy_program_id" json:"legacyProgramID"`
	AssetID            null_v4.Int           `db:"asset_id" json:"assetID"`
	EpisodeNumber      null_v4.Int           `db:"episode_number" json:"episodeNumber"`
	PublishDate        time.Time             `db:"publish_date" json:"publishDate"`
	AvailableFrom      time.Time             `db:"available_from" json:"availableFrom"`
	AvailableTo        time.Time             `db:"available_to" json:"availableTo"`
	PublishDateInTitle bool                  `db:"publish_date_in_title" json:"publishDateInTitle"`
	ImageFileName      null_v4.String        `db:"image_file_name" json:"imageFileName"`
	SeasonID           null_v4.Int           `db:"season_id" json:"seasonID"`
	Type               string                `db:"type" json:"type"`
	Images             json.RawMessage       `db:"images" json:"images"`
	Title              pqtype.NullRawMessage `db:"title" json:"title"`
	Description        pqtype.NullRawMessage `db:"description" json:"description"`
	ExtraDescription   pqtype.NullRawMessage `db:"extra_description" json:"extraDescription"`
	TagIds             []int32               `db:"tag_ids" json:"tagIds"`
	Duration           null_v4.Int           `db:"duration" json:"duration"`
	Agerating          string                `db:"agerating" json:"agerating"`
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
			&i.LegacyID,
			&i.LegacyProgramID,
			&i.AssetID,
			&i.EpisodeNumber,
			&i.PublishDate,
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
