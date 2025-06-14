// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: sessions.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (id, user_id, expires_at)
VALUES ($1, $2, $3)
RETURNING id, user_id, created_at, expires_at
`

type CreateSessionParams struct {
	ID        string        `json:"id"`
	UserID    sql.NullInt32 `json:"user_id"`
	ExpiresAt time.Time     `json:"expires_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Sessions, error) {
	row := q.db.QueryRowContext(ctx, createSession, arg.ID, arg.UserID, arg.ExpiresAt)
	var i Sessions
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const deleteExpiredSessions = `-- name: DeleteExpiredSessions :exec
DELETE FROM sessions WHERE expires_at < CURRENT_TIMESTAMP
`

func (q *Queries) DeleteExpiredSessions(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteExpiredSessions)
	return err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = $1
`

func (q *Queries) DeleteSession(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteSession, id)
	return err
}

const deleteUserSessions = `-- name: DeleteUserSessions :exec
DELETE FROM sessions WHERE user_id = $1
`

func (q *Queries) DeleteUserSessions(ctx context.Context, userID sql.NullInt32) error {
	_, err := q.db.ExecContext(ctx, deleteUserSessions, userID)
	return err
}

const getSession = `-- name: GetSession :one
SELECT id, user_id, created_at, expires_at
FROM sessions
WHERE id = $1 AND expires_at > CURRENT_TIMESTAMP
`

func (q *Queries) GetSession(ctx context.Context, id string) (Sessions, error) {
	row := q.db.QueryRowContext(ctx, getSession, id)
	var i Sessions
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const getUserBySession = `-- name: GetUserBySession :one
SELECT u.id, u.email, u.first_name, u.last_name, u.zodiac, u.created_at, u.updated_at
FROM users u
JOIN sessions s ON u.id = s.user_id
WHERE s.id = $1 AND s.expires_at > CURRENT_TIMESTAMP
`

type GetUserBySessionRow struct {
	ID        int32              `json:"id"`
	Email     string             `json:"email"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Zodiac    NullZodiacSignEnum `json:"zodiac"`
	CreatedAt sql.NullTime       `json:"created_at"`
	UpdatedAt sql.NullTime       `json:"updated_at"`
}

func (q *Queries) GetUserBySession(ctx context.Context, id string) (GetUserBySessionRow, error) {
	row := q.db.QueryRowContext(ctx, getUserBySession, id)
	var i GetUserBySessionRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Zodiac,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
