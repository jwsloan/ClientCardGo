// Package http provides HTTP handlers for admin dashboard.
package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// AdminHandler serves the /admin dashboard for admins only.
type AdminHandler struct{}

func (h *AdminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userRole := getRoleFromJWT(r)
	if userRole != "admin" {
		http.Redirect(w, r, "/dashboard?error=not_admin", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome, Admin",
	})
}

// getRoleFromJWT parses JWT from the cookie and returns the user's role or "".
func getRoleFromJWT(r *http.Request) string {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return ""
	}
	tokenStr := cookie.Value
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// This should use your actual JWT secret
		return []byte("your-secret-key"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		role, _ := claims["role"].(string)
		return role
	}
	return ""
}