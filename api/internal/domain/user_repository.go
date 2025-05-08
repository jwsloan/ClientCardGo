// Package domain contains repository interfaces for persistence.
package domain

// UserRepository abstracts user persistence.
type UserRepository interface {
	CreateUser(u *User) (string, error)
	GetUserByEmail(email string) (*User, error)
	SetInterviewComplete(userID string, complete bool) error
	SetEULAAccepted(userID string, accepted bool) error
}