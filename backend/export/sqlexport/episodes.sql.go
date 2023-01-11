// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: episodes.sql

package sqlexport

import (
	"context"
	"database/sql"
)

const insertEpisode = `-- name: InsertEpisode :exec
INSERT INTO episodes ( id, legacy_id, legacy_program_id, age_rating, title, description, extra_description, images, production_date, season_id, duration, number)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type InsertEpisodeParams struct {
	ID               int64          `db:"id" json:"id"`
	LegacyID         sql.NullInt64  `db:"legacy_id" json:"legacyID"`
	LegacyProgramID  sql.NullInt64  `db:"legacy_program_id" json:"legacyProgramID"`
	AgeRating        string         `db:"age_rating" json:"ageRating"`
	Title            string         `db:"title" json:"title"`
	Description      string         `db:"description" json:"description"`
	ExtraDescription string         `db:"extra_description" json:"extraDescription"`
	Images           string         `db:"images" json:"images"`
	ProductionDate   sql.NullString `db:"production_date" json:"productionDate"`
	SeasonID         sql.NullInt64  `db:"season_id" json:"seasonID"`
	Duration         int64          `db:"duration" json:"duration"`
	Number           int64          `db:"number" json:"number"`
}

func (q *Queries) InsertEpisode(ctx context.Context, arg InsertEpisodeParams) error {
	_, err := q.db.ExecContext(ctx, insertEpisode,
		arg.ID,
		arg.LegacyID,
		arg.LegacyProgramID,
		arg.AgeRating,
		arg.Title,
		arg.Description,
		arg.ExtraDescription,
		arg.Images,
		arg.ProductionDate,
		arg.SeasonID,
		arg.Duration,
		arg.Number,
	)
	return err
}
