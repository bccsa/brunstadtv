-- name: getSurveys :many
WITH ts AS (SELECT ts.surveys_id                                   AS id,
                   json_object_agg(languages_code, ts.title)       AS title,
                   json_object_agg(languages_code, ts.description) AS description
            FROM surveys_translations ts
            GROUP BY ts.surveys_id)
SELECT s.id,
       s.title       AS original_title,
       s.description AS original_description,
       ts.title,
       ts.description
FROM surveys s
         LEFT JOIN ts ON ts.id = s.id
WHERE s.id = ANY (@ids::uuid[]);

-- name: getSurveyQuestions :many
WITH ts AS (SELECT ts.surveyquestions_id                           AS id,
                   json_object_agg(languages_code, ts.title)       AS title,
                   json_object_agg(languages_code, ts.description) AS description
            FROM surveyquestions_translations ts
            GROUP BY ts.surveyquestions_id)
SELECT s.id,
       s.title       AS original_title,
       s.description AS original_description,
       s.placeholder AS original_placeholder,
       s.survey_id,
       s.type,
       ts.title,
       ts.description
FROM surveyquestions s
         LEFT JOIN ts ON ts.id = s.id
WHERE s.id = ANY (@ids::uuid[]);

-- name: getQuestionIDsForSurveyIDs :many
SELECT q.id, q.survey_id AS parent_id
FROM surveyquestions q
WHERE q.survey_id = ANY (@ids::uuid[])
ORDER BY q.sort;

-- name: GetSurveyIDFromQuestionID :one
SELECT q.survey_id
FROM surveyquestions q
WHERE q.id = @id::uuid;

-- name: GetPromptIDsForRoles :many
WITH roles AS (SELECT pt.prompts_id,
                      array_agg(u.usergroups_code) AS roles
               FROM prompts_targets pt
                        LEFT JOIN targets_usergroups u ON u.targets_id = pt.targets_id
               GROUP BY pt.prompts_id)
SELECT p.id
FROM prompts p
         LEFT JOIN roles ON roles.prompts_id = p.id
WHERE p.status = 'published'
  AND p.from < (NOW() + interval '7 day')
  AND p.to > NOW()
  AND roles.roles && @roles::varchar[];

-- name: GetPrompts :many
WITH ts AS (SELECT ts.prompts_id                                       AS id,
                   json_object_agg(languages_code, ts.title)           AS title,
                   json_object_agg(languages_code, ts.secondary_title) AS secondary_title
            FROM prompts_translations ts
            GROUP BY ts.prompts_id)
SELECT p.id,
       p.title           as original_title,
       p.secondary_title as original_secondary_title,
       p.from,
       p.to,
       p.type,
       p.survey_id,
       ts.title,
       ts.secondary_title
FROM prompts p
         LEFT JOIN ts ON ts.id = p.id
WHERE p.id = ANY (@ids::uuid[]);

-- name: UpsertSurveyAnswer :exec
INSERT INTO users.surveyquestionanswers (profile_id, question_id, updated_at)
VALUES (@profile_id::uuid, @question_id::uuid, now())
ON CONFLICT(profile_id, question_id) DO UPDATE SET updated_at = EXCLUDED.updated_at;
