// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: seasons.sql

package sqlexport

import (
	"context"
	"database/sql"
)

const insertSeason = `-- name: InsertSeason :exec
INSERT INTO seasons (id, legacy_id, tag_ids, number, age_rating, title, description, show_id, images)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type InsertSeasonParams struct {
	ID          int64         `db:"id" json:"id"`
	LegacyID    sql.NullInt64 `db:"legacy_id" json:"legacyId"`
	TagIds      string        `db:"tag_ids" json:"tagIds"`
	Number      int64         `db:"number" json:"number"`
	AgeRating   string        `db:"age_rating" json:"ageRating"`
	Title       string        `db:"title" json:"title"`
	Description string        `db:"description" json:"description"`
	ShowID      int64         `db:"show_id" json:"showId"`
	Images      string        `db:"images" json:"images"`
}

func (q *Queries) InsertSeason(ctx context.Context, arg InsertSeasonParams) error {
	_, err := q.db.ExecContext(ctx, insertSeason,
		arg.ID,
		arg.LegacyID,
		arg.TagIds,
		arg.Number,
		arg.AgeRating,
		arg.Title,
		arg.Description,
		arg.ShowID,
		arg.Images,
	)
	return err
}
