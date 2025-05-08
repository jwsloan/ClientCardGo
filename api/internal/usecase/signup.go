// Package usecase provides application business logic for signup.
package usecase

import (
	"errors"
	"strings"
	"time"

	"api/internal/domain"
	"github.com/google/uuid"
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

func (s *Signup) Register(input SignupInput) (*SignupOutput, error) {
	user := &domain.User{
		ID:       uuid.NewString(),
		Email:    strings.TrimSpace(input.Email),
		Name:     strings.TrimSpace(input.Name),
		Password: input.Password, // plain, will be hashed below
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