-- name: CreateLink :one
INSERT INTO links (
    user_id,
    title,
    url,
    description,
    domain,
    favicon_url,
    is_favorite,
    archived
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetLinkByID :one
SELECT * FROM links
WHERE id = $1 AND user_id = $2;

-- name: ListLinksByUser :many
SELECT * FROM links
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateLink :one
UPDATE links
SET
    title = COALESCE(sqlc.narg('title'), title),
    url = COALESCE(sqlc.narg('url'), url),
    description = COALESCE(sqlc.narg('description'), description),
    domain = COALESCE(sqlc.narg('domain'), domain),
    favicon_url = COALESCE(sqlc.narg('favicon_url'), favicon_url),
    is_favorite = COALESCE(sqlc.narg('is_favorite'), is_favorite),
    archived = COALESCE(sqlc.narg('archived'), archived),
    updated_at = NOW()
WHERE id = sqlc.arg('id') AND user_id = sqlc.arg('user_id')
RETURNING *;

-- name: DeleteLink :exec
DELETE FROM links
WHERE id = $1 AND user_id = $2;

-- name: ListFavoriteLinksByUser :many
SELECT * FROM links
WHERE user_id = $1 AND is_favorite = TRUE AND archived = FALSE
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListArchivedLinksByUser :many
SELECT * FROM links
WHERE user_id = $1 AND archived = TRUE
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

