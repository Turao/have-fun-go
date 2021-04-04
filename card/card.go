package card

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Card interface {
	Id() uuid.UUID
	HasOwner() bool
	OwnerId() uuid.UUID
	AssignOwner(ownerId uuid.UUID) error
	UnassignOwner() error
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

func (c card) HasOwner() bool {
	return c.ownerId != uuid.Nil
}

func (c *card) AssignOwner(ownerId uuid.UUID) error {
	if c.HasOwner() {
		return errors.New("card already has an owner")
	}

	c.ownerId = ownerId
	return nil
}

func (c *card) UnassignOwner() error {
	if !c.HasOwner() {
		return errors.New("card has no owner")
	}

	c.ownerId = uuid.Nil
	return nil
}
