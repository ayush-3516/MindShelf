-- name: AddTagToLink :exec
INSERT INTO link_tags (
    link_id,
    tag_id
) VALUES (
    $1, $2
) ON CONFLICT (link_id, tag_id) DO NOTHING;

-- name: RemoveTagFromLink :exec
DELETE FROM link_tags
WHERE link_id = $1 AND tag_id = $2;

-- name: ListTagsByLink :many
SELECT t.* FROM tags t
JOIN link_tags lt ON t.id = lt.tag_id
WHERE lt.link_id = $1
ORDER BY t.name;

-- name: ListLinksByTag :many
SELECT l.* FROM links l
JOIN link_tags lt ON l.id = lt.link_id
WHERE lt.tag_id = $1 AND l.user_id = $2 AND l.archived = FALSE
ORDER BY l.created_at DESC
LIMIT $3 OFFSET $4;

