-- name: CreateUser :one
INSERT INTO users (email, password_hash, first_name, last_name, date_of_birth, zodiac)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, email, first_name, last_name, date_of_birth, zodiac, created_at, updated_at;

-- name: GetUserByEmail :one
SELECT id, email, password_hash, first_name, last_name, date_of_birth, zodiac, created_at, updated_at
FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT id, email, first_name, last_name, date_of_birth, zodiac, created_at, updated_at
FROM users
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;