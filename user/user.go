package user

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type User interface {
	ID() uuid.UUID
	Name() string
	Cards() map[uuid.UUID]bool
	AddCard(cardID uuid.UUID) error
	RemoveCard(cardID uuid.UUID) error
}

type user struct {
	id    uuid.UUID
	name  string
	cards map[uuid.UUID]bool
}

func (user user) MarshalJSON() ([]byte, error) {
	cards := make([]uuid.UUID, 0, len(user.cards))
	for card := range user.cards {
		cards = append(cards, card)
	}

	return json.Marshal(struct {
		Id    uuid.UUID   `json:"id"`
		Name  string      `json:"name"`
		Cards []uuid.UUID `json:"cards"`
	}{
		user.id,
		user.name,
		cards,
	})
}

func New(name string) *user {
	return &user{uuid.New(), name, make(map[uuid.UUID]bool)}
}

func (u user) ID() uuid.UUID {
	return u.id
}

func (u user) Name() string {
	return u.name
}

func (u user) Cards() map[uuid.UUID]bool {
	return u.cards
}

func (u *user) AddCard(cardID uuid.UUID) error {
	found := u.cards[cardID]
	if found {
		return errors.New("user already has this card")
	}

	u.cards[cardID] = true
	return nil
}

func (u *user) RemoveCard(cardID uuid.UUID) error {
	found := u.cards[cardID]
	if !found {
		return errors.New("user does not have this card")
	}

	delete(u.cards, cardID)
	return nil
}
