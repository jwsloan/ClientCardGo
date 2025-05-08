// Package chat provides an in-memory ChatRepository for development/testing.
package chat

import (
	"sync"
	"time"

	"api/internal/domain"
	"github.com/gofrs/uuid/v5"
)

type InMemoryChatRepo struct {
	sync.Mutex
	sessions map[string]*domain.ChatSession // sessionID -> session
	messages map[string][]*domain.ChatMessage // sessionID -> messages
}

func NewInMemoryChatRepo() *InMemoryChatRepo {
	return &InMemoryChatRepo{
		sessions: make(map[string]*domain.ChatSession),
		messages: make(map[string][]*domain.ChatMessage),
	}
}

func (r *InMemoryChatRepo) CreateSession(userID string) (*domain.ChatSession, error) {
	r.Lock()
	defer r.Unlock()
	id, _ := uuid.NewV7()
	session := &domain.ChatSession{
		ID:        id.String(),
		UserID:    userID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Completed: false,
		Messages:  []*domain.ChatMessage{},
	}
	r.sessions[session.ID] = session
	return session, nil
}

func (r *InMemoryChatRepo) GetSession(sessionID, userID string) (*domain.ChatSession, error) {
	r.Lock()
	defer r.Unlock()
	s, ok := r.sessions[sessionID]
	if !ok || s.UserID != userID {
		return nil, nil
	}
	return s, nil
}

func (r *InMemoryChatRepo) AddMessage(sessionID, sender, content string) (*domain.ChatMessage, error) {
	r.Lock()
	defer r.Unlock()
	id, _ := uuid.NewV7()
	msg := &domain.ChatMessage{
		ID:        id.String(),
		SessionID: sessionID,
		Sender:    sender,
		Content:   content,
		CreatedAt: time.Now().UTC(),
	}
	r.messages[sessionID] = append(r.messages[sessionID], msg)
	if s, ok := r.sessions[sessionID]; ok {
		s.Messages = append(s.Messages, msg)
		s.UpdatedAt = time.Now().UTC()
	}
	return msg, nil
}

func (r *InMemoryChatRepo) ListMessages(sessionID string) ([]*domain.ChatMessage, error) {
	r.Lock()
	defer r.Unlock()
	return r.messages[sessionID], nil
}

func (r *InMemoryChatRepo) MarkCompleted(sessionID string) error {
	r.Lock()
	defer r.Unlock()
	if s, ok := r.sessions[sessionID]; ok {
		s.Completed = true
		s.UpdatedAt = time.Now().UTC()
	}
	return nil
}