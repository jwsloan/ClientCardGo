// Package http provides the EULA acceptance endpoint and static EULA content.
package http

import (
	"encoding/json"
	"net/http"
	"os"

	"api/internal/domain"
)

const (
	EULAHTMLPath = "EULA.html"
	EULATXTPath  = "EULA.txt"
)

type EULAHandler struct {
	UserRepo domain.UserRepository
}

type ctxKey int

const ctxKeyUserID ctxKey = iota

func (h *EULAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Serve EULA text (prefers HTML, falls back to plain text)
		var (
			data []byte
			err  error
		)
		if _, err = os.Stat(EULAHTMLPath); err == nil {
			data, err = os.ReadFile(EULAHTMLPath)
			if err == nil {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				w.Write(data)
				return
			}
		}
		data, err = os.ReadFile(EULATXTPath)
		if err != nil {
			http.Error(w, "could not load EULA", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write(data)
	case http.MethodPost:
		// Accept EULA for this user (safe context key and type assertion)
		userID, ok := r.Context().Value(ctxKeyUserID).(string)
		if !ok || userID == "" {
			http.Error(w, "not authenticated", http.StatusUnauthorized)
			return
		}
		if err := h.UserRepo.SetEULAAccepted(userID); err != nil {
			http.Error(w, "could not save EULA acceptance", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}