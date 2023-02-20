// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: user-devices.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getDevicesForProfiles = `-- name: getDevicesForProfiles :many
SELECT token, profile_id, updated_at, name, languages FROM users.devices d WHERE d.profile_id = ANY($1::uuid[]) AND d.updated_at > (NOW() - interval '1 month') ORDER BY updated_at DESC
`

func (q *Queries) getDevicesForProfiles(ctx context.Context, dollar_1 []uuid.UUID) ([]UsersDevice, error) {
	rows, err := q.db.QueryContext(ctx, getDevicesForProfiles, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersDevice
	for rows.Next() {
		var i UsersDevice
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
SELECT token, profile_id, updated_at, name, languages FROM users.devices d WHERE d.updated_at > (NOW() - interval '1 month') ORDER BY updated_at DESC
`

func (q *Queries) listDevices(ctx context.Context) ([]UsersDevice, error) {
	rows, err := q.db.QueryContext(ctx, listDevices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersDevice
	for rows.Next() {
		var i UsersDevice
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
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (token, profile_id) DO UPDATE SET updated_at = EXCLUDED.updated_at, name = EXCLUDED.name, languages = EXCLUDED.languages
`

type setDeviceTokenParams struct {
	Token     string    `db:"token" json:"token"`
	Languages []string  `db:"languages" json:"languages"`
	ProfileID uuid.UUID `db:"profile_id" json:"profileID"`
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
