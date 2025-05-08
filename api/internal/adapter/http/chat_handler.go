// Package http provides HTTP handlers for chat-based interview/profile insights.
package http

import (
	"encoding/json"
	"net/http"

	"api/internal/usecase"
	"api/internal/adapter/auth"
)

type ChatHandler struct {
	UC *usecase.Chat
}

func (h *ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _, err := auth.ParseJWTFromRequest(r)
	if err != nil || userID == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	switch r.Method {
	case http.MethodPost:
		var req struct {
			SessionID string `json:"session_id"`
			Message   string `json:"message"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		msg, err := h.UC.SendMessage(req.SessionID, userID, req.Message)
		if err != nil {
			http.Error(w, "could not send message", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(msg)
	case http.MethodGet:
		sessionID := r.URL.Query().Get("session_id")
		if sessionID == "" {
			// Create new session
			session, err := h.UC.StartSession(userID)
			if err != nil {
				http.Error(w, "could not start session", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(session)
			return
		}
		// List messages for session
		messages, err := h.UC.ListMessages(sessionID, userID)
		if err != nil {
			http.Error(w, "could not fetch messages", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(messages)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// ChatCompleteHandler marks a chat session as completed (for "Finish Interview" button).
func (h *ChatHandler) Complete(w http.ResponseWriter, r *http.Request) {
	userID, _, err := auth.ParseJWTFromRequest(r)
	if err != nil || userID == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		SessionID string `json:"session_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if err := h.UC.MarkCompleted(req.SessionID, userID); err != nil {
		http.Error(w, "could not complete interview", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}