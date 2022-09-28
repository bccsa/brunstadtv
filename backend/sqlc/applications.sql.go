// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: applications.sql

package sqlc

import (
	"context"

	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

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
WITH roles AS (SELECT r.applications_id,
                      array_agg(DISTINCT r.usergroups_code) as roles
               FROM applications_usergroups r
               GROUP BY r.applications_id)
SELECT a.id::int                          AS id,
       a.code::varchar                    AS code,
       a.default                          AS "default",
       a.client_version,
       a.status = 'published'             AS published,
       a.page_id                          AS default_page_id,
       COALESCE(r.roles, '{}')::varchar[] AS roles
FROM applications a
         LEFT JOIN roles r ON a.id = r.applications_id
WHERE a.id = ANY ($1::int[])
  AND a.status = 'published'
`

type getApplicationsRow struct {
	ID            int32          `db:"id" json:"id"`
	Code          string         `db:"code" json:"code"`
	Default       bool           `db:"default" json:"default"`
	ClientVersion null_v4.String `db:"client_version" json:"clientVersion"`
	Published     bool           `db:"published" json:"published"`
	DefaultPageID null_v4.Int    `db:"default_page_id" json:"defaultPageID"`
	Roles         []string       `db:"roles" json:"roles"`
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
			&i.Code,
			&i.Default,
			&i.ClientVersion,
			&i.Published,
			&i.DefaultPageID,
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
WITH roles AS (SELECT r.applications_id,
                      array_agg(DISTINCT r.usergroups_code) as roles
               FROM applications_usergroups r
               GROUP BY r.applications_id)
SELECT a.id::int                          AS id,
       a.code::varchar                    AS code,
       a.default                          AS "default",
       a.client_version,
       a.status = 'published'             AS published,
       a.page_id                          AS default_page_id,
       COALESCE(r.roles, '{}')::varchar[] AS roles
FROM applications a
         LEFT JOIN roles r ON a.id = r.applications_id
WHERE a.status = 'published'
`

type listApplicationsRow struct {
	ID            int32          `db:"id" json:"id"`
	Code          string         `db:"code" json:"code"`
	Default       bool           `db:"default" json:"default"`
	ClientVersion null_v4.String `db:"client_version" json:"clientVersion"`
	Published     bool           `db:"published" json:"published"`
	DefaultPageID null_v4.Int    `db:"default_page_id" json:"defaultPageID"`
	Roles         []string       `db:"roles" json:"roles"`
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
			&i.Code,
			&i.Default,
			&i.ClientVersion,
			&i.Published,
			&i.DefaultPageID,
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