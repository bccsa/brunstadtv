// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: redirects.sql

package sqlc

import (
	"context"
)

const getRedirectByCode = `-- name: GetRedirectByCode :one
SELECT id, status, user_created, date_created, user_updated, date_updated, target_url, code FROM redirects WHERE status = 'published' AND code = $1
`

func (q *Queries) GetRedirectByCode(ctx context.Context, code string) (Redirect, error) {
	row := q.db.QueryRowContext(ctx, getRedirectByCode, code)
	var i Redirect
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.UserCreated,
		&i.DateCreated,
		&i.UserUpdated,
		&i.DateUpdated,
		&i.TargetUrl,
		&i.Code,
	)
	return i, err
}
