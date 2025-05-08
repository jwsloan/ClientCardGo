// Package domain defines repository interfaces for chat sessions.
package domain

type ChatRepository interface {
	// CreateSession creates a new chat session for a user.
	CreateSession(userID string) (*ChatSession, error)
	// GetSession retrieves a chat session by session ID and user ID.
	GetSession(sessionID, userID string) (*ChatSession, error)
	// AddMessage adds a message to the session.
	AddMessage(sessionID, sender, content string) (*ChatMessage, error)
	// ListMessages returns all messages for a session.
	ListMessages(sessionID string) ([]*ChatMessage, error)
	// MarkCompleted marks a session as completed.
	MarkCompleted(sessionID string) error
}