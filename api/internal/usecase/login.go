// Package usecase provides business logic for login.
package usecase

import (
	"errors"
	"strings"

	"api/internal/domain"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
)

// LoginInput is the input for the login usecase.
type LoginInput struct {
	Email    string
	Password string
}

// LoginOutput is the output from the login usecase.
type LoginOutput struct {
	UserID string
	Name   string
}

type AuthTokenGenerator interface {
	Generate(userID string) (string, error)
}

// Login handles user authentication.
type Login struct {
	Repo    domain.UserRepository
	Tokens  AuthTokenGenerator
}

func NewLogin(repo domain.UserRepository, tokens AuthTokenGenerator) *Login {
	return &Login{
		Repo:   repo,
		Tokens: tokens,
	}
}

func (l *Login) Authenticate(input LoginInput) (*LoginOutput, string, error) {
	email := strings.TrimSpace(input.Email)
	pw := input.Password

	user, err := l.Repo.GetUserByEmail(email)
	if err != nil || user == nil {
		return nil, "", ErrInvalidCredentials
	}
	if err := user.CheckPassword(pw); err != nil {
		return nil, "", ErrInvalidCredentials
	}
	token, err := l.Tokens.Generate(user.ID)
	if err != nil {
		return nil, "", err
	}
	return &LoginOutput{
		UserID: user.ID,
		Name:   user.Name,
	}, token, nil
}