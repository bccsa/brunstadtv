-- name: listEpisodes :many
WITH ts AS (SELECT episodes_id,
                   json_object_agg(languages_code, title)             AS title,
                   json_object_agg(languages_code, description)       AS description,
                   json_object_agg(languages_code, extra_description) AS extra_description
            FROM episodes_translations
            GROUP BY episodes_id),
     tags AS (SELECT episodes_id,
                     array_agg(tags_id) AS tags
              FROM episodes_tags
              GROUP BY episodes_id)
SELECT e.id,
       e.legacy_id,
       e.asset_id,
       e.episode_number,
       e.image_file_id,
       e.season_id,
       e.type,
       ts.title,
       ts.description,
       ts.extra_description,
       tags.tags::int[] AS tag_ids
FROM episodes e
         LEFT JOIN ts ON e.id = ts.episodes_id
         LEFT JOIN tags ON tags.episodes_id = e.id;

-- name: getEpisodes :many
WITH ts AS (SELECT episodes_id,
                  json_object_agg(languages_code, title)             AS title,
                  json_object_agg(languages_code, description)       AS description,
                  json_object_agg(languages_code, extra_description) AS extra_description
           FROM episodes_translations
           GROUP BY episodes_id),
     tags AS (SELECT episodes_id,
                     array_agg(tags_id) AS tags
              FROM episodes_tags
              GROUP BY episodes_id)
SELECT e.id,
       e.legacy_id,
       e.asset_id,
       e.episode_number,
       e.image_file_id,
       e.season_id,
       e.type,
       ts.title,
       ts.description,
       ts.extra_description,
       tags.tags::int[] AS tag_ids
FROM episodes e
         LEFT JOIN ts ON e.id = ts.episodes_id
         LEFT JOIN tags ON tags.episodes_id = e.id
WHERE id = ANY($1::int[])
ORDER BY e.episode_number;

-- name: getEpisodeIDsForSeasons :many
SELECT
    e.id,
    e.season_id
FROM episodes e
WHERE e.season_id = ANY($1::int[])
ORDER BY e.episode_number;

-- name: getPermissionsForEpisodes :many
WITH er AS (SELECT e.id,
                   COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                             FROM episodes_usergroups eu
                             WHERE eu.episodes_id = e.id), ARRAY []::character varying[]) AS roles,
                   COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                             FROM episodes_usergroups_download eu
                             WHERE eu.episodes_id = e.id), ARRAY []::character varying[]) AS roles_download,
                   COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                             FROM episodes_usergroups_earlyaccess eu
                             WHERE eu.episodes_id = e.id),
                            ARRAY []::character varying[])                                AS roles_earlyaccess
            FROM episodes e),
     ea AS (SELECT e.id,
                   e.status::text = 'published'::text AND (e.season_id IS NULL OR (se.status::text = 'published'::text AND
                                                                                   s.status::text = 'published'::text))                  AS published,
                   COALESCE(GREATEST(e.available_from, se.available_from, s.available_from),
                            '1800-01-01 00:00:00'::timestamp without time zone) AS available_from,
                   COALESCE(LEAST(e.available_to, se.available_to, s.available_to),
                            '3000-01-01 00:00:00'::timestamp without time zone) AS available_to
            FROM episodes e
                     LEFT JOIN seasons se ON e.season_id = se.id
                     LEFT JOIN shows s ON se.show_id = s.id)
SELECT e.id,
       access.published::bool AS published,
       access.available_from::timestamp              AS available_from,
       access.available_to::timestamp                AS available_to,
       roles.roles::varchar[]                        AS usergroups,
       roles.roles_download::varchar[]               AS usergroups_downloads,
       roles.roles_earlyaccess::varchar[]            AS usergroups_earlyaccess
FROM episodes e
         LEFT JOIN ea access ON access.id = e.id
         LEFT JOIN er roles ON roles.id = e.id
WHERE e.id = ANY($1::int[]);

-- name: RefreshEpisodeAccessView :one
SELECT update_access('episodes_access');
