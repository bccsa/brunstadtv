// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: directus-files.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getFiles = `-- name: getFiles :many
SELECT id, storage, filename_disk, filename_download, title, type, folder, uploaded_by, uploaded_on, modified_by, modified_on, charset, filesize, width, height, duration, embed, description, location, tags, metadata FROM directus_files WHERE id = ANY($1::uuid[])
`

func (q *Queries) getFiles(ctx context.Context, dollar_1 []uuid.UUID) ([]DirectusFile, error) {
	rows, err := q.db.QueryContext(ctx, getFiles, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DirectusFile
	for rows.Next() {
		var i DirectusFile
		if err := rows.Scan(
			&i.ID,
			&i.Storage,
			&i.FilenameDisk,
			&i.FilenameDownload,
			&i.Title,
			&i.Type,
			&i.Folder,
			&i.UploadedBy,
			&i.UploadedOn,
			&i.ModifiedBy,
			&i.ModifiedOn,
			&i.Charset,
			&i.Filesize,
			&i.Width,
			&i.Height,
			&i.Duration,
			&i.Embed,
			&i.Description,
			&i.Location,
			&i.Tags,
			&i.Metadata,
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
