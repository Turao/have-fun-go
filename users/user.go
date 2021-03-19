package users

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type User interface {
	Id() uuid.UUID
	Name() string
	Cards() []uuid.UUID
}

type user struct {
	id    uuid.UUID
	name  string
	cards []uuid.UUID
}

func (user user) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id    uuid.UUID   `json:"id"`
		Name  string      `json:"name"`
		Cards []uuid.UUID `json:"cards"`
	}{
		user.id,
		user.name,
		user.cards,
	})
}

func New(name string) user {
	user := user{uuid.New(), name, make([]uuid.UUID, 0)}

	return user
}

func (u *user) Id() uuid.UUID {
	return u.id
}

func (u *user) Name() string {
	return u.name
}

func (u *user) Cards() []uuid.UUID {
	return u.cards
}

func (u *user) AddCard(cardId uuid.UUID) (*user, error) {
	for _, c := range u.cards {
		if c == cardId {
			return u, errors.New("User already has this card")
		}
	}

	u.cards = append(u.cards, cardId)
	return u, nil
}
