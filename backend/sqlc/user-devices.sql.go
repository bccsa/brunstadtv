// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user-devices.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const deleteDevices = `-- name: DeleteDevices :exec
DELETE
FROM users.devices d
WHERE d.token = ANY ($1::varchar[])
`

func (q *Queries) DeleteDevices(ctx context.Context, tokens []string) error {
	_, err := q.db.ExecContext(ctx, deleteDevices, pq.Array(tokens))
	return err
}

const listDevicesInApplicationGroup = `-- name: ListDevicesInApplicationGroup :many
SELECT d.token, d.profile_id, d.updated_at, d.name, d.languages::varchar[] as languages
FROM users.devices d
         JOIN users.profiles p ON p.id = d.profile_id
WHERE p.applicationgroup_id = $1::uuid
`

type ListDevicesInApplicationGroupRow struct {
	Token     string    `db:"token" json:"token"`
	ProfileID uuid.UUID `db:"profile_id" json:"profileId"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	Name      string    `db:"name" json:"name"`
	Languages []string  `db:"languages" json:"languages"`
}

func (q *Queries) ListDevicesInApplicationGroup(ctx context.Context, groupID uuid.UUID) ([]ListDevicesInApplicationGroupRow, error) {
	rows, err := q.db.QueryContext(ctx, listDevicesInApplicationGroup, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListDevicesInApplicationGroupRow
	for rows.Next() {
		var i ListDevicesInApplicationGroupRow
		if err := rows.Scan(
			&i.Token,
			&i.ProfileID,
			&i.UpdatedAt,
			&i.Name,
			pq.Array(&i.Languages),
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

const getDevicesForProfiles = `-- name: getDevicesForProfiles :many
SELECT d.token, d.profile_id, d.updated_at, d.name, d.languages::varchar[] as languages
FROM users.devices d
WHERE d.profile_id = ANY ($1::uuid[])
  AND d.updated_at > (NOW() - interval '6 months')
ORDER BY updated_at DESC
`

type getDevicesForProfilesRow struct {
	Token     string    `db:"token" json:"token"`
	ProfileID uuid.UUID `db:"profile_id" json:"profileId"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	Name      string    `db:"name" json:"name"`
	Languages []string  `db:"languages" json:"languages"`
}

func (q *Queries) getDevicesForProfiles(ctx context.Context, dollar_1 []uuid.UUID) ([]getDevicesForProfilesRow, error) {
	rows, err := q.db.QueryContext(ctx, getDevicesForProfiles, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getDevicesForProfilesRow
	for rows.Next() {
		var i getDevicesForProfilesRow
		if err := rows.Scan(
			&i.Token,
			&i.ProfileID,
			&i.UpdatedAt,
			&i.Name,
			pq.Array(&i.Languages),
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

const listDevices = `-- name: listDevices :many
SELECT d.token, d.profile_id, d.updated_at, d.name, d.languages::varchar[] as languages
FROM users.devices d
WHERE d.updated_at > (NOW() - interval '6 months')
ORDER BY updated_at DESC
`

type listDevicesRow struct {
	Token     string    `db:"token" json:"token"`
	ProfileID uuid.UUID `db:"profile_id" json:"profileId"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	Name      string    `db:"name" json:"name"`
	Languages []string  `db:"languages" json:"languages"`
}

func (q *Queries) listDevices(ctx context.Context) ([]listDevicesRow, error) {
	rows, err := q.db.QueryContext(ctx, listDevices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listDevicesRow
	for rows.Next() {
		var i listDevicesRow
		if err := rows.Scan(
			&i.Token,
			&i.ProfileID,
			&i.UpdatedAt,
			&i.Name,
			pq.Array(&i.Languages),
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

const setDeviceToken = `-- name: setDeviceToken :exec
INSERT INTO users.devices (token, languages, profile_id, updated_at, name)
VALUES ($1::varchar, $2::varchar[], $3, $4, $5)
ON CONFLICT (token, profile_id) DO UPDATE SET updated_at = EXCLUDED.updated_at,
                                              name       = EXCLUDED.name,
                                              languages  = EXCLUDED.languages
`

type setDeviceTokenParams struct {
	Token     string    `db:"token" json:"token"`
	Languages []string  `db:"languages" json:"languages"`
	ProfileID uuid.UUID `db:"profile_id" json:"profileId"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	Name      string    `db:"name" json:"name"`
}

func (q *Queries) setDeviceToken(ctx context.Context, arg setDeviceTokenParams) error {
	_, err := q.db.ExecContext(ctx, setDeviceToken,
		arg.Token,
		pq.Array(arg.Languages),
		arg.ProfileID,
		arg.UpdatedAt,
		arg.Name,
	)
	return err
}
