package auth

import (
	"crypto/rand"
	"encoding/base64"
)

const SessionIDLength = 32 // bytes (256 bits)

func GenerateSessionID() string {
	// Create a byte slice for random data
	bytes := make([]byte, SessionIDLength)

	// Fill with cryptographically secure random bytes
	_, err := rand.Read(bytes)
	if err != nil {
		// This should never happen with a properly functioning system
		panic("crypto/rand is unavailable: " + err.Error())
	}

	// Encode to URL-safe base64 (no padding)
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes)
}
