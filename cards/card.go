package cards

import (
	"errors"

	"github.com/google/uuid"
)

type Card interface {
	Id()
	OwnerId()
}

type card struct {
	id      uuid.UUID `json:"id"`
	ownerId uuid.UUID `json:"ownerId"`
}

func New() card {
	return card{uuid.New(), uuid.Nil}
}

func (c *card) Id() uuid.UUID {
	return c.id
}

func (c *card) OwnerId() uuid.UUID {
	return c.ownerId
}

func (c *card) AssignOwner(ownerId uuid.UUID) (*card, error) {
	if c.ownerId != uuid.Nil {
		return c, errors.New("Card already has an owner")
	}

	c.ownerId = ownerId
	return c, nil
}

func (c *card) UnassignOwner() (*card, error) {
	if c.ownerId == uuid.Nil {
		return c, errors.New("Card has no owner")
	}

	c.ownerId = uuid.Nil
	return c, nil
}
