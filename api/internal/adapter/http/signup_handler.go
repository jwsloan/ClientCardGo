// Package http provides HTTP handlers for the API.
package http

import (
	"encoding/json"
	"net/http"

	"api/internal/usecase"
)

// SignupHandler handles POST /signup requests.
type SignupHandler struct {
	SignupUC *usecase.Signup
}

type signupRequest struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	InvitationToken string `json:"invitation_token"`
}

type signupResponse struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

import (
	"github.com/rs/zerolog/log"
	"context"
)

func (h *SignupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req signupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	out, err := h.SignupUC.Register(r.Context(), usecase.SignupInput{
		Email:           req.Email,
		Name:            req.Name,
		Password:        req.Password,
		InvitationToken: req.InvitationToken,
	})
	if err != nil {
		log.Error().Err(err).Msg("signup failed")
		switch err {
		case usecase.ErrEmailTaken:
			http.Error(w, "email already registered", http.StatusConflict)
		case domain.ErrInvalidToken:
			http.Error(w, "invalid invitation token", http.StatusBadRequest)
		case domain.ErrTokenUsed:
			http.Error(w, "invitation token already used", http.StatusBadRequest)
		case domain.ErrTokenExpired:
			http.Error(w, "invitation token expired", http.StatusBadRequest)
		default:
			http.Error(w, "signup failed", http.StatusBadRequest)
		}
		return
	}

	resp := signupResponse{
		UserID: out.UserID,
		Name:   out.Name,
	}
	w.Header().Set("Content-Type", "application/json")
	// After signup, user always needs to complete interview (unless admin, but signup is not for admin)
	json.NewEncoder(w).Encode(map[string]string{
		"redirect": "/profile-interview",
		"message":  "Welcome, " + out.Name + "! Please complete your profile interview.",
	})
}