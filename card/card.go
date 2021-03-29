package card

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Card interface {
	Id() uuid.UUID
	OwnerId() uuid.UUID
	AssignOwner(ownerId uuid.UUID) (*Card, error)
	UnassignOwner() (*Card, error)
}

type card struct {
	id      uuid.UUID
	ownerId uuid.UUID
}

func (card card) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id      uuid.UUID `json:"id"`
		OwnerId uuid.UUID `json:"ownerId"`
	}{
		card.id,
		card.ownerId,
	})
}

func New() *card {
	return &card{uuid.New(), uuid.Nil}
}

func (c card) Id() uuid.UUID {
	return c.id
}

func (c card) OwnerId() uuid.UUID {
	return c.ownerId
}

func (c *card) AssignOwner(ownerId uuid.UUID) (*card, error) {
	if c.ownerId != uuid.Nil {
		return c, errors.New("card already has an owner")
	}

	c.ownerId = ownerId
	return c, nil
}

func (c *card) UnassignOwner() (*card, error) {
	if c.ownerId == uuid.Nil {
		return c, errors.New("card has no owner")
	}

	c.ownerId = uuid.Nil
	return c, nil
}
