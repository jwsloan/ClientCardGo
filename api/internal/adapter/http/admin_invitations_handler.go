// Package http provides the admin invitation list handler.
package http

import (
	"encoding/json"
	"net/http"

	"api/internal/domain"
	"api/internal/adapter/auth"
)

type AdminInvitationsHandler struct {
	Invitations domain.InvitationRepository
}

func (h *AdminInvitationsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Only allow GET for listing.
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Auth + admin role enforced by middleware.
	invs, err := h.Invitations.ListWithUserInfo()
	if err != nil {
		http.Error(w, "could not load invitations", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invs)
}