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
	Repo        domain.UserRepository
	Invitations domain.InvitationRepository
}

func NewSignup(repo domain.UserRepository, invitations domain.InvitationRepository) *Signup {
	return &Signup{Repo: repo, Invitations: invitations}
}

// generateUUIDv7 generates a UUIDv7 string using github.com/gofrs/uuid/v5.
func generateUUIDv7() (string, error) {
	u7, err := gofrsuuid.NewV7()
	if err != nil {
		return "", err
	}
	return u7.String(), nil
}

import (
	"context"
	"fmt"
)

func (s *Signup) Register(ctx context.Context, input SignupInput) (*SignupOutput, error) {
	if input.InvitationToken == "" {
		return nil, domain.ErrInvalidToken
	}
	// Validate invitation token (case-sensitive, unused)
	inv, err := s.Invitations.FindByToken(input.InvitationToken)
	if err != nil || inv == nil {
		return nil, domain.ErrInvalidToken
	}
	if inv.Status == domain.InvitationUsed {
		return nil, domain.ErrTokenUsed
	}
	if inv.Status == domain.InvitationExpired {
		return nil, domain.ErrTokenExpired
	}

	uid, err := generateUUIDv7()
	if err != nil {
		return nil, fmt.Errorf("could not generate user ID: %w", err)
	}
	user := &domain.User{
		ID:       uid,
		Email:    strings.TrimSpace(input.Email),
		Name:     strings.TrimSpace(input.Name),
		Password: input.Password,
		Role:     "member",
		CreatedAt: time.Now().UTC(),
	}

	if err := user.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	existing, err := s.Repo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, fmt.Errorf("could not check for existing user: %w", err)
	}
	if existing != nil {
		return nil, ErrEmailTaken
	}

	if err := user.HashPassword(); err != nil {
		return nil, fmt.Errorf("could not hash password: %w", err)
	}

	if _, err := s.Repo.CreateUser(user); err != nil {
		return nil, fmt.Errorf("could not create user: %w", err)
	}

	// Mark invitation used, associate with user
	if err := s.Invitations.MarkUsed(inv.Token, user.ID); err != nil {
		return nil, fmt.Errorf("could not mark invitation as used: %w", err)
	}

	return &SignupOutput{
		UserID: user.ID,
		Name:   user.Name,
	}, nil
}