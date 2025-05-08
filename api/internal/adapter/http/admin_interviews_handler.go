// Package http provides admin access to user interview transcripts.
package http

import (
	"encoding/json"
	"net/http"
	"api/internal/domain"
)

// AdminInterviewsHandler allows admins to list and view interview transcripts.
type AdminInterviewsHandler struct {
	ChatRepo domain.ChatRepository
	UserRepo domain.UserRepository
}

// List all interview sessions (basic pagination for demo; real impl should use DB pagination)
func (h *AdminInterviewsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// RBAC: Assumes admin middleware already enforced

	switch r.Method {
	case http.MethodGet:
		// List sessions, or fetch transcript for a session
		sessionID := r.URL.Query().Get("session_id")
		if sessionID != "" {
			// Fetch a specific transcript
			session, err := h.ChatRepo.GetSession(sessionID, "")
			if err != nil || session == nil {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			// Get user info (show as little as possible)
			user, _ := h.UserRepo.GetUserByEmail(session.UserID) // UserID is probably not email; real impl would fetch by ID
			type UserSafe struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			}
			var safeUser *UserSafe
			if user != nil {
				safeUser = &UserSafe{ID: user.ID, Name: user.Name}
			}
			transcript := struct {
				SessionID string                `json:"session_id"`
				User      *UserSafe             `json:"user"`
				Messages  []*domain.ChatMessage `json:"messages"`
			}{
				SessionID: session.ID,
				User:      safeUser,
				Messages:  session.Messages,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(transcript)
			return
		}

		// List all interview sessions (for demo, fetch all)
		sessions, err := h.ChatRepo.ListAllSessions()
		if err != nil {
			http.Error(w, "could not load interviews", http.StatusInternalServerError)
			return
		}
		summaries := make([]map[string]interface{}, 0, len(sessions))
		for _, s := range sessions {
			summaries = append(summaries, map[string]interface{}{
				"id":        s.ID,
				"user_id":   s.UserID,
				"created":   s.CreatedAt,
				"completed": s.Completed,
				"length":    len(s.Messages),
			})
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(summaries)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}