// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: achievements.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const confirmAchievement = `-- name: ConfirmAchievement :exec
UPDATE "users"."achievements"
SET confirmed_at = NOW()
WHERE profile_id = $1
  AND achievement_id = $2
`

type ConfirmAchievementParams struct {
	ProfileID     uuid.UUID `db:"profile_id" json:"profileID"`
	AchievementID uuid.UUID `db:"achievement_id" json:"achievementID"`
}

func (q *Queries) ConfirmAchievement(ctx context.Context, arg ConfirmAchievementParams) error {
	_, err := q.db.ExecContext(ctx, confirmAchievement, arg.ProfileID, arg.AchievementID)
	return err
}

const getAchievedAchievements = `-- name: GetAchievedAchievements :many
SELECT a.achievement_id as id
FROM "users"."achievements" a
WHERE a.profile_id = $1
  AND a.achievement_id = ANY ($2::uuid[])
`

type GetAchievedAchievementsParams struct {
	ProfileID uuid.UUID   `db:"profile_id" json:"profileID"`
	Column2   []uuid.UUID `db:"column_2" json:"column2"`
}

func (q *Queries) GetAchievedAchievements(ctx context.Context, arg GetAchievedAchievementsParams) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getAchievedAchievements, arg.ProfileID, pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
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

const getAchievementsWithConditionAchieved = `-- name: GetAchievementsWithConditionAchieved :many
SELECT c.achievement_id AS id, array_agg(c.id)::uuid[] AS condition_ids
FROM "public"."achievementconditions" c
         LEFT JOIN "users"."achievements" achieved
                   ON achieved.profile_id = $1 AND achieved.achievement_id = c.achievement_id
WHERE achieved IS NULL
  AND c.collection = $2
  AND c.action = $3
  AND c.amount <= $4
GROUP BY c.achievement_id
`

type GetAchievementsWithConditionAchievedParams struct {
	ProfileID  uuid.UUID `db:"profile_id" json:"profileID"`
	Collection string    `db:"collection" json:"collection"`
	Action     string    `db:"action" json:"action"`
	Amount     int32     `db:"amount" json:"amount"`
}

type GetAchievementsWithConditionAchievedRow struct {
	ID           uuid.UUID   `db:"id" json:"id"`
	ConditionIds []uuid.UUID `db:"condition_ids" json:"conditionIds"`
}

func (q *Queries) GetAchievementsWithConditionAchieved(ctx context.Context, arg GetAchievementsWithConditionAchievedParams) ([]GetAchievementsWithConditionAchievedRow, error) {
	rows, err := q.db.QueryContext(ctx, getAchievementsWithConditionAchieved,
		arg.ProfileID,
		arg.Collection,
		arg.Action,
		arg.Amount,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAchievementsWithConditionAchievedRow
	for rows.Next() {
		var i GetAchievementsWithConditionAchievedRow
		if err := rows.Scan(&i.ID, pq.Array(&i.ConditionIds)); err != nil {
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

const listAchievementGroups = `-- name: ListAchievementGroups :many
SELECT id
FROM "public"."achievementgroups"
WHERE status = 'published'
`

func (q *Queries) ListAchievementGroups(ctx context.Context) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, listAchievementGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
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

const listAchievements = `-- name: ListAchievements :many
SELECT id
FROM "public"."achievements"
WHERE status = 'published'
ORDER BY sort
`

func (q *Queries) ListAchievements(ctx context.Context) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, listAchievements)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
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

const setAchievementAchieved = `-- name: SetAchievementAchieved :exec
INSERT INTO "users"."achievements" (profile_id, achievement_id, achieved_at, condition_ids)
VALUES ($1, $2, now(), $3)
ON CONFLICT(profile_id, achievement_id) DO UPDATE SET achieved_at = now()
`

type SetAchievementAchievedParams struct {
	ProfileID     uuid.UUID   `db:"profile_id" json:"profileID"`
	AchievementID uuid.UUID   `db:"achievement_id" json:"achievementID"`
	ConditionIds  []uuid.UUID `db:"condition_ids" json:"conditionIds"`
}

func (q *Queries) SetAchievementAchieved(ctx context.Context, arg SetAchievementAchievedParams) error {
	_, err := q.db.ExecContext(ctx, setAchievementAchieved, arg.ProfileID, arg.AchievementID, pq.Array(arg.ConditionIds))
	return err
}

const achievementAchievedAt = `-- name: achievementAchievedAt :many
SELECT a.achievement_id,
       a.achieved_at,
       a.confirmed_at
FROM "users"."achievements" a
WHERE a.profile_id = $1
  AND a.achievement_id = ANY ($2::uuid[])
`

type achievementAchievedAtParams struct {
	ProfileID uuid.UUID   `db:"profile_id" json:"profileID"`
	Column2   []uuid.UUID `db:"column_2" json:"column2"`
}

type achievementAchievedAtRow struct {
	AchievementID uuid.UUID    `db:"achievement_id" json:"achievementID"`
	AchievedAt    time.Time    `db:"achieved_at" json:"achievedAt"`
	ConfirmedAt   null_v4.Time `db:"confirmed_at" json:"confirmedAt"`
}

func (q *Queries) achievementAchievedAt(ctx context.Context, arg achievementAchievedAtParams) ([]achievementAchievedAtRow, error) {
	rows, err := q.db.QueryContext(ctx, achievementAchievedAt, arg.ProfileID, pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []achievementAchievedAtRow
	for rows.Next() {
		var i achievementAchievedAtRow
		if err := rows.Scan(&i.AchievementID, &i.AchievedAt, &i.ConfirmedAt); err != nil {
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

const getAchievedAchievementsForProfiles = `-- name: getAchievedAchievementsForProfiles :many
SELECT a.achievement_id as id, a.profile_id as parent_id
FROM "users"."achievements" a
WHERE a.profile_id = ANY ($1::uuid[])
ORDER BY a.achieved_at DESC
`

type getAchievedAchievementsForProfilesRow struct {
	ID       uuid.UUID `db:"id" json:"id"`
	ParentID uuid.UUID `db:"parent_id" json:"parentID"`
}

func (q *Queries) getAchievedAchievementsForProfiles(ctx context.Context, dollar_1 []uuid.UUID) ([]getAchievedAchievementsForProfilesRow, error) {
	rows, err := q.db.QueryContext(ctx, getAchievedAchievementsForProfiles, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getAchievedAchievementsForProfilesRow
	for rows.Next() {
		var i getAchievedAchievementsForProfilesRow
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

const getAchievementGroups = `-- name: getAchievementGroups :many
WITH ts AS (SELECT achievementgroups_id, json_object_agg(languages_code, title) as title
            FROM achievementgroups_translations
            GROUP BY achievementgroups_id)
SELECT ag.id, ag.title as original_title, ts.title
FROM achievementgroups ag
         LEFT JOIN ts ON ts.achievementgroups_id = ag.id
WHERE ag.id = ANY ($1::uuid[])
`

type getAchievementGroupsRow struct {
	ID            uuid.UUID             `db:"id" json:"id"`
	OriginalTitle null_v4.String        `db:"original_title" json:"originalTitle"`
	Title         pqtype.NullRawMessage `db:"title" json:"title"`
}

func (q *Queries) getAchievementGroups(ctx context.Context, dollar_1 []uuid.UUID) ([]getAchievementGroupsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAchievementGroups, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getAchievementGroupsRow
	for rows.Next() {
		var i getAchievementGroupsRow
		if err := rows.Scan(&i.ID, &i.OriginalTitle, &i.Title); err != nil {
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

const getAchievements = `-- name: getAchievements :many
WITH ts AS (SELECT achievements_id,
                   json_object_agg(languages_code, title)       as title,
                   json_object_agg(languages_code, description) as description
            FROM achievements_translations
            GROUP BY achievements_id),
     cons AS (SELECT achievement_id,
                     json_agg(c) as conditions
              FROM achievementconditions c
              GROUP BY achievement_id),
     images AS (SELECT achievement_id, json_object_agg(COALESCE(language, 'no'), df.filename_disk) as images
                FROM achievements_images
                         JOIN directus_files df on achievements_images.image = df.id
                GROUP BY achievement_id)
SELECT a.id,
       a.group_id,
       a.title as original_title,
       a.description as original_description,
       ts.title,
       ts.description,
       images.images,
       cons.conditions
FROM "public"."achievements" a
         LEFT JOIN ts ON ts.achievements_id = a.id
         LEFT JOIN cons ON cons.achievement_id = a.id
         LEFT JOIN images ON images.achievement_id = a.id
WHERE a.id = ANY ($1::uuid[])
ORDER BY sort
`

type getAchievementsRow struct {
	ID                  uuid.UUID             `db:"id" json:"id"`
	GroupID             uuid.NullUUID         `db:"group_id" json:"groupID"`
	OriginalTitle       string                `db:"original_title" json:"originalTitle"`
	OriginalDescription null_v4.String        `db:"original_description" json:"originalDescription"`
	Title               pqtype.NullRawMessage `db:"title" json:"title"`
	Description         pqtype.NullRawMessage `db:"description" json:"description"`
	Images              pqtype.NullRawMessage `db:"images" json:"images"`
	Conditions          pqtype.NullRawMessage `db:"conditions" json:"conditions"`
}

func (q *Queries) getAchievements(ctx context.Context, dollar_1 []uuid.UUID) ([]getAchievementsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAchievements, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getAchievementsRow
	for rows.Next() {
		var i getAchievementsRow
		if err := rows.Scan(
			&i.ID,
			&i.GroupID,
			&i.OriginalTitle,
			&i.OriginalDescription,
			&i.Title,
			&i.Description,
			&i.Images,
			&i.Conditions,
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

const getAchievementsForActions = `-- name: getAchievementsForActions :many
SELECT achievement_id::uuid as id, action::varchar as parent_id
FROM achievementconditions
WHERE collection = $1::varchar
  AND action = ANY ($2::varchar[])
`

type getAchievementsForActionsParams struct {
	Column1 string   `db:"column_1" json:"column1"`
	Column2 []string `db:"column_2" json:"column2"`
}

type getAchievementsForActionsRow struct {
	ID       uuid.UUID `db:"id" json:"id"`
	ParentID string    `db:"parent_id" json:"parentID"`
}

func (q *Queries) getAchievementsForActions(ctx context.Context, arg getAchievementsForActionsParams) ([]getAchievementsForActionsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAchievementsForActions, arg.Column1, pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getAchievementsForActionsRow
	for rows.Next() {
		var i getAchievementsForActionsRow
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

const getAchievementsForGroups = `-- name: getAchievementsForGroups :many
SELECT id, group_id::uuid as parent_id
FROM "public"."achievements"
WHERE group_id = ANY ($1::uuid[])
ORDER BY sort
`

type getAchievementsForGroupsRow struct {
	ID       uuid.UUID `db:"id" json:"id"`
	ParentID uuid.UUID `db:"parent_id" json:"parentID"`
}

func (q *Queries) getAchievementsForGroups(ctx context.Context, dollar_1 []uuid.UUID) ([]getAchievementsForGroupsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAchievementsForGroups, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getAchievementsForGroupsRow
	for rows.Next() {
		var i getAchievementsForGroupsRow
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

const getUnconfirmedAchievedAchievementsForProfiles = `-- name: getUnconfirmedAchievedAchievementsForProfiles :many
SELECT a.achievement_id as id, a.profile_id as parent_id
FROM "users"."achievements" a
WHERE a.profile_id = ANY ($1::uuid[])
  AND a.confirmed_at IS NULL
`

type getUnconfirmedAchievedAchievementsForProfilesRow struct {
	ID       uuid.UUID `db:"id" json:"id"`
	ParentID uuid.UUID `db:"parent_id" json:"parentID"`
}

func (q *Queries) getUnconfirmedAchievedAchievementsForProfiles(ctx context.Context, dollar_1 []uuid.UUID) ([]getUnconfirmedAchievedAchievementsForProfilesRow, error) {
	rows, err := q.db.QueryContext(ctx, getUnconfirmedAchievedAchievementsForProfiles, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getUnconfirmedAchievedAchievementsForProfilesRow
	for rows.Next() {
		var i getUnconfirmedAchievedAchievementsForProfilesRow
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
