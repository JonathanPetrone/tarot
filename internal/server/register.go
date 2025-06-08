package server

import (
	"errors"
	"regexp"
	"strings"
)

type RegisterRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Password    string `json:"password"`
}

func (r *RegisterRequest) Validate() error {
	// Sanitize inputs
	r.FirstName = strings.TrimSpace(r.FirstName)
	r.LastName = strings.TrimSpace(r.LastName)
	r.Email = strings.ToLower(strings.TrimSpace(r.Email))

	// Validate required fields
	if r.FirstName == "" {
		return errors.New("first name is required")
	}
	if r.LastName == "" {
		return errors.New("last name is required")
	}
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}

	// Validate business rules
	if !isValidEmail(r.Email) {
		return errors.New("invalid email format")
	}

	if !isValidPassword(r.Password) {
		return errors.New("password must be at least 8 characters with uppercase, lowercase, and number")
	}

	return nil
}

func isValidEmail(email string) bool {
	// Basic length check
	if len(email) < 3 || len(email) > 254 {
		return false
	}

	// Must contain exactly one @
	if strings.Count(email, "@") != 1 {
		return false
	}

	// Split into local and domain parts
	parts := strings.Split(email, "@")
	local, domain := parts[0], parts[1]

	// Local part checks
	if len(local) == 0 || len(local) > 64 {
		return false
	}

	// Domain part checks
	if len(domain) == 0 || len(domain) > 253 {
		return false
	}

	// Must contain at least one dot in domain
	if !strings.Contains(domain, ".") {
		return false
	}

	// Regex for valid characters
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 72 {
		return false
	}

	// Must contain at least one uppercase, one lowercase, and one digit
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	return hasUpper && hasLower && hasNumber
}
