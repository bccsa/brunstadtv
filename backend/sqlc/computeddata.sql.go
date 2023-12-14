// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: computeddata.sql

package sqlc

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getComputedForGroups = `-- name: getComputedForGroups :many
WITH conditions AS (SELECT c.computeddata_id, json_agg(c) as conditions
                    FROM computeddata_conditions c
                    GROUP BY c.computeddata_id)
SELECT g.id    as group_id,
       d.id    as id,
       d.value as result,
       c.conditions
FROM computeddatagroups g
         JOIN computeddata d ON d.group_id = g.id
         JOIN conditions c ON c.computeddata_id = d.id
WHERE g.id = ANY ($1::uuid[])
`

type getComputedForGroupsRow struct {
	GroupID    uuid.UUID       `db:"group_id" json:"groupId"`
	ID         uuid.UUID       `db:"id" json:"id"`
	Result     string          `db:"result" json:"result"`
	Conditions json.RawMessage `db:"conditions" json:"conditions"`
}

func (q *Queries) getComputedForGroups(ctx context.Context, dollar_1 []uuid.UUID) ([]getComputedForGroupsRow, error) {
	rows, err := q.db.QueryContext(ctx, getComputedForGroups, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getComputedForGroupsRow
	for rows.Next() {
		var i getComputedForGroupsRow
		if err := rows.Scan(
			&i.GroupID,
			&i.ID,
			&i.Result,
			&i.Conditions,
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
