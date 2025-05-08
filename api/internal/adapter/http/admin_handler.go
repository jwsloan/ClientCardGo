// Package http provides HTTP handlers for admin dashboard.
package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// AdminHandler serves the /admin dashboard for admins only.
package http

import (
	"encoding/json"
	"net/http"
)

// AdminHandler serves the /admin dashboard for admins only.
type AdminHandler struct{}

func (h *AdminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome, Admin",
	})
}