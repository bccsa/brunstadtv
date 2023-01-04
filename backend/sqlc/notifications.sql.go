// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: notifications.sql

package sqlc

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

const notificationMarkSendCompleted = `-- name: NotificationMarkSendCompleted :exec
UPDATE notifications n
SET send_completed = NOW()
WHERE id = $1
`

func (q *Queries) NotificationMarkSendCompleted(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, notificationMarkSendCompleted, id)
	return err
}

const notificationMarkSendStarted = `-- name: NotificationMarkSendStarted :exec
UPDATE notifications n
SET send_started = NOW()
WHERE id = $1
`

func (q *Queries) NotificationMarkSendStarted(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, notificationMarkSendStarted, id)
	return err
}

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
                GROUP BY item_id),
     target_ids AS (SELECT notifications_id, array_agg(targets_id)::uuid[] AS targets
                    FROM notifications_targets
                    GROUP BY notifications_id)
SELECT n.id,
       n.status,
       COALESCE(ts.title, '{}')       AS title,
       COALESCE(ts.description, '{}') AS description,
       COALESCE(img.json, '[]')       AS images,
       n.action,
       n.deep_link,
       n.schedule_at,
       n.send_started,
       n.send_completed,
       ti.targets                     AS target_ids
FROM notifications n
         LEFT JOIN notificationtemplates t ON n.template_id = t.id
         LEFT JOIN ts ON ts.notificationtemplates_id = t.id
         LEFT JOIN images img ON img.item_id = t.id
         LEFT JOIN target_ids ti ON ti.notifications_id = n.id
WHERE n.id = ANY ($1::uuid[])
`

type getNotificationsRow struct {
	ID            uuid.UUID       `db:"id" json:"id"`
	Status        string          `db:"status" json:"status"`
	Title         json.RawMessage `db:"title" json:"title"`
	Description   json.RawMessage `db:"description" json:"description"`
	Images        json.RawMessage `db:"images" json:"images"`
	Action        null_v4.String  `db:"action" json:"action"`
	DeepLink      null_v4.String  `db:"deep_link" json:"deepLink"`
	ScheduleAt    null_v4.Time    `db:"schedule_at" json:"scheduleAt"`
	SendStarted   null_v4.Time    `db:"send_started" json:"sendStarted"`
	SendCompleted null_v4.Time    `db:"send_completed" json:"sendCompleted"`
	TargetIds     []uuid.UUID     `db:"target_ids" json:"targetIds"`
}

func (q *Queries) getNotifications(ctx context.Context, dollar_1 []uuid.UUID) ([]getNotificationsRow, error) {
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
			&i.SendStarted,
			&i.SendCompleted,
			pq.Array(&i.TargetIds),
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
