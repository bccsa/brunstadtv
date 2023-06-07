// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: applications.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getApplicationGroups = `-- name: getApplicationGroups :many
WITH roles AS (SELECT r.applicationgroups_id,
                      array_agg(DISTINCT r.usergroups_code) as roles
               FROM applicationgroups_usergroups r
               GROUP BY r.applicationgroups_id)
SELECT g.id,
       COALESCE(r.roles, '{}')::varchar[] AS roles
FROM applicationgroups g
         LEFT JOIN roles r ON g.id = r.applicationgroups_id
WHERE g.id = ANY ($1::uuid[])
`

type getApplicationGroupsRow struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Roles []string  `db:"roles" json:"roles"`
}

func (q *Queries) getApplicationGroups(ctx context.Context, id []uuid.UUID) ([]getApplicationGroupsRow, error) {
	rows, err := q.db.QueryContext(ctx, getApplicationGroups, pq.Array(id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getApplicationGroupsRow
	for rows.Next() {
		var i getApplicationGroupsRow
		if err := rows.Scan(&i.ID, pq.Array(&i.Roles)); err != nil {
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

const getApplicationIDsForCodes = `-- name: getApplicationIDsForCodes :many
SELECT p.id, p.code
FROM applications p
WHERE p.code = ANY ($1::varchar[])
`

type getApplicationIDsForCodesRow struct {
	ID   int32  `db:"id" json:"id"`
	Code string `db:"code" json:"code"`
}

func (q *Queries) getApplicationIDsForCodes(ctx context.Context, dollar_1 []string) ([]getApplicationIDsForCodesRow, error) {
	rows, err := q.db.QueryContext(ctx, getApplicationIDsForCodes, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getApplicationIDsForCodesRow
	for rows.Next() {
		var i getApplicationIDsForCodesRow
		if err := rows.Scan(&i.ID, &i.Code); err != nil {
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

const getApplications = `-- name: getApplications :many
WITH roles AS (SELECT r.applicationgroups_id,
                      array_agg(DISTINCT r.usergroups_code) as roles
               FROM applicationgroups_usergroups r
               GROUP BY r.applicationgroups_id)
SELECT a.id::int                          AS id,
       a.uuid                             AS uuid,
       a.group_id                         AS group_id,
       a.code::varchar                    AS code,
       a.default                          AS "default",
       a.client_version,
       a.status = 'published'             AS published,
       a.page_id                          AS default_page_id,
       a.search_page_id                   AS search_page_id,
       a.games_page_id                    AS games_page_id,
       a.standalone_related_collection_id AS standalone_related_collection_id,
       COALESCE(r.roles, '{}')::varchar[] AS roles
FROM applications a
         JOIN applicationgroups g ON g.id = a.group_id
         LEFT JOIN roles r ON g.id = r.applicationgroups_id
WHERE a.id = ANY ($1::int[])
  AND a.status = 'published'
`

type getApplicationsRow struct {
	ID                            int32          `db:"id" json:"id"`
	Uuid                          uuid.UUID      `db:"uuid" json:"uuid"`
	GroupID                       uuid.UUID      `db:"group_id" json:"GroupID"`
	Code                          string         `db:"code" json:"code"`
	Default                       bool           `db:"default" json:"default"`
	ClientVersion                 null_v4.String `db:"client_version" json:"clientVersion"`
	Published                     bool           `db:"published" json:"published"`
	DefaultPageID                 null_v4.Int    `db:"default_page_id" json:"defaultPageID"`
	SearchPageID                  null_v4.Int    `db:"search_page_id" json:"searchPageID"`
	GamesPageID                   null_v4.Int    `db:"games_page_id" json:"gamesPageID"`
	StandaloneRelatedCollectionID null_v4.Int    `db:"standalone_related_collection_id" json:"standaloneRelatedCollectionID"`
	Roles                         []string       `db:"roles" json:"roles"`
}

func (q *Queries) getApplications(ctx context.Context, dollar_1 []int32) ([]getApplicationsRow, error) {
	rows, err := q.db.QueryContext(ctx, getApplications, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getApplicationsRow
	for rows.Next() {
		var i getApplicationsRow
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.GroupID,
			&i.Code,
			&i.Default,
			&i.ClientVersion,
			&i.Published,
			&i.DefaultPageID,
			&i.SearchPageID,
			&i.GamesPageID,
			&i.StandaloneRelatedCollectionID,
			pq.Array(&i.Roles),
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

const listApplications = `-- name: listApplications :many
WITH roles AS (SELECT r.applicationgroups_id,
                      array_agg(DISTINCT r.usergroups_code) as roles
               FROM applicationgroups_usergroups r
               GROUP BY r.applicationgroups_id)
SELECT a.id::int                          AS id,
       a.uuid                             AS uuid,
       a.group_id                         AS group_id,
       a.code::varchar                    AS code,
       a.default                          AS "default",
       a.client_version,
       a.status = 'published'             AS published,
       a.page_id                          AS default_page_id,
       a.search_page_id                   AS search_page_id,
       a.games_page_id                    AS games_page_id,
       a.standalone_related_collection_id AS standalone_related_collection_id,
       COALESCE(r.roles, '{}')::varchar[] AS roles
FROM applications a
         JOIN applicationgroups g ON g.id = a.group_id
         LEFT JOIN roles r ON g.id = r.applicationgroups_id
WHERE a.status = 'published'
`

type listApplicationsRow struct {
	ID                            int32          `db:"id" json:"id"`
	Uuid                          uuid.UUID      `db:"uuid" json:"uuid"`
	GroupID                       uuid.UUID      `db:"group_id" json:"GroupID"`
	Code                          string         `db:"code" json:"code"`
	Default                       bool           `db:"default" json:"default"`
	ClientVersion                 null_v4.String `db:"client_version" json:"clientVersion"`
	Published                     bool           `db:"published" json:"published"`
	DefaultPageID                 null_v4.Int    `db:"default_page_id" json:"defaultPageID"`
	SearchPageID                  null_v4.Int    `db:"search_page_id" json:"searchPageID"`
	GamesPageID                   null_v4.Int    `db:"games_page_id" json:"gamesPageID"`
	StandaloneRelatedCollectionID null_v4.Int    `db:"standalone_related_collection_id" json:"standaloneRelatedCollectionID"`
	Roles                         []string       `db:"roles" json:"roles"`
}

func (q *Queries) listApplications(ctx context.Context) ([]listApplicationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listApplications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listApplicationsRow
	for rows.Next() {
		var i listApplicationsRow
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.GroupID,
			&i.Code,
			&i.Default,
			&i.ClientVersion,
			&i.Published,
			&i.DefaultPageID,
			&i.SearchPageID,
			&i.GamesPageID,
			&i.StandaloneRelatedCollectionID,
			pq.Array(&i.Roles),
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
