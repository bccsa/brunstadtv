// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: surveys.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/sqlc-dev/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getPromptIDsForRoles = `-- name: GetPromptIDsForRoles :many
WITH roles AS (SELECT pt.prompts_id,
                      array_agg(u.usergroups_code) AS roles
               FROM prompts_targets pt
                        LEFT JOIN targets_usergroups u ON u.targets_id = pt.targets_id
               GROUP BY pt.prompts_id)
SELECT p.id
FROM prompts p
         LEFT JOIN roles ON roles.prompts_id = p.id
WHERE p.status = 'published'
  AND p.from < (NOW() + interval '7 day')
  AND p.to > NOW()
  AND roles.roles && $1::varchar[]
`

func (q *Queries) GetPromptIDsForRoles(ctx context.Context, roles []string) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getPromptIDsForRoles, pq.Array(roles))
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

const getPrompts = `-- name: GetPrompts :many
WITH ts AS (SELECT ts.prompts_id                                       AS id,
                   json_object_agg(languages_code, ts.title)           AS title,
                   json_object_agg(languages_code, ts.secondary_title) AS secondary_title
            FROM prompts_translations ts
            GROUP BY ts.prompts_id)
SELECT p.id,
       p.title           as original_title,
       p.secondary_title as original_secondary_title,
       p.from,
       p.to,
       p.type,
       p.survey_id,
       ts.title,
       ts.secondary_title
FROM prompts p
         LEFT JOIN ts ON ts.id = p.id
WHERE p.id = ANY ($1::uuid[])
`

type GetPromptsRow struct {
	ID                     uuid.UUID             `db:"id" json:"id"`
	OriginalTitle          string                `db:"original_title" json:"originalTitle"`
	OriginalSecondaryTitle null_v4.String        `db:"original_secondary_title" json:"originalSecondaryTitle"`
	From                   time.Time             `db:"from" json:"from"`
	To                     time.Time             `db:"to" json:"to"`
	Type                   string                `db:"type" json:"type"`
	SurveyID               uuid.NullUUID         `db:"survey_id" json:"surveyId"`
	Title                  pqtype.NullRawMessage `db:"title" json:"title"`
	SecondaryTitle         pqtype.NullRawMessage `db:"secondary_title" json:"secondaryTitle"`
}

func (q *Queries) GetPrompts(ctx context.Context, ids []uuid.UUID) ([]GetPromptsRow, error) {
	rows, err := q.db.QueryContext(ctx, getPrompts, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPromptsRow
	for rows.Next() {
		var i GetPromptsRow
		if err := rows.Scan(
			&i.ID,
			&i.OriginalTitle,
			&i.OriginalSecondaryTitle,
			&i.From,
			&i.To,
			&i.Type,
			&i.SurveyID,
			&i.Title,
			&i.SecondaryTitle,
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

const upsertSurveyAnswer = `-- name: UpsertSurveyAnswer :exec
INSERT INTO users.surveyquestionanswers (profile_id, question_id, updated_at)
VALUES ($1::uuid, $2::uuid, now())
ON CONFLICT(profile_id, question_id) DO UPDATE SET updated_at = EXCLUDED.updated_at
`

type UpsertSurveyAnswerParams struct {
	ProfileID  uuid.UUID `db:"profile_id" json:"profileId"`
	QuestionID uuid.UUID `db:"question_id" json:"questionId"`
}

func (q *Queries) UpsertSurveyAnswer(ctx context.Context, arg UpsertSurveyAnswerParams) error {
	_, err := q.db.ExecContext(ctx, upsertSurveyAnswer, arg.ProfileID, arg.QuestionID)
	return err
}

const getQuestionIDsForSurveyIDs = `-- name: getQuestionIDsForSurveyIDs :many
SELECT q.id, q.survey_id AS parent_id
FROM surveyquestions q
WHERE q.survey_id = ANY ($1::uuid[])
ORDER BY q.sort
`

type getQuestionIDsForSurveyIDsRow struct {
	ID       uuid.UUID `db:"id" json:"id"`
	ParentID uuid.UUID `db:"parent_id" json:"parentId"`
}

func (q *Queries) getQuestionIDsForSurveyIDs(ctx context.Context, ids []uuid.UUID) ([]getQuestionIDsForSurveyIDsRow, error) {
	rows, err := q.db.QueryContext(ctx, getQuestionIDsForSurveyIDs, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getQuestionIDsForSurveyIDsRow
	for rows.Next() {
		var i getQuestionIDsForSurveyIDsRow
		if err := rows.Scan(&i.ID, &i.ParentID); err != nil {
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

const getSurveyQuestions = `-- name: getSurveyQuestions :many
WITH ts AS (SELECT ts.surveyquestions_id                           AS id,
                   json_object_agg(languages_code, ts.title)       AS title,
                   json_object_agg(languages_code, ts.description) AS description
            FROM surveyquestions_translations ts
            GROUP BY ts.surveyquestions_id)
SELECT s.id,
       s.title       AS original_title,
       s.description AS original_description,
       s.placeholder AS original_placeholder,
       s.survey_id,
       s.type,
       ts.title,
       ts.description
FROM surveyquestions s
         LEFT JOIN ts ON ts.id = s.id
WHERE s.id = ANY ($1::uuid[])
`

type getSurveyQuestionsRow struct {
	ID                  uuid.UUID             `db:"id" json:"id"`
	OriginalTitle       string                `db:"original_title" json:"originalTitle"`
	OriginalDescription null_v4.String        `db:"original_description" json:"originalDescription"`
	OriginalPlaceholder null_v4.String        `db:"original_placeholder" json:"originalPlaceholder"`
	SurveyID            uuid.UUID             `db:"survey_id" json:"surveyId"`
	Type                string                `db:"type" json:"type"`
	Title               pqtype.NullRawMessage `db:"title" json:"title"`
	Description         pqtype.NullRawMessage `db:"description" json:"description"`
}

func (q *Queries) getSurveyQuestions(ctx context.Context, ids []uuid.UUID) ([]getSurveyQuestionsRow, error) {
	rows, err := q.db.QueryContext(ctx, getSurveyQuestions, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getSurveyQuestionsRow
	for rows.Next() {
		var i getSurveyQuestionsRow
		if err := rows.Scan(
			&i.ID,
			&i.OriginalTitle,
			&i.OriginalDescription,
			&i.OriginalPlaceholder,
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
