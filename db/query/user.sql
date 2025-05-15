-- name: CreateUser :one
INSERT INTO users (
    email,
    password_hash,
    name
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET
    email = COALESCE(sqlc.narg('email'), email),
    name = COALESCE(sqlc.narg('name'), name),
    password_hash = COALESCE(sqlc.narg('password_hash'), password_hash),
    updated_at = NOW()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

