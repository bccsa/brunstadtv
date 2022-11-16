// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: notifications.sql

package sqlc

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getNotifications = `-- name: getNotifications :many
WITH ts AS (SELECT ts.notificationtemplates_id,
                   json_object_agg(languages_code, title)       AS title,
                   json_object_agg(languages_code, description) AS description
            FROM notificationtemplates_translations ts
            GROUP BY ts.notificationtemplates_id),
     imgs AS (SELECT notificationtemplate_id as item_id, style, language, filename_disk
              FROM images img
                       JOIN directus_files df on img.file = df.id),
     images AS (SELECT item_id, json_agg(imgs) as json
                FROM imgs
                GROUP BY item_id)
SELECT n.id,
       n.status,
       COALESCE(ts.title, '{}')       AS title,
       COALESCE(ts.description, '{}') AS description,
       COALESCE(img.json, '[]')       AS images,
       n.action,
       n.deep_link,
       n.schedule_at,
       n.sent
FROM notifications n
         LEFT JOIN notificationtemplates t ON n.template_id = t.id
         LEFT JOIN ts ON ts.notificationtemplates_id = t.id
         LEFT JOIN images img ON img.item_id = t.id
WHERE n.id = ANY ($1::int[])
`

type getNotificationsRow struct {
	ID          uuid.UUID       `db:"id" json:"id"`
	Status      string          `db:"status" json:"status"`
	Title       json.RawMessage `db:"title" json:"title"`
	Description json.RawMessage `db:"description" json:"description"`
	Images      json.RawMessage `db:"images" json:"images"`
	Action      null_v4.String  `db:"action" json:"action"`
	DeepLink    null_v4.String  `db:"deep_link" json:"deepLink"`
	ScheduleAt  null_v4.Time    `db:"schedule_at" json:"scheduleAt"`
	Sent        sql.NullBool    `db:"sent" json:"sent"`
}

func (q *Queries) getNotifications(ctx context.Context, dollar_1 []int32) ([]getNotificationsRow, error) {
	rows, err := q.db.QueryContext(ctx, getNotifications, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getNotificationsRow
	for rows.Next() {
		var i getNotificationsRow
		if err := rows.Scan(
			&i.ID,
			&i.Status,
			&i.Title,
			&i.Description,
			&i.Images,
			&i.Action,
			&i.DeepLink,
			&i.ScheduleAt,
			&i.Sent,
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
