package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"api/internal/adapter/http"
	"api/internal/adapter/invitation"
	"api/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestAdminInvitationsListingShowsUserDetails(t *testing.T) {
	repo := invitation.NewInMemoryInvitationRepo()
	// Add an unused and a used invitation
	repo.invites["TOKEN1"] = &domain.Invitation{
		Token:  "TOKEN1",
		Status: domain.InvitationUnused,
	}
	userID := "user-uuid"
	repo.invites["TOKEN2"] = &domain.Invitation{
		Token:  "TOKEN2",
		Status: domain.InvitationUsed,
		UserID: &userID,
	}
	repo.SetUser(userID, "Alice Admin", "alice@example.com")

	handler := &http.AdminInvitationsHandler{Invitations: repo}
	server := httptest.NewServer(handler)
	defer server.Close()

	req, _ := http.NewRequest("GET", server.URL, nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var data []domain.InvitationWithUser
	json.NewDecoder(resp.Body).Decode(&data)
	assert.Len(t, data, 2)
	for _, inv := range data {
		if inv.Token == "TOKEN2" {
			assert.Equal(t, domain.InvitationUsed, inv.Status)
			assert.Equal(t, "Alice Admin", *inv.UserName)
			assert.Equal(t, "alice@example.com", *inv.UserEmail)
		}
		if inv.Token == "TOKEN1" {
			assert.Equal(t, domain.InvitationUnused, inv.Status)
			assert.Nil(t, inv.UserName)
			assert.Nil(t, inv.UserEmail)
		}
	}
}