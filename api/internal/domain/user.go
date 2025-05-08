// Package domain defines the core business entities for ClientCard.
package domain

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents a client or contractor account.
type User struct {
	ID        string
	Email     string
	Name      string
	Password  string // hashed
	CreatedAt time.Time
}

// User validation errors
var (
	ErrInvalidEmail    = errors.New("invalid email address")
	ErrWeakPassword    = errors.New("password does not meet strength requirements")
	ErrNameRequired    = errors.New("name is required")
)

var (
	emailRegex = regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,})
)

// Validate checks if the user fields are valid for signup.
func (u *User) Validate() error {
	if !emailRegex.MatchString(u.Email) {
		return ErrInvalidEmail
	}
	if len(u.Name) == 0 {
		return ErrNameRequired
	}
	if err := validatePassword(u.Password); err != nil {
		return err
	}
	return nil
}

// validatePassword checks password strength requirements.
func validatePassword(pw string) error {
	pw = strings.TrimSpace(pw)
	if len(pw) < 8 {
		return ErrWeakPassword
	}
	// At least one uppercase, one lowercase, one digit
	var (
		hasUpper = false
		hasLower = false
		hasDigit = false
	)
	for _, c := range pw {
		switch {
		case 'A' <= c && c <= 'Z':
			hasUpper = true
		case 'a' <= c && c <= 'z':
			hasLower = true
		case '0' <= c && c <= '9':
			hasDigit = true
		}
	}
	if !hasUpper || !hasLower || !hasDigit {
		return ErrWeakPassword
	}
	return nil
}

// HashPassword hashes the user's password.
func (u *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

// CheckPassword verifies a plaintext password against the hash.
func (u *User) CheckPassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
}