package auth

import (
	"context"
	"net/http"

	"github.com/jonathanpetrone/aitarot/internal/database"
)

type contextKey string

const UserContextKey contextKey = "user"

type AuthMiddleware struct {
	sessionService *SessionService
}

func NewAuthMiddleware(sessionService *SessionService) *AuthMiddleware {
	return &AuthMiddleware{
		sessionService: sessionService,
	}
}

// RequireAuth is middleware that requires authentication
func (a *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID := a.sessionService.GetSessionFromRequest(r)
		if sessionID == "" {
			http.Redirect(w, r, "/login-user", http.StatusSeeOther)
			return
		}

		user, err := a.sessionService.GetUserBySession(r.Context(), sessionID)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if user == nil {
			// Session expired or invalid
			a.sessionService.ClearSessionCookie(w)
			http.Redirect(w, r, "/login-user", http.StatusSeeOther)
			return
		}

		// Add user to request context
		ctx := context.WithValue(r.Context(), UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// OptionalAuth middleware that loads user if session exists but doesn't require it
func (a *AuthMiddleware) OptionalAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID := a.sessionService.GetSessionFromRequest(r)
		if sessionID != "" {
			user, err := a.sessionService.GetUserBySession(r.Context(), sessionID)
			if err == nil && user != nil {
				ctx := context.WithValue(r.Context(), UserContextKey, user)
				r = r.WithContext(ctx)
			}
		}
		next.ServeHTTP(w, r)
	}
}

// GetUserFromContext extracts user from request context
func GetUserFromContext(ctx context.Context) *database.GetUserBySessionRow {
	user, ok := ctx.Value(UserContextKey).(*database.GetUserBySessionRow)
	if !ok {
		return nil
	}
	return user
}
