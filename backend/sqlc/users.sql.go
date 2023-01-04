// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users.sql

package sqlc

import (
	"context"

	"github.com/lib/pq"
)

const getRoles = `-- name: GetRoles :many
SELECT code, string_to_array(emails, E'\n')::text[] as emails  FROM usergroups
`

type GetRolesRow struct {
	Code   string   `db:"code" json:"code"`
	Emails []string `db:"emails" json:"emails"`
}

func (q *Queries) GetRoles(ctx context.Context) ([]GetRolesRow, error) {
	rows, err := q.db.QueryContext(ctx, getRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRolesRow
	for rows.Next() {
		var i GetRolesRow
		if err := rows.Scan(&i.Code, pq.Array(&i.Emails)); err != nil {
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

const getRolesByEmail = `-- name: GetRolesByEmail :one
SELECT array_agg(code)::text[] as groups FROM usergroups WHERE $1::text = ANY(string_to_array(emails, E'\n'))
`

func (q *Queries) GetRolesByEmail(ctx context.Context, dollar_1 string) ([]string, error) {
	row := q.db.QueryRowContext(ctx, getRolesByEmail, dollar_1)
	var groups []string
	err := row.Scan(pq.Array(&groups))
	return groups, err
}

const getRolesWithCode = `-- name: GetRolesWithCode :many
SELECT code, string_to_array(emails, E'\n')::text[] as emails  FROM usergroups WHERE code = ANY($1::varchar[])
`

type GetRolesWithCodeRow struct {
	Code   string   `db:"code" json:"code"`
	Emails []string `db:"emails" json:"emails"`
}

func (q *Queries) GetRolesWithCode(ctx context.Context, dollar_1 []string) ([]GetRolesWithCodeRow, error) {
	rows, err := q.db.QueryContext(ctx, getRolesWithCode, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRolesWithCodeRow
	for rows.Next() {
		var i GetRolesWithCodeRow
		if err := rows.Scan(&i.Code, pq.Array(&i.Emails)); err != nil {
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
