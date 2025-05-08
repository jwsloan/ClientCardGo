// Package auth provides JWT token generation for authentication.
package auth

import (
	"time"
	"os"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type JWTGenerator struct {
	Secret []byte
	Expiry time.Duration
}

func NewJWTGeneratorFromEnv() (*JWTGenerator, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET not set")
	}
	expiry := 24 * time.Hour
	return &JWTGenerator{
		Secret: []byte(secret),
		Expiry: expiry,
	}, nil
}

func (j *JWTGenerator) GenerateWithRole(userID, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(j.Expiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.Secret)
}