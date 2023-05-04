// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: redirects.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getRedirectIDsForCodes = `-- name: getRedirectIDsForCodes :many
SELECT id, code
FROM redirects
WHERE status = 'published'
  AND code = ANY ($1::varchar[])
`

type getRedirectIDsForCodesRow struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Code string    `db:"code" json:"code"`
}

func (q *Queries) getRedirectIDsForCodes(ctx context.Context, dollar_1 []string) ([]getRedirectIDsForCodesRow, error) {
	rows, err := q.db.QueryContext(ctx, getRedirectIDsForCodes, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getRedirectIDsForCodesRow
	for rows.Next() {
		var i getRedirectIDsForCodesRow
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

const getRedirects = `-- name: getRedirects :many
SELECT r.id, r.code, r.target_url, COALESCE(r.include_token, true)::bool as include_token
FROM redirects r
WHERE r.status = 'published'
  AND r.id = ANY ($1::uuid[])
`

type getRedirectsRow struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Code         string    `db:"code" json:"code"`
	TargetUrl    string    `db:"target_url" json:"targetUrl"`
	IncludeToken bool      `db:"include_token" json:"includeToken"`
}

func (q *Queries) getRedirects(ctx context.Context, dollar_1 []uuid.UUID) ([]getRedirectsRow, error) {
	rows, err := q.db.QueryContext(ctx, getRedirects, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getRedirectsRow
	for rows.Next() {
		var i getRedirectsRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.TargetUrl,
			&i.IncludeToken,
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
