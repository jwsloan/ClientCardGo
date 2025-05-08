// Package usecase implements business logic for profile interview chat.
package usecase

import (
	"api/internal/domain"
	"time"

	"github.com/gofrs/uuid/v5"
)

type Chat struct {
	Repo domain.ChatRepository
}

func NewChat(repo domain.ChatRepository) *Chat {
	return &Chat{Repo: repo}
}

func (c *Chat) StartSession(userID string) (*domain.ChatSession, error) {
	return c.Repo.CreateSession(userID)
}

func (c *Chat) SendMessage(sessionID, userID, content string) (*domain.ChatMessage, error) {
	// Find session and validate it belongs to user
	session, err := c.Repo.GetSession(sessionID, userID)
	if err != nil || session == nil {
		return nil, err
	}
	// Save user message
	msg, err := c.Repo.AddMessage(sessionID, "user", content)
	if err != nil {
		return nil, err
	}
	session.UpdatedAt = time.Now().UTC()
	return msg, nil
}

func (c *Chat) ListMessages(sessionID, userID string) ([]*domain.ChatMessage, error) {
	session, err := c.Repo.GetSession(sessionID, userID)
	if err != nil || session == nil {
		return nil, err
	}
	return c.Repo.ListMessages(sessionID)
}

func (c *Chat) MarkCompleted(sessionID, userID string) error {
	return c.Repo.MarkCompleted(sessionID)
}