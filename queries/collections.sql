-- name: getCollections :many
WITH ts AS (SELECT collections_id,
                   json_object_agg(languages_code, title) AS title,
                   json_object_agg(languages_code, slug)  AS slugs
            FROM collections_translations
            GROUP BY collections_id)
SELECT c.id,
       c.advanced_type,
       c.date_created,
       c.date_updated,
       c.filter_type,
       c.query_filter,
       c.number_in_titles,
       ts.title,
       ts.slugs
FROM collections c
         LEFT JOIN ts ON ts.collections_id = c.id
WHERE c.id = ANY ($1::int[]);

-- name: getCollectionEntriesForCollections :many
SELECT *
FROM collections_entries ci
WHERE ci.collections_id = ANY ($1::int[]);

-- name: getCollectionEntriesForCollectionsWithRoles :many
WITH game_roles AS (SELECT r.games_id,
                           array_agg(r.usergroups_code) AS roles
                    FROM games_usergroups r
                    GROUP BY r.games_id)
SELECT ce.*
FROM collections_entries ce
         LEFT JOIN episode_roles er ON ce.collection = 'episodes' AND er.id::varchar = ce.item
         LEFT JOIN episode_availability ea ON ce.collection = 'episodes' AND ea.id::varchar = ce.item
         LEFT JOIN season_roles sr ON ce.collection = 'seasons' AND sr.id::varchar = ce.item
         LEFT JOIN season_availability sa ON ce.collection = 'seasons' AND sa.id::varchar = ce.item
         LEFT JOIN show_roles shr ON ce.collection = 'shows' AND shr.id::varchar = ce.item
         LEFT JOIN show_availability sha ON ce.collection = 'shows' AND sha.id::varchar = ce.item
         LEFT JOIN game_roles gr ON ce.collection = 'games' AND gr.games_id::varchar = ce.item
WHERE ce.collections_id = ANY (@collectionIDs::int[])
  AND (ce.collection != 'episodes' OR (
        ea.published
        AND ea.available_to > now()
        AND er.roles && @roles::varchar[] AND ea.available_from < now()
    ))
  AND (ce.collection != 'seasons' OR (
        sa.published
        AND sa.available_to > now()
        AND sr.roles && @roles::varchar[] AND sa.available_from < now()
    ))
  AND (ce.collection != 'shows' OR (
        sha.published
        AND sha.available_to > now()
        AND shr.roles && @roles::varchar[] AND sha.available_from < now()
    ))
  AND (ce.collection != 'games' OR (gr.roles && @roles::varchar[]))
ORDER BY ce.sort;

-- name: getCollectionIDsForCodes :many
SELECT c.id, ct.slug
FROM collections c
         JOIN collections_translations ct ON c.id = ct.collections_id AND ct.slug = ANY ($1::varchar[]);
