// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: shows.sql

package sqlexport

import (
	"context"
	"database/sql"
)

const insertShow = `-- name: InsertShow :exec
INSERT INTO shows (id, type, legacy_id, title, description, images, default_episode) VALUES (?, ?, ?, ?, ?, ?, ?)
`

type InsertShowParams struct {
	ID             int64         `db:"id" json:"id"`
	Type           string        `db:"type" json:"type"`
	LegacyID       sql.NullInt64 `db:"legacy_id" json:"legacyId"`
	Title          string        `db:"title" json:"title"`
	Description    string        `db:"description" json:"description"`
	Images         string        `db:"images" json:"images"`
	DefaultEpisode sql.NullInt64 `db:"default_episode" json:"defaultEpisode"`
}

func (q *Queries) InsertShow(ctx context.Context, arg InsertShowParams) error {
	_, err := q.db.ExecContext(ctx, insertShow,
		arg.ID,
		arg.Type,
		arg.LegacyID,
		arg.Title,
		arg.Description,
		arg.Images,
		arg.DefaultEpisode,
	)
	return err
}
