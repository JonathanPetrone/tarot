package auth

import (
	"context"
	"net/http"

	"github.com/jonathanpetrone/aitarot/internal/database"
)

type User struct {
	ID    int32
	Email string
}

type contextKey string

const userContextKey contextKey = "user"

// Core Middleware Function
func RequireAuth(db *database.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract session ID from cookie
			cookie, err := r.Cookie("session_id")
			if err != nil {
				// No session cookie, redirect to login
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			sessionID := cookie.Value
			if sessionID == "" {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			// Get user by session from database
			userRow, err := db.GetUserBySession(r.Context(), sessionID)
			if err != nil {
				// Invalid session, redirect to login
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			// Create user struct
			user := &User{
				ID:    userRow.ID,
				Email: userRow.Email,
			}

			// Inject user into request context
			ctx := context.WithValue(r.Context(), userContextKey, user)
			r = r.WithContext(ctx)

			// Continue to next handler
			next.ServeHTTP(w, r)
		})
	}
}

// Supporting Functions
func GetUserFromContext(r *http.Request) (*User, bool) {
	user, ok := r.Context().Value(userContextKey).(*User)
	if !ok || user == nil {
		return nil, false
	}
	return user, true
}
