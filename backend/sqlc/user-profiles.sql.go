// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: user-profiles.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getProfiles = `-- name: getProfiles :many
SELECT p.id,
       p.user_id,
       p.name,
       p.applicationgroup_id AS application_group_id
FROM users.profiles p
WHERE applicationgroup_id = $1::uuid
  AND user_id = ANY ($2::varchar[])
`

type getProfilesParams struct {
	ApplicationgroupID uuid.UUID `db:"applicationgroup_id" json:"applicationgroupId"`
	UserID             []string  `db:"user_id" json:"userId"`
}

type getProfilesRow struct {
	ID                 uuid.UUID `db:"id" json:"id"`
	UserID             string    `db:"user_id" json:"userId"`
	Name               string    `db:"name" json:"name"`
	ApplicationGroupID uuid.UUID `db:"application_group_id" json:"applicationGroupId"`
}

func (q *Queries) getProfiles(ctx context.Context, arg getProfilesParams) ([]getProfilesRow, error) {
	rows, err := q.db.QueryContext(ctx, getProfiles, arg.ApplicationgroupID, pq.Array(arg.UserID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getProfilesRow
	for rows.Next() {
		var i getProfilesRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.ApplicationGroupID,
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

const saveProfile = `-- name: saveProfile :exec
INSERT INTO users.profiles (id, user_id, name, applicationgroup_id)
VALUES ($1::uuid, $2::varchar, $3::varchar, $4::uuid)
ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name
`

type saveProfileParams struct {
	ID                 uuid.UUID `db:"id" json:"id"`
	UserID             string    `db:"user_id" json:"userId"`
	Name               string    `db:"name" json:"name"`
	ApplicationgroupID uuid.UUID `db:"applicationgroup_id" json:"applicationgroupId"`
}

func (q *Queries) saveProfile(ctx context.Context, arg saveProfileParams) error {
	_, err := q.db.ExecContext(ctx, saveProfile,
		arg.ID,
		arg.UserID,
		arg.Name,
		arg.ApplicationgroupID,
	)
	return err
}
