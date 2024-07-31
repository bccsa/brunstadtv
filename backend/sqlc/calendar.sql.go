// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: calendar.sql

package sqlc

import (
	"context"
	"encoding/json"
	"time"

	"github.com/lib/pq"
	"github.com/sqlc-dev/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getCalendarEntries = `-- name: getCalendarEntries :many
SELECT e.id,
       e.event_id,
       e.link_type,
       e.start,
       e.end,
       COALESCE(e.is_replay, false) = true AS is_replay,
       ea.id                               AS episode_id,
       se.id                               AS season_id,
       sh.id                               AS show_id,
       ts.title,
       ts.description
FROM calendarentries e
         LEFT JOIN LATERAL (SELECT json_object_agg(ts.languages_code, ts.title)       AS title,
                                   json_object_agg(ts.languages_code, ts.description) AS description
                            FROM calendarentries_translations ts
                            WHERE ts.calendarentries_id = e.id) ts ON true
         LEFT JOIN episode_roles er ON er.id = e.episode_id AND er.roles && $2::varchar[]
         LEFT JOIN episode_availability ea ON ea.id = er.id AND ea.published
         LEFT JOIN seasons se ON se.id = e.season_id AND se.status = 'published'
         LEFT JOIN shows sh ON sh.id = e.show_id AND sh.status = 'published'
WHERE e.status = 'published'
  AND e.id = ANY ($1::int[])
`

type getCalendarEntriesParams struct {
	Column1 []int32  `db:"column_1" json:"column1"`
	Column2 []string `db:"column_2" json:"column2"`
}

type getCalendarEntriesRow struct {
	ID          int32           `db:"id" json:"id"`
	EventID     null_v4.Int     `db:"event_id" json:"eventId"`
	LinkType    null_v4.String  `db:"link_type" json:"linkType"`
	Start       time.Time       `db:"start" json:"start"`
	End         time.Time       `db:"end" json:"end"`
	IsReplay    bool            `db:"is_replay" json:"isReplay"`
	EpisodeID   null_v4.Int     `db:"episode_id" json:"episodeId"`
	SeasonID    null_v4.Int     `db:"season_id" json:"seasonId"`
	ShowID      null_v4.Int     `db:"show_id" json:"showId"`
	Title       json.RawMessage `db:"title" json:"title"`
	Description json.RawMessage `db:"description" json:"description"`
}

func (q *Queries) getCalendarEntries(ctx context.Context, arg getCalendarEntriesParams) ([]getCalendarEntriesRow, error) {
	rows, err := q.db.QueryContext(ctx, getCalendarEntries, pq.Array(arg.Column1), pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getCalendarEntriesRow
	for rows.Next() {
		var i getCalendarEntriesRow
		if err := rows.Scan(
			&i.ID,
			&i.EventID,
			&i.LinkType,
			&i.Start,
			&i.End,
			&i.IsReplay,
			&i.EpisodeID,
			&i.SeasonID,
			&i.ShowID,
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

const getCalendarEntriesByID = `-- name: getCalendarEntriesByID :many
SELECT e.id,
       e.event_id,
       e.link_type,
       e.start,
       e.end,
       COALESCE(e.is_replay, false) = true AS is_replay,
       ea.id                               AS episode_id,
       se.id                               AS season_id,
       sh.id                               AS show_id,
       ts.title,
       ts.description
FROM calendarentries e
         LEFT JOIN LATERAL (SELECT json_object_agg(ts.languages_code, ts.title)       AS title,
                                   json_object_agg(ts.languages_code, ts.description) AS description
                            FROM calendarentries_translations ts
                            WHERE ts.calendarentries_id = e.id) ts ON true
         LEFT JOIN episode_roles er ON er.id = e.episode_id
         LEFT JOIN episode_availability ea ON ea.id = er.id
         LEFT JOIN seasons se ON se.id = e.season_id
         LEFT JOIN shows sh ON sh.id = e.show_id
WHERE e.id = ANY ($1::int[])
`

type getCalendarEntriesByIDRow struct {
	ID          int32           `db:"id" json:"id"`
	EventID     null_v4.Int     `db:"event_id" json:"eventId"`
	LinkType    null_v4.String  `db:"link_type" json:"linkType"`
	Start       time.Time       `db:"start" json:"start"`
	End         time.Time       `db:"end" json:"end"`
	IsReplay    bool            `db:"is_replay" json:"isReplay"`
	EpisodeID   null_v4.Int     `db:"episode_id" json:"episodeId"`
	SeasonID    null_v4.Int     `db:"season_id" json:"seasonId"`
	ShowID      null_v4.Int     `db:"show_id" json:"showId"`
	Title       json.RawMessage `db:"title" json:"title"`
	Description json.RawMessage `db:"description" json:"description"`
}

func (q *Queries) getCalendarEntriesByID(ctx context.Context, dollar_1 []int32) ([]getCalendarEntriesByIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getCalendarEntriesByID, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getCalendarEntriesByIDRow
	for rows.Next() {
		var i getCalendarEntriesByIDRow
		if err := rows.Scan(
			&i.ID,
			&i.EventID,
			&i.LinkType,
			&i.Start,
			&i.End,
			&i.IsReplay,
			&i.EpisodeID,
			&i.SeasonID,
			&i.ShowID,
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

const getCalendarEntryIDsForEvents = `-- name: getCalendarEntryIDsForEvents :many
SELECT e.id, e.event_id as parent_id
FROM calendarentries e
WHERE e.status = 'published'
  AND e.event_id = ANY ($1::int[])
ORDER BY e.start
`

type getCalendarEntryIDsForEventsRow struct {
	ID       int32       `db:"id" json:"id"`
	ParentID null_v4.Int `db:"parent_id" json:"parentId"`
}

func (q *Queries) getCalendarEntryIDsForEvents(ctx context.Context, ids []int32) ([]getCalendarEntryIDsForEventsRow, error) {
	rows, err := q.db.QueryContext(ctx, getCalendarEntryIDsForEvents, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getCalendarEntryIDsForEventsRow
	for rows.Next() {
		var i getCalendarEntryIDsForEventsRow
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

const getCalendarEntryIDsForPeriod = `-- name: getCalendarEntryIDsForPeriod :many
SELECT e.id
FROM calendarentries e
WHERE e.status = 'published'
  AND (e.start >= $1::timestamptz AND e.start <= $2::timestamptz)
ORDER BY e.start
`

type getCalendarEntryIDsForPeriodParams struct {
	Column1 time.Time `db:"column_1" json:"column1"`
	Column2 time.Time `db:"column_2" json:"column2"`
}

func (q *Queries) getCalendarEntryIDsForPeriod(ctx context.Context, arg getCalendarEntryIDsForPeriodParams) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, getCalendarEntryIDsForPeriod, arg.Column1, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
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

const getEventIDsForPeriod = `-- name: getEventIDsForPeriod :many
SELECT e.id
FROM events e
WHERE e.status = 'published'
  AND ((e.start >= $1::timestamptz AND e.start <= $2::timestamptz) OR
       (e.end >= $1::timestamptz AND e.end <= $2::timestamptz))
ORDER BY e.start
`

type getEventIDsForPeriodParams struct {
	Column1 time.Time `db:"column_1" json:"column1"`
	Column2 time.Time `db:"column_2" json:"column2"`
}

func (q *Queries) getEventIDsForPeriod(ctx context.Context, arg getEventIDsForPeriodParams) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, getEventIDsForPeriod, arg.Column1, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
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

const getEvents = `-- name: getEvents :many
WITH t AS (SELECT ts.events_id,
                  json_object_agg(ts.languages_code, ts.title) AS title
           FROM events_translations ts
           GROUP BY ts.events_id)
SELECT e.id,
       e.start,
       e.end,
       t.title
FROM events e
         LEFT JOIN t ON e.id = t.events_id
WHERE e.status = 'published'
  AND e.id = ANY ($1::int[])
`

type getEventsRow struct {
	ID    int32                 `db:"id" json:"id"`
	Start time.Time             `db:"start" json:"start"`
	End   time.Time             `db:"end" json:"end"`
	Title pqtype.NullRawMessage `db:"title" json:"title"`
}

func (q *Queries) getEvents(ctx context.Context, dollar_1 []int32) ([]getEventsRow, error) {
	rows, err := q.db.QueryContext(ctx, getEvents, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getEventsRow
	for rows.Next() {
		var i getEventsRow
		if err := rows.Scan(
			&i.ID,
			&i.Start,
			&i.End,
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

const listEvents = `-- name: listEvents :many
WITH t AS (SELECT ts.events_id,
                  json_object_agg(ts.languages_code, ts.title) AS title
           FROM events_translations ts
           GROUP BY ts.events_id)
SELECT e.id,
       e.start,
       e.end,
       t.title
FROM events e
         LEFT JOIN t ON e.id = t.events_id
WHERE e.status = 'published'
  AND e.end >= now() - '1 year'::interval
ORDER BY e.start
`

type listEventsRow struct {
	ID    int32                 `db:"id" json:"id"`
	Start time.Time             `db:"start" json:"start"`
	End   time.Time             `db:"end" json:"end"`
	Title pqtype.NullRawMessage `db:"title" json:"title"`
}

func (q *Queries) listEvents(ctx context.Context) ([]listEventsRow, error) {
	rows, err := q.db.QueryContext(ctx, listEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listEventsRow
	for rows.Next() {
		var i listEventsRow
		if err := rows.Scan(
			&i.ID,
			&i.Start,
			&i.End,
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
