// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: directus-users.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getUserIDByEmail = `-- name: GetUserIDByEmail :one
SELECT id FROM directus_users WHERE email = $1
`

func (q *Queries) GetUserIDByEmail(ctx context.Context, email null_v4.String) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getUserIDByEmail, email)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
