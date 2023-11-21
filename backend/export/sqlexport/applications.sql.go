// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: applications.sql

package sqlexport

import (
	"context"
	"database/sql"
)

const insertApplication = `-- name: InsertApplication :exec
INSERT INTO applications (id, code, client_version, default_page_id)
VALUES (?,?,?,?)
`

type InsertApplicationParams struct {
	ID            int64         `db:"id" json:"id"`
	Code          string        `db:"code" json:"code"`
	ClientVersion string        `db:"client_version" json:"clientVersion"`
	DefaultPageID sql.NullInt64 `db:"default_page_id" json:"defaultPageId"`
}

func (q *Queries) InsertApplication(ctx context.Context, arg InsertApplicationParams) error {
	_, err := q.db.ExecContext(ctx, insertApplication,
		arg.ID,
		arg.Code,
		arg.ClientVersion,
		arg.DefaultPageID,
	)
	return err
}
