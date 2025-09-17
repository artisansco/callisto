-- name: GetUser :one
SELECT id, name, email FROM users
WHERE id = ? LIMIT 1;

-- name: GetAllUsers :many
SELECT id, name, email FROM users
ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users
set name = ?, email = ?
WHERE id = ?
RETURNING *;
