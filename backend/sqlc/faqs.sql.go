// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: faqs.sql

package sqlc

import (
	"context"

	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
)

const getFAQCategories = `-- name: getFAQCategories :many
WITH t AS (SELECT ts.faq_categories_id,
                  json_object_agg(ts.languages_code, ts.title)       AS title
           FROM faq_categories_translations ts
           GROUP BY ts.faq_categories_id)
SELECT c.id,
       t.title
FROM faq_categories c
         LEFT JOIN t ON c.id = t.faq_categories_id
WHERE c.status = 'published'
  AND c.id = ANY($1::int[])
`

type getFAQCategoriesRow struct {
	ID    int32                 `db:"id" json:"id"`
	Title pqtype.NullRawMessage `db:"title" json:"title"`
}

func (q *Queries) getFAQCategories(ctx context.Context, dollar_1 []int32) ([]getFAQCategoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, getFAQCategories, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getFAQCategoriesRow
	for rows.Next() {
		var i getFAQCategoriesRow
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

const getQuestionIDsForCategories = `-- name: getQuestionIDsForCategories :many
SELECT f.id, f.category
FROM faqs f
         LEFT JOIN faq_categories fc on f.category = fc.id
WHERE f.status = 'published'
  AND fc.status = 'published'
  AND f.category = ANY($1::int[])
`

type getQuestionIDsForCategoriesRow struct {
	ID       int32 `db:"id" json:"id"`
	Category int32 `db:"category" json:"category"`
}

func (q *Queries) getQuestionIDsForCategories(ctx context.Context, dollar_1 []int32) ([]getQuestionIDsForCategoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, getQuestionIDsForCategories, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getQuestionIDsForCategoriesRow
	for rows.Next() {
		var i getQuestionIDsForCategoriesRow
		if err := rows.Scan(&i.ID, &i.Category); err != nil {
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

const getQuestions = `-- name: getQuestions :many
WITH t AS (SELECT ts.faqs_id,
                  json_object_agg(ts.languages_code, ts.question) AS question,
                  json_object_agg(ts.languages_code, ts.answer)   AS answer
           FROM faqs_translations ts
           GROUP BY ts.faqs_id)
SELECT f.id,
       f.category,
       t.question,
       t.answer
FROM faqs f
    LEFT JOIN t ON f.id = t.faqs_id
    LEFT JOIN faq_categories fc on f.category = fc.id
WHERE f.status = 'published'
  AND fc.status = 'published'
  AND f.id = ANY($1::int[])
`

type getQuestionsRow struct {
	ID       int32                 `db:"id" json:"id"`
	Category int32                 `db:"category" json:"category"`
	Question pqtype.NullRawMessage `db:"question" json:"question"`
	Answer   pqtype.NullRawMessage `db:"answer" json:"answer"`
}

func (q *Queries) getQuestions(ctx context.Context, dollar_1 []int32) ([]getQuestionsRow, error) {
	rows, err := q.db.QueryContext(ctx, getQuestions, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getQuestionsRow
	for rows.Next() {
		var i getQuestionsRow
		if err := rows.Scan(
			&i.ID,
			&i.Category,
			&i.Question,
			&i.Answer,
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

const listFAQCategories = `-- name: listFAQCategories :many
WITH t AS (SELECT ts.faq_categories_id,
                  json_object_agg(ts.languages_code, ts.title)    AS title
           FROM faq_categories_translations ts
           GROUP BY ts.faq_categories_id)
SELECT c.id,
       t.title
FROM faq_categories c
         LEFT JOIN t ON c.id = t.faq_categories_id
WHERE c.status = 'published'
`

type listFAQCategoriesRow struct {
	ID    int32                 `db:"id" json:"id"`
	Title pqtype.NullRawMessage `db:"title" json:"title"`
}

func (q *Queries) listFAQCategories(ctx context.Context) ([]listFAQCategoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, listFAQCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listFAQCategoriesRow
	for rows.Next() {
		var i listFAQCategoriesRow
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
