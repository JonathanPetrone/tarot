package server

import (
	"net/url"
	"regexp"
	"strings"
)

var cleanupPattern = regexp.MustCompile(`[^\p{L}\p{Zs}]+`) // keep only letters and spaces

func cleanCardName(raw string) string {
	// Decode URL encoding
	decoded, _ := url.QueryUnescape(raw)

	// Remove emoji, numbers, and position labels (e.g., "🌴 1. ")
	decoded = strings.TrimSpace(decoded)

	// Remove leading emoji and digits (e.g., "🌾 4. ")
	parts := strings.SplitN(decoded, "–", 2) // keep only before the dash
	cardPart := parts[0]

	// Remove anything that's not a letter or space
	clean := cleanupPattern.ReplaceAllString(cardPart, "")
	return strings.TrimSpace(clean)
}

// Helper function to detect duplicate email errors
func isDuplicateEmail(err error) bool {
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "duplicate") ||
		strings.Contains(errStr, "unique constraint") ||
		strings.Contains(errStr, "already exists") ||
		strings.Contains(errStr, "violates unique constraint")
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

	// Updated regex for international characters
	// Local part: ASCII for compatibility, Domain: Unicode support
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[\p{L}\p{M}\p{N}.-]+\.[\p{L}]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 72 {
		return false
	}

	// Updated to use Unicode character classes
	hasUpper := regexp.MustCompile(`[\p{Lu}]`).MatchString(password)
	hasLower := regexp.MustCompile(`[\p{Ll}]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[\p{N}]`).MatchString(password)

	return hasUpper && hasLower && hasNumber
}
