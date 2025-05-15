-- name: CreateTag :one
INSERT INTO tags (
    user_id,
    name
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetTagByID :one
SELECT * FROM tags
WHERE id = $1 AND user_id = $2;

-- name: GetTagByName :one
SELECT * FROM tags
WHERE name = $1 AND user_id = $2;

-- name: ListTagsByUser :many
SELECT * FROM tags
WHERE user_id = $1
ORDER BY name;

-- name: UpdateTag :one
UPDATE tags
SET name = $2
WHERE id = $1 AND user_id = $3
RETURNING *;

-- name: DeleteTag :exec
DELETE FROM tags
WHERE id = $1 AND user_id = $2;

