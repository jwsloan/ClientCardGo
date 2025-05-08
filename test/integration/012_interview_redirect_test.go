package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"api/internal/adapter/http"
	"api/internal/domain"
	"api/internal/usecase"

	"github.com/stretchr/testify/assert"
)

type testUserRepoWithInterview struct {
	users map[string]*domain.User
}

func newTestUserRepoWithInterview() *testUserRepoWithInterview {
	return &testUserRepoWithInterview{users: map[string]*domain.User{}}
}
func (r *testUserRepoWithInterview) CreateUser(u *domain.User) (string, error) {
	r.users[u.Email] = u
	return u.ID, nil
}
func (r *testUserRepoWithInterview) GetUserByEmail(email string) (*domain.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, nil
	}
	return u, nil
}
func (r *testUserRepoWithInterview) SetInterviewComplete(userID string, complete bool) error {
	for _, u := range r.users {
		if u.ID == userID {
			u.InterviewComplete = complete
			return nil
		}
	}
	return nil
}

func TestInterviewRedirectLogic(t *testing.T) {
	userRepo := newTestUserRepoWithInterview()
	// Add a member who has NOT completed interview
	user := &domain.User{
		ID:                "u1",
		Email:             "member@example.com",
		Name:              "Member",
		Password:          "hash",
		Role:              "member",
		InterviewComplete: false,
		CreatedAt:         time.Now(),
	}
	userRepo.users[user.Email] = user

	// Add a member who HAS completed interview
	userDone := &domain.User{
		ID:                "u2",
		Email:             "done@example.com",
		Name:              "Done",
		Password:          "hash",
		Role:              "member",
		InterviewComplete: true,
		CreatedAt:         time.Now(),
	}
	userRepo.users[userDone.Email] = userDone

	// Add an admin
	admin := &domain.User{
		ID:                "a1",
		Email:             "admin@example.com",
		Name:              "Admin",
		Password:          "hash",
		Role:              "admin",
		InterviewComplete: false,
		CreatedAt:         time.Now(),
	}
	userRepo.users[admin.Email] = admin

	loginUC := usecase.NewLogin(userRepo, nil)
	loginHandler := &http.LoginHandler{LoginUC: loginUC}

	server := httptest.NewServer(loginHandler)
	defer server.Close()

	tests := []struct {
		email       string
		wantRedirect string
	}{
		{email: user.Email, wantRedirect: "/profile-interview"},
		{email: userDone.Email, wantRedirect: "/dashboard"},
		{email: admin.Email, wantRedirect: "/admin"},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			payload := map[string]string{"email": tt.email, "password": "any"}
			b, _ := json.Marshal(payload)
			resp, err := http.Post(server.URL, "application/json", bytes.NewReader(b))
			assert.NoError(t, err)
			defer resp.Body.Close()
			var out map[string]string
			json.NewDecoder(resp.Body).Decode(&out)
			assert.Equal(t, tt.wantRedirect, out["redirect"])
		})
	}
}

func TestCompleteInterviewSetsFlag(t *testing.T) {
	userRepo := newTestUserRepoWithInterview()
	user := &domain.User{
		ID:                "u3",
		Email:             "flow@example.com",
		Name:              "Flow",
		Password:          "hash",
		Role:              "member",
		InterviewComplete: false,
		CreatedAt:         time.Now(),
	}
	userRepo.users[user.Email] = user

	chatUC := usecase.NewChat(nil)
	chatUC.SetUserRepository(userRepo)
	sessionID := "sess1"
	// Simulate backend: MarkCompleted should set InterviewComplete
	err := chatUC.MarkCompleted(sessionID, user.ID)
	assert.NoError(t, err)
	assert.True(t, user.InterviewComplete)
}