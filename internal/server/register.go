package server

import (
	"errors"
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
