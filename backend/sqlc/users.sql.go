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
SELECT code, string_to_array(emails, E'\n')::text[] as emails
FROM usergroups
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
SELECT array_agg(code)::text[] as groups
FROM usergroups
WHERE $1::text = ANY (string_to_array(emails, E'\n'))
`

func (q *Queries) GetRolesByEmail(ctx context.Context, dollar_1 string) ([]string, error) {
	row := q.db.QueryRowContext(ctx, getRolesByEmail, dollar_1)
	var groups []string
	err := row.Scan(pq.Array(&groups))
	return groups, err
}

const getRolesWithCode = `-- name: GetRolesWithCode :many
SELECT code, string_to_array(emails, E'\n')::text[] as emails
FROM usergroups
WHERE code = ANY ($1::varchar[])
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

const upsertUser = `-- name: UpsertUser :exec
INSERT INTO users.users (id, email, display_name, age, church_ids, active_bcc, roles, age_group, updated_at)
VALUES ($1, $2, $3, $4, $7::int[], $5, $8::varchar[], $6, NOW())
ON CONFLICT (id) DO UPDATE SET email        = excluded.email,
                               display_name = excluded.display_name,
                               age          = excluded.age,
                               church_ids   = excluded.church_ids,
                               active_bcc   = excluded.active_bcc,
                               roles        = excluded.roles,
                               age_group    = excluded.age_group,
                               updated_at   = NOW()
`

type UpsertUserParams struct {
	ID          string   `db:"id" json:"id"`
	Email       string   `db:"email" json:"email"`
	DisplayName string   `db:"display_name" json:"displayName"`
	Age         int32    `db:"age" json:"age"`
	ActiveBcc   bool     `db:"active_bcc" json:"activeBcc"`
	AgeGroup    string   `db:"age_group" json:"ageGroup"`
	ChurchIds   []int32  `db:"church_ids" json:"churchIds"`
	Roles       []string `db:"roles" json:"roles"`
}

func (q *Queries) UpsertUser(ctx context.Context, arg UpsertUserParams) error {
	_, err := q.db.ExecContext(ctx, upsertUser,
		arg.ID,
		arg.Email,
		arg.DisplayName,
		arg.Age,
		arg.ActiveBcc,
		arg.AgeGroup,
		pq.Array(arg.ChurchIds),
		pq.Array(arg.Roles),
	)
	return err
}

const getUsers = `-- name: getUsers :many
SELECT u.id,
       u.email,
       u.display_name,
       u.age,
       u.age_group,
       u.church_ids::int[] as church_ids,
       u.active_bcc,
       u.roles::varchar[]  as roles
FROM users.users u
WHERE u.id = ANY ($1::varchar[])
`

type getUsersRow struct {
	ID          string   `db:"id" json:"id"`
	Email       string   `db:"email" json:"email"`
	DisplayName string   `db:"display_name" json:"displayName"`
	Age         int32    `db:"age" json:"age"`
	AgeGroup    string   `db:"age_group" json:"ageGroup"`
	ChurchIds   []int32  `db:"church_ids" json:"churchIds"`
	ActiveBcc   bool     `db:"active_bcc" json:"activeBcc"`
	Roles       []string `db:"roles" json:"roles"`
}

func (q *Queries) getUsers(ctx context.Context, dollar_1 []string) ([]getUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsers, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getUsersRow
	for rows.Next() {
		var i getUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.DisplayName,
			&i.Age,
			&i.AgeGroup,
			pq.Array(&i.ChurchIds),
			&i.ActiveBcc,
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
