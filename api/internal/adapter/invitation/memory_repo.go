// Package invitation provides an in-memory InvitationRepository for testing/demo.
package invitation

import (
	"api/internal/domain"
	"time"
)

type InMemoryInvitationRepo struct {
	invites map[string]*domain.Invitation
	users   map[string]struct {
		Name  string
		Email string
	}
}

func NewInMemoryInvitationRepo() *InMemoryInvitationRepo {
	return &InMemoryInvitationRepo{
		invites: map[string]*domain.Invitation{},
		users:   map[string]struct{ Name, Email string }{},
	}
}

func (r *InMemoryInvitationRepo) FindByToken(token string) (*domain.Invitation, error) {
	if inv, ok := r.invites[token]; ok {
		return inv, nil
	}
	return nil, nil
}

func (r *InMemoryInvitationRepo) MarkUsed(token, userID string) error {
	inv, ok := r.invites[token]
	if !ok {
		return domain.ErrInvalidToken
	}
	if inv.Status == domain.InvitationUsed {
		return domain.ErrTokenUsed
	}
	now := time.Now()
	inv.Status = domain.InvitationUsed
	inv.UsedAt = &now
	inv.UserID = &userID
	return nil
}

func (r *InMemoryInvitationRepo) ListWithUserInfo() ([]domain.InvitationWithUser, error) {
	list := []domain.InvitationWithUser{}
	for _, inv := range r.invites {
		var name, email *string
		if inv.UserID != nil {
			u, ok := r.users[*inv.UserID]
			if ok {
				name = &u.Name
				email = &u.Email
			}
		}
		list = append(list, domain.InvitationWithUser{
			Invitation: *inv,
			UserName:   name,
			UserEmail:  email,
		})
	}
	return list, nil
}

// For test/demo: add a user record for a given user ID
func (r *InMemoryInvitationRepo) SetUser(id, name, email string) {
	r.users[id] = struct {
		Name  string
		Email string
	}{Name: name, Email: email}
}