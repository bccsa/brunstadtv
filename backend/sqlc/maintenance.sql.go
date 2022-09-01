// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: maintenance.sql

package sqlc

import (
	"context"
	"database/sql"
	"encoding/json"
)

const getMaintenanceMessage = `-- name: getMaintenanceMessage :one
WITH ts AS (SELECT messagetemplates_id,
                   json_object_agg(languages_code, message) AS message,
                   json_object_agg(languages_code, details) AS details
            FROM messagetemplates_translations
            GROUP BY messagetemplates_id),
     messages AS (SELECT mt.id,
                         mm.maintenancemessage_id,
                         ts.message,
                         ts.details
                  FROM messagetemplates mt
                           LEFT JOIN ts ON ts.messagetemplates_id = mt.id
                           JOIN maintenancemessage_messagetemplates mm on mt.id = mm.messagetemplates_id)
SELECT m.id, m.active, json_agg(ms) as messages
FROM maintenancemessage m
         JOIN messages ms ON ms.maintenancemessage_id = m.id
GROUP BY m.id, m.active
LIMIT 1
`

type getMaintenanceMessageRow struct {
	ID       int32           `db:"id" json:"id"`
	Active   sql.NullBool    `db:"active" json:"active"`
	Messages json.RawMessage `db:"messages" json:"messages"`
}

func (q *Queries) getMaintenanceMessage(ctx context.Context) (getMaintenanceMessageRow, error) {
	row := q.db.QueryRowContext(ctx, getMaintenanceMessage)
	var i getMaintenanceMessageRow
	err := row.Scan(&i.ID, &i.Active, &i.Messages)
	return i, err
}
