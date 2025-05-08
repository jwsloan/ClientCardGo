// Package domain defines the InvitationRepository interface.
package domain

type InvitationWithUser struct {
	Invitation
	UserName  *string
	UserEmail *string
}

type InvitationRepository interface {
	// FindByToken returns the invitation for a given token (case-sensitive).
	FindByToken(token string) (*Invitation, error)
	// MarkUsed marks the invitation as used by the given user.
	MarkUsed(token, userID string) error
	// ListWithUserInfo returns all invitations including user name/email if used.
	ListWithUserInfo() ([]InvitationWithUser, error)
}