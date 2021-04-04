package card

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Card interface {
	ID() uuid.UUID
	HasOwner() bool
	OwnerID() uuid.UUID
	AssignOwner(ownerID uuid.UUID) error
	UnassignOwner() error
}

type card struct {
	id      uuid.UUID
	ownerID uuid.UUID
}

func (card card) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id      uuid.UUID `json:"id"`
		OwnerID uuid.UUID `json:"ownerID"`
	}{
		card.id,
		card.ownerID,
	})
}

func New() *card {
	return &card{uuid.New(), uuid.Nil}
}

func (c card) ID() uuid.UUID {
	return c.id
}

func (c card) OwnerID() uuid.UUID {
	return c.ownerID
}

func (c card) HasOwner() bool {
	return c.ownerID != uuid.Nil
}

func (c *card) AssignOwner(ownerID uuid.UUID) error {
	if c.HasOwner() {
		return errors.New("card already has an owner")
	}

	c.ownerID = ownerID
	return nil
}

func (c *card) UnassignOwner() error {
	if !c.HasOwner() {
		return errors.New("card has no owner")
	}

	c.ownerID = uuid.Nil
	return nil
}
