// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: collections.sql

package sqlc

import (
	"context"
	"time"

	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getCollectionEntriesForCollections = `-- name: getCollectionEntriesForCollections :many
SELECT id, collections_id, item, collection, sort
FROM collections_entries ci
WHERE ci.collections_id = ANY ($1::int[])
`

func (q *Queries) getCollectionEntriesForCollections(ctx context.Context, dollar_1 []int32) ([]CollectionsEntry, error) {
	rows, err := q.db.QueryContext(ctx, getCollectionEntriesForCollections, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CollectionsEntry
	for rows.Next() {
		var i CollectionsEntry
		if err := rows.Scan(
			&i.ID,
			&i.CollectionsID,
			&i.Item,
			&i.Collection,
			&i.Sort,
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

const getCollectionEntriesForCollectionsWithRoles = `-- name: getCollectionEntriesForCollectionsWithRoles :many
SELECT ce.id, ce.collections_id, ce.item, ce.collection, ce.sort
FROM collections_entries ce
         LEFT JOIN episode_roles er ON ce.collection = 'episodes' AND er.id::varchar = ce.item
         LEFT JOIN episode_availability ea ON ce.collection = 'episodes' AND ea.id::varchar = ce.item
         LEFT JOIN season_roles sr ON ce.collection = 'seasons' AND sr.id::varchar = ce.item
         LEFT JOIN season_availability sa ON ce.collection = 'seasons' AND sa.id::varchar = ce.item
         LEFT JOIN show_roles shr ON ce.collection = 'shows' AND shr.id::varchar = ce.item
         LEFT JOIN show_availability sha ON ce.collection = 'shows' AND sha.id::varchar = ce.item
WHERE ce.collections_id = ANY ($1::int[])
  AND (ce.collection != 'episodes' OR (
        ea.published
        AND ea.available_to > now()
        AND er.roles && $2::varchar[] AND ea.available_from < now()
    ))
  AND ce.collection != 'seasons'
  AND (ce.collection != 'shows' OR (
        sha.published
        AND sha.available_to > now()
        AND shr.roles && $2::varchar[] AND sha.available_from < now()
    ))
ORDER BY ce.sort
`

type getCollectionEntriesForCollectionsWithRolesParams struct {
	Column1 []int32  `db:"column_1" json:"column1"`
	Column2 []string `db:"column_2" json:"column2"`
}

func (q *Queries) getCollectionEntriesForCollectionsWithRoles(ctx context.Context, arg getCollectionEntriesForCollectionsWithRolesParams) ([]CollectionsEntry, error) {
	rows, err := q.db.QueryContext(ctx, getCollectionEntriesForCollectionsWithRoles, pq.Array(arg.Column1), pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CollectionsEntry
	for rows.Next() {
		var i CollectionsEntry
		if err := rows.Scan(
			&i.ID,
			&i.CollectionsID,
			&i.Item,
			&i.Collection,
			&i.Sort,
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

const getCollectionIDsForCodes = `-- name: getCollectionIDsForCodes :many
SELECT c.id, ct.slug
FROM collections c
         JOIN collections_translations ct ON c.id = ct.collections_id AND ct.slug = ANY ($1::varchar[])
`

type getCollectionIDsForCodesRow struct {
	ID   int32          `db:"id" json:"id"`
	Slug null_v4.String `db:"slug" json:"slug"`
}

func (q *Queries) getCollectionIDsForCodes(ctx context.Context, dollar_1 []string) ([]getCollectionIDsForCodesRow, error) {
	rows, err := q.db.QueryContext(ctx, getCollectionIDsForCodes, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getCollectionIDsForCodesRow
	for rows.Next() {
		var i getCollectionIDsForCodesRow
		if err := rows.Scan(&i.ID, &i.Slug); err != nil {
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

const getCollections = `-- name: getCollections :many
WITH ts AS (SELECT collections_id,
                   json_object_agg(languages_code, title) AS title,
                   json_object_agg(languages_code, slug)  AS slugs
            FROM collections_translations
            GROUP BY collections_id)
SELECT c.id,
       c.advanced_type,
       c.date_created,
       c.date_updated,
       c.filter_type,
       c.query_filter,
       ts.title,
       ts.slugs
FROM collections c
         LEFT JOIN ts ON ts.collections_id = c.id
WHERE c.id = ANY ($1::int[])
`

type getCollectionsRow struct {
	ID           int32                 `db:"id" json:"id"`
	AdvancedType null_v4.String        `db:"advanced_type" json:"advancedType"`
	DateCreated  time.Time             `db:"date_created" json:"dateCreated"`
	DateUpdated  time.Time             `db:"date_updated" json:"dateUpdated"`
	FilterType   null_v4.String        `db:"filter_type" json:"filterType"`
	QueryFilter  pqtype.NullRawMessage `db:"query_filter" json:"queryFilter"`
	Title        pqtype.NullRawMessage `db:"title" json:"title"`
	Slugs        pqtype.NullRawMessage `db:"slugs" json:"slugs"`
}

func (q *Queries) getCollections(ctx context.Context, dollar_1 []int32) ([]getCollectionsRow, error) {
	rows, err := q.db.QueryContext(ctx, getCollections, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getCollectionsRow
	for rows.Next() {
		var i getCollectionsRow
		if err := rows.Scan(
			&i.ID,
			&i.AdvancedType,
			&i.DateCreated,
			&i.DateUpdated,
			&i.FilterType,
			&i.QueryFilter,
			&i.Title,
			&i.Slugs,
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
