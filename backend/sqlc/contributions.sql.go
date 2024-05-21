// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: contributions.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const insertContribution = `-- name: InsertContribution :exec
INSERT INTO "public"."contributions" (person_id, "type", mediaitem_id, timedmetadata_id)
VALUES ($1::uuid, $2, $3::uuid, $4::uuid)
`

type InsertContributionParams struct {
	PersonID        uuid.UUID     `db:"person_id" json:"personId"`
	Type            string        `db:"type" json:"type"`
	MediaitemID     uuid.NullUUID `db:"mediaitem_id" json:"mediaitemId"`
	TimedmetadataID uuid.NullUUID `db:"timedmetadata_id" json:"timedmetadataId"`
}

func (q *Queries) InsertContribution(ctx context.Context, arg InsertContributionParams) error {
	_, err := q.db.ExecContext(ctx, insertContribution,
		arg.PersonID,
		arg.Type,
		arg.MediaitemID,
		arg.TimedmetadataID,
	)
	return err
}

const getContributionIDsForPersonsWithRoles = `-- name: getContributionIDsForPersonsWithRoles :many
WITH RelevantContributions AS (
  SELECT
    m.primary_episode_id::text as item_id,
    c.type,
    c.person_id,
    'episode' as item_type,
    m.id as mediaitem_id
  FROM
    public.mediaitems m
  INNER JOIN contributions c ON c.mediaitem_id = m.id
    and c.person_id = ANY ($2::uuid[])
    and m.primary_episode_id is not null
  UNION
  ALL
  SELECT
    tm.id::text as item_id,
    c.type,
    c.person_id,
    'chapter' as item_type,
    m.id as mediaitem_id
  FROM timedmetadata tm
  INNER JOIN mediaitems m ON
    (m.timedmetadata_from_asset AND tm.asset_id = m.asset_id)
    OR (NOT m.timedmetadata_from_asset AND tm.mediaitem_id = m.id)
  INNER JOIN contributions c ON c.timedmetadata_id = tm.id
  and c.person_id = ANY ($2::uuid[])
)
SELECT
  rc.type,
  rc.person_id,
  rc.item_type,
  rc.item_id
FROM
  RelevantContributions rc
  JOIN public.mediaitems m ON rc.mediaitem_id = m.id
  JOIN public.episode_availability access ON access.id = m.primary_episode_id
  JOIN public.episode_roles roles ON roles.id = m.primary_episode_id
WHERE
  access.published
  AND access.available_to > now()
  AND (
    (
      roles.roles && $1::varchar[]
      AND access.available_from < now()
    )
    OR (roles.roles_earlyaccess && $1::varchar[])
  )
ORDER BY
  m.published_at DESC
`

type getContributionIDsForPersonsWithRolesParams struct {
	Roles     []string    `db:"roles" json:"roles"`
	PersonIds []uuid.UUID `db:"person_ids" json:"personIds"`
}

type getContributionIDsForPersonsWithRolesRow struct {
	Type     string    `db:"type" json:"type"`
	PersonID uuid.UUID `db:"person_id" json:"personId"`
	ItemType string    `db:"item_type" json:"itemType"`
	ItemID   string    `db:"item_id" json:"itemId"`
}

func (q *Queries) getContributionIDsForPersonsWithRoles(ctx context.Context, arg getContributionIDsForPersonsWithRolesParams) ([]getContributionIDsForPersonsWithRolesRow, error) {
	rows, err := q.db.QueryContext(ctx, getContributionIDsForPersonsWithRoles, pq.Array(arg.Roles), pq.Array(arg.PersonIds))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getContributionIDsForPersonsWithRolesRow
	for rows.Next() {
		var i getContributionIDsForPersonsWithRolesRow
		if err := rows.Scan(
			&i.Type,
			&i.PersonID,
			&i.ItemType,
			&i.ItemID,
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
