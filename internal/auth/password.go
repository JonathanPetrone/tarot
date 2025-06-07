package auth

import (
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const DefaultCost = 12

func HashPassword(plaintext string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(plaintext, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintext))
	return err == nil
}

// NeedsRehash checks if a password hash should be upgraded
func NeedsRehash(hash string) bool {
	cost, err := bcrypt.Cost([]byte(hash))
	if err != nil {
		return true // If we can't read the cost, assume it needs rehashing
	}
	return cost < DefaultCost
}

func ValidatePassword(password, email string) []string {
	var errors []string

	if len(password) < 8 {
		errors = append(errors, "Password must be at least 8 characters")
	}

	if !hasLetter(password) {
		errors = append(errors, "Password must contain at least one letter")
	}

	if !hasNumber(password) {
		errors = append(errors, "Password must contain at least one number")
	}

	if isCommonPassword(password) {
		errors = append(errors, "Password is too common, please choose a different one")
	}

	if strings.EqualFold(password, email) {
		errors = append(errors, "Password cannot be the same as your email")
	}

	return errors
}

func hasLetter(s string) bool {
	for _, char := range s {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}

func hasNumber(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

var commonPasswordsSet = map[string]struct{}{
	"password": {}, "123456789": {}, "qwerty": {}, "abc123": {},
	"password123": {}, "admin": {}, "letmein": {}, "welcome": {},
	"monkey": {}, "dragon": {}, "princess": {}, "hello": {},
	"password1": {}, "qwerty123": {}, "1234qwer": {}, "abc12345": {},
	"123qweasd": {}, "qwertyui1": {}, "admin1234": {}, "welcome1": {},
	"1q2w3e4r": {}, "qwe12345": {}, "1234567a": {}, "passw0rd": {},
	"1qaz2wsx": {}, "test1234": {}, "sunshine1": {}, "monkey123": {},
	"football1": {}, "princess1": {}, "iloveyou1": {}, "dragon123": {},
	"michael1": {}, "letmein1": {}, "superman1": {}, "shadow123": {},
	"baseball1": {}, "master123": {}, "654321qwe": {}, "123abc456": {},
	"qazwsxedc1": {}, "12345qwert": {}, "123456qwe": {}, "123qazwsx": {},
	"1234asdf": {}, "asdf1234": {}, "1q2w3e4r5t": {}, "12345asdf": {},
	"1234567q": {}, "1qazxsw2": {}, "1q2w3e4r5t6y": {}, "1234zxcv": {},
	"123qwe123": {}, "12345678a": {}, "123456789a": {}, "qazwsx123": {},
	"12345zxcv": {}, "1qaz2wsx3edc": {}, "1234qazx": {}, "1qazwsx2": {},
	"1234qwert": {}, "qwerty1234": {},
}

func isCommonPassword(password string) bool {
	lower := strings.ToLower(password)
	_, exists := commonPasswordsSet[lower]
	return exists
}
