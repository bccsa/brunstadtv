// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: surveys.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getSurveyIDFromQuestionID = `-- name: GetSurveyIDFromQuestionID :one
SELECT q.survey_id
FROM surveyquestions q
WHERE q.id = $1::uuid
`

func (q *Queries) GetSurveyIDFromQuestionID(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getSurveyIDFromQuestionID, id)
	var survey_id uuid.UUID
	err := row.Scan(&survey_id)
	return survey_id, err
}

const getSurveyIDsForRoles = `-- name: GetSurveyIDsForRoles :many
WITH roles AS (SELECT st.surveys_id,
                      array_agg(u.usergroups_code) AS roles
               FROM surveys_targets st
                        LEFT JOIN targets t ON st.targets_id = st.targets_id AND t.type = 'usergroups'
                        LEFT JOIN targets_usergroups u ON u.targets_id = t.id
               GROUP BY st.surveys_id)
SELECT s.id
FROM surveys s
         LEFT JOIN roles ON roles.surveys_id = s.id
WHERE s.from < (NOW() + interval '7 day')
  AND s.to > NOW()
  AND roles.roles && $1::varchar[]
`

func (q *Queries) GetSurveyIDsForRoles(ctx context.Context, roles []string) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getSurveyIDsForRoles, pq.Array(roles))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
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

const getSurveyQuestionsForSurveyIDs = `-- name: getSurveyQuestionsForSurveyIDs :many
WITH ts AS (SELECT ts.surveyquestions_id                           AS id,
                   json_object_agg(languages_code, ts.title)       AS title,
                   json_object_agg(languages_code, ts.description) AS description
            FROM surveyquestions_translations ts
            GROUP BY ts.surveyquestions_id)
SELECT s.id,
       s.title       AS original_title,
       s.description AS original_description,
       s.survey_id,
       s.type,
       ts.title,
       ts.description
FROM surveyquestions s
         LEFT JOIN ts ON ts.id = s.id
WHERE s.survey_id = ANY ($1::uuid[])
`

type getSurveyQuestionsForSurveyIDsRow struct {
	ID                  uuid.UUID             `db:"id" json:"id"`
	OriginalTitle       string                `db:"original_title" json:"originalTitle"`
	OriginalDescription null_v4.String        `db:"original_description" json:"originalDescription"`
	SurveyID            uuid.UUID             `db:"survey_id" json:"surveyID"`
	Type                string                `db:"type" json:"type"`
	Title               pqtype.NullRawMessage `db:"title" json:"title"`
	Description         pqtype.NullRawMessage `db:"description" json:"description"`
}

func (q *Queries) getSurveyQuestionsForSurveyIDs(ctx context.Context, ids []uuid.UUID) ([]getSurveyQuestionsForSurveyIDsRow, error) {
	rows, err := q.db.QueryContext(ctx, getSurveyQuestionsForSurveyIDs, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getSurveyQuestionsForSurveyIDsRow
	for rows.Next() {
		var i getSurveyQuestionsForSurveyIDsRow
		if err := rows.Scan(
			&i.ID,
			&i.OriginalTitle,
			&i.OriginalDescription,
			&i.SurveyID,
			&i.Type,
			&i.Title,
			&i.Description,
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

const getSurveys = `-- name: getSurveys :many
WITH ts AS (SELECT ts.surveys_id                                   AS id,
                   json_object_agg(languages_code, ts.title)       AS title,
                   json_object_agg(languages_code, ts.description) AS description
            FROM surveys_translations ts
            GROUP BY ts.surveys_id)
SELECT s.id,
       s.title       AS original_title,
       s.description AS original_description,
       s.from,
       s.to,
       ts.title,
       ts.description
FROM surveys s
         LEFT JOIN ts ON ts.id = s.id
WHERE s.id = ANY ($1::uuid[])
`

type getSurveysRow struct {
	ID                  uuid.UUID             `db:"id" json:"id"`
	OriginalTitle       string                `db:"original_title" json:"originalTitle"`
	OriginalDescription null_v4.String        `db:"original_description" json:"originalDescription"`
	From                time.Time             `db:"from" json:"from"`
	To                  time.Time             `db:"to" json:"to"`
	Title               pqtype.NullRawMessage `db:"title" json:"title"`
	Description         pqtype.NullRawMessage `db:"description" json:"description"`
}

func (q *Queries) getSurveys(ctx context.Context, ids []uuid.UUID) ([]getSurveysRow, error) {
	rows, err := q.db.QueryContext(ctx, getSurveys, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getSurveysRow
	for rows.Next() {
		var i getSurveysRow
		if err := rows.Scan(
			&i.ID,
			&i.OriginalTitle,
			&i.OriginalDescription,
			&i.From,
			&i.To,
			&i.Title,
			&i.Description,
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
