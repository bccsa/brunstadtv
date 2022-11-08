// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: collections.sql

package sqlexport

import (
	"context"
)

const insertCollection = `-- name: InsertCollection :exec
INSERT INTO collections (id, name, type, collection_items)
VALUES (?,?,?,?)
`

type InsertCollectionParams struct {
	ID              int64  `db:"id" json:"id"`
	Name            string `db:"name" json:"name"`
	Type            string `db:"type" json:"type"`
	CollectionItems string `db:"collection_items" json:"collectionItems"`
}

func (q *Queries) InsertCollection(ctx context.Context, arg InsertCollectionParams) error {
	_, err := q.db.ExecContext(ctx, insertCollection,
		arg.ID,
		arg.Name,
		arg.Type,
		arg.CollectionItems,
	)
	return err
}