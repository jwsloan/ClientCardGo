// Package http provides the EULA endpoints.
package http

import (
	"encoding/json"
	"net/http"
	"os"

	"api/internal/domain"
	"api/internal/adapter/auth"
)

type EULAHandler struct {
	UserRepo domain.UserRepository
}

func (h *EULAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _, err := auth.ParseJWTFromRequest(r)
	if err != nil || userID == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	switch r.Method {
	case http.MethodGet:
		// Serve the EULA markdown or HTML (could render or just send as text)
		content, err := os.ReadFile("eula.md")
		if err != nil {
			http.Error(w, "EULA not found", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/markdown")
		w.Write(content)
	case http.MethodPost:
		// Accept EULA
		err := h.UserRepo.SetEULAAccepted(userID, true)
		if err != nil {
			http.Error(w, "Could not record EULA acceptance", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}