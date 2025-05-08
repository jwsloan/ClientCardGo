// Package domain defines the chat session and message model for profile interviews.
package domain

import (
	"time"
)

type ChatSession struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Completed bool
	Messages  []*ChatMessage
}

type ChatMessage struct {
	ID         string
	SessionID  string
	Sender     string // "user" or "system"
	Content    string
	CreatedAt  time.Time
}