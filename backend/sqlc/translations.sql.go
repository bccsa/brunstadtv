// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: translations.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

const listAlternativeTranslations = `-- name: ListAlternativeTranslations :many
WITH items AS (SELECT i.id
               FROM questionalternatives i)
SELECT ts.id, questionalternatives_id as parent_id, languages_code, title
FROM questionalternatives_translations ts
         JOIN items i ON i.id = ts.questionalternatives_id
WHERE ts.languages_code = ANY ($1::varchar[])
`

type ListAlternativeTranslationsRow struct {
	ID            int32          `db:"id" json:"id"`
	ParentID      uuid.NullUUID  `db:"parent_id" json:"parentID"`
	LanguagesCode string         `db:"languages_code" json:"languagesCode"`
	Title         null_v4.String `db:"title" json:"title"`
}

func (q *Queries) ListAlternativeTranslations(ctx context.Context, dollar_1 []string) ([]ListAlternativeTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listAlternativeTranslations, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAlternativeTranslationsRow
	for rows.Next() {
		var i ListAlternativeTranslationsRow
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.LanguagesCode,
			&i.Title,
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

const listEpisodeTranslations = `-- name: ListEpisodeTranslations :many
WITH episodes AS (SELECT e.id
                  FROM episodes e
                           LEFT JOIN seasons s ON s.id = e.season_id
                           LEFT JOIN shows sh ON sh.id = s.show_id
                  WHERE e.status = 'published'
                    AND s.status = 'published'
                    AND sh.status = 'published')
SELECT et.id, episodes_id as parent_id, languages_code, title, description, extra_description
FROM episodes_translations et
         JOIN episodes e ON e.id = et.episodes_id
WHERE et.languages_code = ANY ($1::varchar[])
`

type ListEpisodeTranslationsRow struct {
	ID               int32          `db:"id" json:"id"`
	ParentID         int32          `db:"parent_id" json:"parentID"`
	LanguagesCode    string         `db:"languages_code" json:"languagesCode"`
	Title            null_v4.String `db:"title" json:"title"`
	Description      null_v4.String `db:"description" json:"description"`
	ExtraDescription null_v4.String `db:"extra_description" json:"extraDescription"`
}

func (q *Queries) ListEpisodeTranslations(ctx context.Context, dollar_1 []string) ([]ListEpisodeTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listEpisodeTranslations, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListEpisodeTranslationsRow
	for rows.Next() {
		var i ListEpisodeTranslationsRow
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.LanguagesCode,
			&i.Title,
			&i.Description,
			&i.ExtraDescription,
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

const listLessonOriginalTranslations = `-- name: ListLessonOriginalTranslations :many
SELECT items.id, items.title, items.description
FROM lessons items
WHERE status = 'published'
`

type ListLessonOriginalTranslationsRow struct {
	ID          uuid.UUID      `db:"id" json:"id"`
	Title       string         `db:"title" json:"title"`
	Description null_v4.String `db:"description" json:"description"`
}

func (q *Queries) ListLessonOriginalTranslations(ctx context.Context) ([]ListLessonOriginalTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listLessonOriginalTranslations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLessonOriginalTranslationsRow
	for rows.Next() {
		var i ListLessonOriginalTranslationsRow
		if err := rows.Scan(&i.ID, &i.Title, &i.Description); err != nil {
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

const listLessonTranslations = `-- name: ListLessonTranslations :many
WITH lessons AS (SELECT s.id
                 FROM lessons s
                 WHERE s.status = 'published')
SELECT st.id, lessons_id as parent_id, languages_code, title, description
FROM lessons_translations st
         JOIN lessons e ON e.id = st.lessons_id
WHERE st.languages_code = ANY ($1::varchar[])
`

type ListLessonTranslationsRow struct {
	ID            int32          `db:"id" json:"id"`
	ParentID      uuid.NullUUID  `db:"parent_id" json:"parentID"`
	LanguagesCode null_v4.String `db:"languages_code" json:"languagesCode"`
	Title         null_v4.String `db:"title" json:"title"`
	Description   null_v4.String `db:"description" json:"description"`
}

func (q *Queries) ListLessonTranslations(ctx context.Context, dollar_1 []string) ([]ListLessonTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listLessonTranslations, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLessonTranslationsRow
	for rows.Next() {
		var i ListLessonTranslationsRow
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.LanguagesCode,
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

const listPageTranslations = `-- name: ListPageTranslations :many
WITH pages AS (SELECT s.id
               FROM pages s
               WHERE s.status = 'published')
SELECT st.id, pages_id as parent_id, languages_code, title, description
FROM pages_translations st
         JOIN pages e ON e.id = st.pages_id
WHERE st.languages_code = ANY ($1::varchar[])
`

type ListPageTranslationsRow struct {
	ID            int32          `db:"id" json:"id"`
	ParentID      null_v4.Int    `db:"parent_id" json:"parentID"`
	LanguagesCode null_v4.String `db:"languages_code" json:"languagesCode"`
	Title         null_v4.String `db:"title" json:"title"`
	Description   null_v4.String `db:"description" json:"description"`
}

func (q *Queries) ListPageTranslations(ctx context.Context, dollar_1 []string) ([]ListPageTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listPageTranslations, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPageTranslationsRow
	for rows.Next() {
		var i ListPageTranslationsRow
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.LanguagesCode,
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

const listQuestionAlternativesOriginalTranslations = `-- name: ListQuestionAlternativesOriginalTranslations :many
SELECT items.id, items.title
FROM questionalternatives items
         JOIN tasks t ON t.id = items.task_id
WHERE t.status = 'published'
`

type ListQuestionAlternativesOriginalTranslationsRow struct {
	ID    uuid.UUID      `db:"id" json:"id"`
	Title null_v4.String `db:"title" json:"title"`
}

func (q *Queries) ListQuestionAlternativesOriginalTranslations(ctx context.Context) ([]ListQuestionAlternativesOriginalTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listQuestionAlternativesOriginalTranslations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListQuestionAlternativesOriginalTranslationsRow
	for rows.Next() {
		var i ListQuestionAlternativesOriginalTranslationsRow
		if err := rows.Scan(&i.ID, &i.Title); err != nil {
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

const listSeasonTranslations = `-- name: ListSeasonTranslations :many
WITH seasons AS (SELECT s.id
                 FROM seasons s
                          LEFT JOIN shows sh ON sh.id = s.show_id
                 WHERE s.status = 'published'
                   AND sh.status = 'published')
SELECT et.id, seasons_id as parent_id, languages_code, title, description
FROM seasons_translations et
         JOIN seasons e ON e.id = et.seasons_id
WHERE et.languages_code = ANY ($1::varchar[])
`

type ListSeasonTranslationsRow struct {
	ID            int32          `db:"id" json:"id"`
	ParentID      int32          `db:"parent_id" json:"parentID"`
	LanguagesCode string         `db:"languages_code" json:"languagesCode"`
	Title         null_v4.String `db:"title" json:"title"`
	Description   null_v4.String `db:"description" json:"description"`
}

func (q *Queries) ListSeasonTranslations(ctx context.Context, dollar_1 []string) ([]ListSeasonTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listSeasonTranslations, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSeasonTranslationsRow
	for rows.Next() {
		var i ListSeasonTranslationsRow
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.LanguagesCode,
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

const listSectionTranslations = `-- name: ListSectionTranslations :many
WITH sections AS (SELECT s.id
                  FROM sections s
                  WHERE s.status = 'published'
                    AND s.show_title = true)
SELECT st.id, sections_id as parent_id, languages_code, title, description
FROM sections_translations st
         JOIN sections e ON e.id = st.sections_id
WHERE st.languages_code = ANY ($1::varchar[])
`

type ListSectionTranslationsRow struct {
	ID            int32          `db:"id" json:"id"`
	ParentID      int32          `db:"parent_id" json:"parentID"`
	LanguagesCode string         `db:"languages_code" json:"languagesCode"`
	Title         null_v4.String `db:"title" json:"title"`
	Description   null_v4.String `db:"description" json:"description"`
}

func (q *Queries) ListSectionTranslations(ctx context.Context, dollar_1 []string) ([]ListSectionTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listSectionTranslations, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSectionTranslationsRow
	for rows.Next() {
		var i ListSectionTranslationsRow
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.LanguagesCode,
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

const listShowTranslations = `-- name: ListShowTranslations :many
WITH shows AS (SELECT s.id
               FROM shows s
               WHERE s.status = 'published')
SELECT et.id, shows_id as parent_id, languages_code, title, description
FROM shows_translations et
         JOIN shows e ON e.id = et.shows_id
WHERE et.languages_code = ANY ($1::varchar[])
`

type ListShowTranslationsRow struct {
	ID            int32          `db:"id" json:"id"`
	ParentID      int32          `db:"parent_id" json:"parentID"`
	LanguagesCode string         `db:"languages_code" json:"languagesCode"`
	Title         null_v4.String `db:"title" json:"title"`
	Description   null_v4.String `db:"description" json:"description"`
}

func (q *Queries) ListShowTranslations(ctx context.Context, dollar_1 []string) ([]ListShowTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listShowTranslations, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListShowTranslationsRow
	for rows.Next() {
		var i ListShowTranslationsRow
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.LanguagesCode,
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

const listStudyTopicOriginalTranslations = `-- name: ListStudyTopicOriginalTranslations :many
SELECT items.id, items.title, items.description
FROM studytopics items
WHERE status = 'published'
`

type ListStudyTopicOriginalTranslationsRow struct {
	ID          uuid.UUID      `db:"id" json:"id"`
	Title       string         `db:"title" json:"title"`
	Description null_v4.String `db:"description" json:"description"`
}

func (q *Queries) ListStudyTopicOriginalTranslations(ctx context.Context) ([]ListStudyTopicOriginalTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listStudyTopicOriginalTranslations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListStudyTopicOriginalTranslationsRow
	for rows.Next() {
		var i ListStudyTopicOriginalTranslationsRow
		if err := rows.Scan(&i.ID, &i.Title, &i.Description); err != nil {
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

const listStudyTopicTranslations = `-- name: ListStudyTopicTranslations :many
WITH items AS (SELECT i.id
               FROM studytopics i
               WHERE i.status = 'published')
SELECT ts.id, studytopics_id as parent_id, languages_code, title, description
FROM studytopics_translations ts
         JOIN items i ON i.id = ts.studytopics_id
WHERE ts.languages_code = ANY ($1::varchar[])
`

type ListStudyTopicTranslationsRow struct {
	ID            int32          `db:"id" json:"id"`
	ParentID      uuid.NullUUID  `db:"parent_id" json:"parentID"`
	LanguagesCode null_v4.String `db:"languages_code" json:"languagesCode"`
	Title         null_v4.String `db:"title" json:"title"`
	Description   null_v4.String `db:"description" json:"description"`
}

func (q *Queries) ListStudyTopicTranslations(ctx context.Context, dollar_1 []string) ([]ListStudyTopicTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listStudyTopicTranslations, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListStudyTopicTranslationsRow
	for rows.Next() {
		var i ListStudyTopicTranslationsRow
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.LanguagesCode,
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

const listTaskOriginalTranslations = `-- name: ListTaskOriginalTranslations :many
SELECT items.id, items.title, items.description
FROM tasks items
WHERE status = 'published'
`

type ListTaskOriginalTranslationsRow struct {
	ID          uuid.UUID      `db:"id" json:"id"`
	Title       null_v4.String `db:"title" json:"title"`
	Description null_v4.String `db:"description" json:"description"`
}

func (q *Queries) ListTaskOriginalTranslations(ctx context.Context) ([]ListTaskOriginalTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listTaskOriginalTranslations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListTaskOriginalTranslationsRow
	for rows.Next() {
		var i ListTaskOriginalTranslationsRow
		if err := rows.Scan(&i.ID, &i.Title, &i.Description); err != nil {
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

const listTaskTranslations = `-- name: ListTaskTranslations :many
WITH items AS (SELECT i.id
               FROM tasks i
               WHERE i.status = 'published')
SELECT ts.id, tasks_id as parent_id, languages_code, title, description
FROM tasks_translations ts
         JOIN items i ON i.id = ts.tasks_id
WHERE ts.languages_code = ANY ($1::varchar[])
`

type ListTaskTranslationsRow struct {
	ID            int32          `db:"id" json:"id"`
	ParentID      uuid.NullUUID  `db:"parent_id" json:"parentID"`
	LanguagesCode null_v4.String `db:"languages_code" json:"languagesCode"`
	Title         null_v4.String `db:"title" json:"title"`
	Description   null_v4.String `db:"description" json:"description"`
}

func (q *Queries) ListTaskTranslations(ctx context.Context, dollar_1 []string) ([]ListTaskTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listTaskTranslations, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListTaskTranslationsRow
	for rows.Next() {
		var i ListTaskTranslationsRow
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.LanguagesCode,
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
