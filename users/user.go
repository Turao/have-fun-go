package users

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type User interface {
	Id() uuid.UUID
	Name() string
	Cards() map[uuid.UUID]bool
	AddCard(cardId uuid.UUID) (*User, error)
	RemoveCard(cardId uuid.UUID) (*User, error)
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

func (u user) Id() uuid.UUID {
	return u.id
}

func (u user) Name() string {
	return u.name
}

func (u user) Cards() map[uuid.UUID]bool {
	return u.cards
}

func (u *user) AddCard(cardId uuid.UUID) (*user, error) {
	found := u.cards[cardId]
	if found {
		return u, errors.New("user already has this card")
	}

	u.cards[cardId] = true
	return u, nil
}

func (u *user) RemoveCard(cardId uuid.UUID) (*user, error) {
	found := u.cards[cardId]
	if !found {
		return u, errors.New("user does not have this card")
	}

	delete(u.cards, cardId)
	return u, nil
}
