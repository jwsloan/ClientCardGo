// Package domain contains repository interfaces for persistence.
package domain

// UserRepository abstracts user persistence.
type UserRepository interface {
	// CreateUser persists a new user and returns its ID.
	CreateUser(u *User) (string, error)
	// GetUserByEmail fetches a user by email, or returns nil if not found.
	GetUserByEmail(email string) (*User, error)
}