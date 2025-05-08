// Package usecase provides application business logic for signup.
package usecase

import (
	"errors"
	"strings"
	"time"

	"api/internal/domain"
	gofrsuuid "github.com/gofrs/uuid/v5"
)

var (
	ErrEmailTaken = errors.New("email already registered")
)

// SignupInput is the input for the signup usecase.
type SignupInput struct {
	Email    string
	Name     string
	Password string
}

// SignupOutput is the output from the signup usecase.
type SignupOutput struct {
	UserID string
	Name   string
}

// Signup handles user registration.
type Signup struct {
	Repo domain.UserRepository
}

func NewSignup(repo domain.UserRepository) *Signup {
	return &Signup{Repo: repo}
}

// generateUUIDv7 generates a UUIDv7 string using github.com/gofrs/uuid/v5.
func generateUUIDv7() (string, error) {
	u7, err := gofrsuuid.NewV7()
	if err != nil {
		return "", err
	}
	return u7.String(), nil
}

func (s *Signup) Register(input SignupInput) (*SignupOutput, error) {
	// Use UUIDv7 for user IDs (requires github.com/gofrs/uuid/v5 or similar)
	uid, err := generateUUIDv7()
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		ID:       uid,
		Email:    strings.TrimSpace(input.Email),
		Name:     strings.TrimSpace(input.Name),
		Password: input.Password, // plain, will be hashed below
		Role:     "member",
		CreatedAt: time.Now().UTC(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	existing, err := s.Repo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrEmailTaken
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	if _, err := s.Repo.CreateUser(user); err != nil {
		return nil, err
	}

	return &SignupOutput{
		UserID: user.ID,
		Name:   user.Name,
	}, nil
}