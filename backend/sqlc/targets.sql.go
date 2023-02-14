// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: targets.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getMemberIDs = `-- name: GetMemberIDs :many
SELECT u.id
FROM users.users u
WHERE $1::bool = true OR u.active_bcc
`

func (q *Queries) GetMemberIDs(ctx context.Context, everyone bool) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getMemberIDs, everyone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var id string
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

const getTargets = `-- name: GetTargets :many
WITH groups AS (SELECT targets_id, array_agg(usergroups_code)::varchar[] as codes
                FROM targets_usergroups
                GROUP BY targets_id)
SELECT t.id,
       t.label,
       t.type,
       g.codes
FROM targets t
         LEFT JOIN groups g ON g.targets_id = t.id
WHERE id = ANY ($1::uuid[])
`

type GetTargetsRow struct {
	ID    uuid.UUID      `db:"id" json:"id"`
	Label null_v4.String `db:"label" json:"label"`
	Type  string         `db:"type" json:"type"`
	Codes []string       `db:"codes" json:"codes"`
}

func (q *Queries) GetTargets(ctx context.Context, dollar_1 []uuid.UUID) ([]GetTargetsRow, error) {
	rows, err := q.db.QueryContext(ctx, getTargets, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTargetsRow
	for rows.Next() {
		var i GetTargetsRow
		if err := rows.Scan(
			&i.ID,
			&i.Label,
			&i.Type,
			pq.Array(&i.Codes),
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
