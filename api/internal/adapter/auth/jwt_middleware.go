// Package auth provides reusable JWT authentication and role enforcement middleware for HTTP handlers.
package auth

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	ContextUserIDKey contextKey = "user_id"
	ContextRoleKey   contextKey = "role"
)

// AuthMiddleware validates the JWT in the session_token cookie, sets userID/role in context, and allows only authenticated users.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, role, err := ParseJWTFromRequest(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ContextUserIDKey, userID)
		ctx = context.WithValue(ctx, ContextRoleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireRole wraps an HTTP handler, allowing only users with the required role (e.g., "admin").
func RequireRole(role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(ContextRoleKey)
		if val != role {
			http.Redirect(w, r, "/dashboard?error=not_"+role, http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ParseJWTFromRequest parses the JWT from the session_token cookie and returns the userID and role.
func ParseJWTFromRequest(r *http.Request) (string, string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return "", "", err
	}
	tokenStr := cookie.Value
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", "", errors.New("JWT_SECRET not set")
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return "", "", errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("invalid claims")
	}
	userID, ok := claims["user_id"].(string)
	if !ok || userID == "" {
		return "", "", errors.New("invalid user_id in token")
	}
	role := ""
	if r, ok := claims["role"].(string); ok {
		role = r
	}
	return userID, role, nil
}