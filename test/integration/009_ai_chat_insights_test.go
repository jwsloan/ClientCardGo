package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"api/internal/adapter/http"
	"api/internal/adapter/chat"
	"api/internal/usecase"

	"github.com/stretchr/testify/assert"
)

func TestProfileInterviewChatFlow(t *testing.T) {
	repo := chat.NewInMemoryChatRepo()
	uc := usecase.NewChat(repo)
	handler := &http.ChatHandler{UC: uc}

	server := httptest.NewServer(handler)
	defer server.Close()

	// Simulate starting a chat session (GET /chat)
	req, _ := http.NewRequest("GET", server.URL, nil)
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "dummy.jwt.token"})
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var session struct {
		ID string `json:"id"`
	}
	json.NewDecoder(resp.Body).Decode(&session)
	assert.NotEmpty(t, session.ID)

	// Send a message (POST /chat)
	payload := map[string]string{
		"session_id": session.ID,
		"message":    "Hello, I'm a contractor.",
	}
	b, _ := json.Marshal(payload)
	req2, _ := http.NewRequest("POST", server.URL, bytes.NewReader(b))
	req2.Header.Set("Content-Type", "application/json")
	req2.AddCookie(&http.Cookie{Name: "session_token", Value: "dummy.jwt.token"})
	resp2, err := http.DefaultClient.Do(req2)
	assert.NoError(t, err)
	defer resp2.Body.Close()
	assert.Equal(t, http.StatusOK, resp2.StatusCode)

	var reply map[string]interface{}
	json.NewDecoder(resp2.Body).Decode(&reply)
	assert.Equal(t, "user", reply["sender"])
	assert.Equal(t, "Hello, I'm a contractor.", reply["content"])
}