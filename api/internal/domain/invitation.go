// Package domain defines the Invitation model and its validation.
package domain

import (
	"time"
	"errors"
)

type InvitationStatus string

const (
	InvitationUnused  InvitationStatus = "unused"
	InvitationUsed    InvitationStatus = "used"
	InvitationExpired InvitationStatus = "expired"
)

type Invitation struct {
	ID        string
	Token     string
	Note      string
	Status    InvitationStatus
	CreatedAt time.Time
	UsedAt    *time.Time
	UserID    *string // who used the token, if any
}

var (
	ErrInvalidToken   = errors.New("invalid invitation token")
	ErrTokenUsed      = errors.New("invitation token already used")
	ErrTokenExpired   = errors.New("invitation token expired")
)