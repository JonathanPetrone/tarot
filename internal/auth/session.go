package auth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/jonathanpetrone/aitarot/internal/database"
)

type SessionService struct {
	db *database.Queries
}

func NewSessionService(db *database.Queries) *SessionService {
	return &SessionService{db: db}
}

// generateSessionID creates a cryptographically secure session ID
func (s *SessionService) generateSessionID() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// CreateSession creates a new session for a user
func (s *SessionService) CreateSession(ctx context.Context, userID int32, duration time.Duration) (string, error) {
	sessionID, err := s.generateSessionID()
	if err != nil {
		return "", fmt.Errorf("failed to generate session ID: %w", err)
	}

	expiresAt := time.Now().Add(duration)

	_, err = s.db.CreateSession(ctx, database.CreateSessionParams{
		ID:        sessionID,
		UserID:    sql.NullInt32{Int32: userID, Valid: true},
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}

	return sessionID, nil
}

// GetUserBySession retrieves user information from a session ID
func (s *SessionService) GetUserBySession(ctx context.Context, sessionID string) (*database.GetUserBySessionRow, error) {
	user, err := s.db.GetUserBySession(ctx, sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Session not found or expired
		}
		return nil, fmt.Errorf("failed to get user by session: %w", err)
	}
	return &user, nil
}

// DeleteSession removes a session
func (s *SessionService) DeleteSession(ctx context.Context, sessionID string) error {
	return s.db.DeleteSession(ctx, sessionID)
}

// SetSessionCookie sets the session cookie on the response
func (s *SessionService) SetSessionCookie(w http.ResponseWriter, sessionID string, duration time.Duration) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		MaxAge:   int(duration.Seconds()),
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}

// ClearSessionCookie removes the session cookie
func (s *SessionService) ClearSessionCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

// GetSessionFromRequest extracts session ID from request cookie
func (s *SessionService) GetSessionFromRequest(r *http.Request) string {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// CleanupExpiredSessions removes expired sessions (run periodically)
func (s *SessionService) CleanupExpiredSessions(ctx context.Context) error {
	return s.db.DeleteExpiredSessions(ctx)
}
