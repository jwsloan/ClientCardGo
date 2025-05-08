package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"api/internal/adapter/http"
	"api/internal/domain"
	"api/internal/usecase"

	"github.com/stretchr/testify/assert"
)

type testUserRepo struct {
	users map[string]*domain.User
}

func newTestUserRepo() *testUserRepo {
	return &testUserRepo{users: map[string]*domain.User{}}
}

func (r *testUserRepo) CreateUser(u *domain.User) (string, error) {
	if _, exists := r.users[u.Email]; exists {
		return "", usecase.ErrEmailTaken
	}
	// enforce allowed roles
	if u.Role != "member" && u.Role != "admin" {
		return "", domain.ErrInvalidRole
	}
	r.users[u.Email] = u
	return u.ID, nil
}

func (r *testUserRepo) GetUserByEmail(email string) (*domain.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, nil
	}
	return u, nil
}
	// Validate role
	if u.Role != "member" && u.Role != "admin" {
		return "", domain.ErrInvalidRole
	}
	r.users[u.Email] = u
	return u.ID, nil
}

func (r *testUserRepo) GetUserByEmail(email string) (*domain.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, nil
	}
	return u, nil
}

func TestSignupHandler(t *testing.T) {
	repo := newTestUserRepo()
	signupUC := usecase.NewSignup(repo)
	handler := &http.SignupHandler{SignupUC: signupUC}
	server := httptest.NewServer(handler)
	defer server.Close()

	tests := []struct {
		name           string
		payload        map[string]string
		wantStatus     int
		wantInResponse string
	}{
		{
			name: "valid signup",
			payload: map[string]string{
				"email":    "test@example.com",
				"name":     "Test User",
				"password": "Password1",
			},
			wantStatus:     http.StatusOK,
			wantInResponse: `"user_id"`,
		},
		{
			name: "uuidv7 id format",
			payload: map[string]string{
				"email":    "uuidv7@example.com",
				"name":     "UUID User",
				"password": "Password1",
			},
			wantStatus:     http.StatusOK,
			wantInResponse: `"user_id"`,
		},
		{
			name: "role defaults to member",
			payload: map[string]string{
				"email":    "roledefault@example.com",
				"name":     "Role Default",
				"password": "Password1",
			},
			wantStatus:     http.StatusOK,
			wantInResponse: `"user_id"`,
		},
		{
			name: "invalid role rejected",
			payload: map[string]string{
				"email":    "badrole@example.com",
				"name":     "Bad Role",
				"password": "Password1",
				"role":     "superuser",
			},
			wantStatus:     http.StatusBadRequest,
			wantInResponse: "invalid user role",
		},
		{
			name: "duplicate email",
			payload: map[string]string{
				"email":    "test@example.com",
				"name":     "Another User",
				"password": "Password1",
			},
			wantStatus:     http.StatusConflict,
			wantInResponse: "email already registered",
		},
		{
			name: "weak password",
			payload: map[string]string{
				"email":    "weakpass@example.com",
				"name":     "Weak User",
				"password": "short",
			},
			wantStatus:     http.StatusBadRequest,
			wantInResponse: "password does not meet strength requirements",
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "not-an-email",
				"name":     "No Email",
				"password": "Password1",
			},
			wantStatus:     http.StatusBadRequest,
			wantInResponse: "invalid email address",
		},
		{
			name: "missing name",
			payload: map[string]string{
				"email":    "noname@example.com",
				"name":     "",
				"password": "Password1",
			},
			wantStatus:     http.StatusBadRequest,
			wantInResponse: "name is required",
		},
	}

	// Arrange - create user for duplicate test
	_, _ = repo.CreateUser(&domain.User{
		ID:       "existing-id",
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "hashedpw",
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.payload)
			resp, err := http.Post(server.URL, "application/json", bytes.NewReader(b))
			assert.NoError(t, err)
			defer resp.Body.Close()
			assert.Equal(t, tt.wantStatus, resp.StatusCode)
			buf := new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			assert.Contains(t, buf.String(), tt.wantInResponse)

			// For role defaults to member, check that the created user has role "member"
			if tt.name == "role defaults to member" && resp.StatusCode == http.StatusOK {
				created := repo.users["roledefault@example.com"]
				assert.Equal(t, "member", created.Role)
			}
		})
	}
}