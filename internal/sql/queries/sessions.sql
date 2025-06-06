-- name: CreateSession :one
INSERT INTO sessions (id, user_id, expires_at)
VALUES ($1, $2, $3)
RETURNING id, user_id, created_at, expires_at;

-- name: GetSession :one
SELECT id, user_id, created_at, expires_at
FROM sessions
WHERE id = $1 AND expires_at > CURRENT_TIMESTAMP;

-- name: GetUserBySession :one
SELECT u.id, u.email, u.created_at, u.updated_at
FROM users u
JOIN sessions s ON u.id = s.user_id
WHERE s.id = $1 AND s.expires_at > CURRENT_TIMESTAMP;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = $1;

-- name: DeleteExpiredSessions :exec
DELETE FROM sessions WHERE expires_at < CURRENT_TIMESTAMP;

-- name: DeleteUserSessions :exec
DELETE FROM sessions WHERE user_id = $1;