package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword123"
	hash, err := HashPassword(password)

	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}

	if len(hash) == 0 {
		t.Fatal("Hash should not be empty")
	}

	// Hash should be different each time
	hash2, _ := HashPassword(password)
	if hash == hash2 {
		t.Fatal("Hashes should be different (salt should vary)")
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "mypassword123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	// Correct password should verify
	if !VerifyPassword(password, hash) {
		t.Fatal("VerifyPassword should return true for correct password")
	}

	// Wrong password should not verify
	if VerifyPassword("wrongpassword", hash) {
		t.Fatal("VerifyPassword should return false for incorrect password")
	}

	// Empty password should not verify
	if VerifyPassword("", hash) {
		t.Fatal("VerifyPassword should return false for empty password")
	}
}

func TestNeedsRehash(t *testing.T) {
	// Test with current cost
	password := "testpassword123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if NeedsRehash(hash) {
		t.Fatal("Newly created hash should not need rehashing")
	}

	// Test with invalid hash
	if !NeedsRehash("invalid_hash") {
		t.Fatal("Invalid hash should need rehashing")
	}

	// Test with empty hash
	if !NeedsRehash("") {
		t.Fatal("Empty hash should need rehashing")
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		email    string
		wantErrs int
	}{
		{
			name:     "valid password",
			password: "mypassword123",
			email:    "user@example.com",
			wantErrs: 0,
		},
		{
			name:     "too short",
			password: "abc123",
			email:    "user@example.com",
			wantErrs: 2, // too short + common password
		},
		{
			name:     "no letter",
			password: "12345678",
			email:    "user@example.com",
			wantErrs: 1, // no letter
		},
		{
			name:     "no number",
			password: "abcdefgh",
			email:    "user@example.com",
			wantErrs: 1, // no number
		},
		{
			name:     "common password",
			password: "password123",
			email:    "user@example.com",
			wantErrs: 1, // common password
		},
		{
			name:     "same as email",
			password: "user@example.com",
			email:    "user@example.com",
			wantErrs: 2, // no number + same as email
		},
		{
			name:     "case insensitive email check",
			password: "USER@EXAMPLE.COM",
			email:    "user@example.com",
			wantErrs: 2, // no number + same as email
		},
		{
			name:     "multiple errors",
			password: "abc",
			email:    "user@example.com",
			wantErrs: 2, // too short + no number (abc isn't in common passwords list)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := ValidatePassword(tt.password, tt.email)
			if len(errors) != tt.wantErrs {
				t.Errorf("ValidatePassword() got %d errors, want %d. Errors: %v",
					len(errors), tt.wantErrs, errors)
			}
		})
	}
}

func TestHasNumber(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abc123", true},
		{"abcdef", false},
		{"123456", true},
		{"", false},
		{"!@#$%", false},
		{"test5", true},
		{"5test", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := hasNumber(tt.input); got != tt.want {
				t.Errorf("hasNumber(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsCommonPassword(t *testing.T) {
	tests := []struct {
		password string
		want     bool
	}{
		{"password123", true},
		{"qwerty", true},
		{"QWERTY", true},      // case insensitive
		{"Password123", true}, // case insensitive
		{"mysecurepassword123", false},
		{"", false},
		{"notcommon456", false},
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			if got := isCommonPassword(tt.password); got != tt.want {
				t.Errorf("isCommonPassword(%q) = %v, want %v", tt.password, got, tt.want)
			}
		})
	}
}
