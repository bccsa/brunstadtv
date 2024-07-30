// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: relations.sql

package sqlc

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getCollectionIDFromKey = `-- name: GetCollectionIDFromKey :one
SELECT id
FROM "public"."songcollections"
WHERE key = $1
LIMIT 1
`

func (q *Queries) GetCollectionIDFromKey(ctx context.Context, key string) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getCollectionIDFromKey, key)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getCollectionSongID = `-- name: GetCollectionSongID :one
SELECT s.id
FROM "public"."songs" s
         JOIN "public"."songcollections" c ON s.collection_id = c.id
WHERE c.key = $1
  AND s.key = $2
LIMIT 1
`

type GetCollectionSongIDParams struct {
	CollectionKey string `db:"collection_key" json:"collectionKey"`
	SongKey       string `db:"song_key" json:"songKey"`
}

func (q *Queries) GetCollectionSongID(ctx context.Context, arg GetCollectionSongIDParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getCollectionSongID, arg.CollectionKey, arg.SongKey)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getPersonIDsByNames = `-- name: GetPersonIDsByNames :many
SELECT p.id, p.name
FROM persons p
WHERE name = ANY ($1::varchar[])
`

func (q *Queries) GetPersonIDsByNames(ctx context.Context, names []string) ([]Person, error) {
	rows, err := q.db.QueryContext(ctx, getPersonIDsByNames, pq.Array(names))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Person
	for rows.Next() {
		var i Person
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const insertPerson = `-- name: InsertPerson :exec
INSERT INTO "public"."persons" (id, name)
VALUES ($1, $2)
`

type InsertPersonParams struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
}

func (q *Queries) InsertPerson(ctx context.Context, arg InsertPersonParams) error {
	_, err := q.db.ExecContext(ctx, insertPerson, arg.ID, arg.Name)
	return err
}

const insertSong = `-- name: InsertSong :exec
INSERT INTO "public"."songs" (id, key, collection_id, title)
VALUES ($1, $2, $3, $4)
`

type InsertSongParams struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Key          string    `db:"key" json:"key"`
	CollectionID uuid.UUID `db:"collection_id" json:"collectionId"`
	Title        string    `db:"title" json:"title"`
}

func (q *Queries) InsertSong(ctx context.Context, arg InsertSongParams) error {
	_, err := q.db.ExecContext(ctx, insertSong,
		arg.ID,
		arg.Key,
		arg.CollectionID,
		arg.Title,
	)
	return err
}

const insertSongCollection = `-- name: InsertSongCollection :exec
INSERT INTO "public"."songcollections" (id, key, title)
VALUES ($1, $2, $3)
`

type InsertSongCollectionParams struct {
	ID    uuid.UUID      `db:"id" json:"id"`
	Key   string         `db:"key" json:"key"`
	Title null_v4.String `db:"title" json:"title"`
}

func (q *Queries) InsertSongCollection(ctx context.Context, arg InsertSongCollectionParams) error {
	_, err := q.db.ExecContext(ctx, insertSongCollection, arg.ID, arg.Key, arg.Title)
	return err
}

const insertTimedMetadataPerson = `-- name: InsertTimedMetadataPerson :exec
INSERT INTO "public"."timedmetadata_persons" (timedmetadata_id, persons_id)
VALUES ($1::uuid, $2::uuid)
`

type InsertTimedMetadataPersonParams struct {
	TimedmetadataID uuid.UUID `db:"timedmetadata_id" json:"timedmetadataId"`
	PersonsID       uuid.UUID `db:"persons_id" json:"personsId"`
}

func (q *Queries) InsertTimedMetadataPerson(ctx context.Context, arg InsertTimedMetadataPersonParams) error {
	_, err := q.db.ExecContext(ctx, insertTimedMetadataPerson, arg.TimedmetadataID, arg.PersonsID)
	return err
}

const getPersons = `-- name: getPersons :many
SELECT id, name
FROM persons
WHERE id = ANY ($1::uuid[])
`

func (q *Queries) getPersons(ctx context.Context, ids []uuid.UUID) ([]Person, error) {
	rows, err := q.db.QueryContext(ctx, getPersons, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Person
	for rows.Next() {
		var i Person
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const getPhrases = `-- name: getPhrases :many
SELECT p.key,
       p.value                                   AS original_value,
       COALESCE((SELECT json_object_agg(value, languages_code)
                 FROM phrases_translations
                 WHERE key = p.key), '{}')::json AS value
FROM phrases p
WHERE key = ANY ($1::varchar[])
`

type getPhrasesRow struct {
	Key           string          `db:"key" json:"key"`
	OriginalValue string          `db:"original_value" json:"originalValue"`
	Value         json.RawMessage `db:"value" json:"value"`
}

func (q *Queries) getPhrases(ctx context.Context, ids []string) ([]getPhrasesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPhrases, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getPhrasesRow
	for rows.Next() {
		var i getPhrasesRow
		if err := rows.Scan(&i.Key, &i.OriginalValue, &i.Value); err != nil {
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

const getSongs = `-- name: getSongs :many
SELECT id, collection_id, title, key
FROM songs
WHERE id = ANY ($1::uuid[])
`

func (q *Queries) getSongs(ctx context.Context, ids []uuid.UUID) ([]Song, error) {
	rows, err := q.db.QueryContext(ctx, getSongs, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Song
	for rows.Next() {
		var i Song
		if err := rows.Scan(
			&i.ID,
			&i.CollectionID,
			&i.Title,
			&i.Key,
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
