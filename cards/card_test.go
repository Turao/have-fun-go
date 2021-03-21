package cards

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAssignOwner(t *testing.T) {
	card := New()
	ownerId := uuid.New()

	newCard, err := card.AssignOwner(ownerId)

	assert.Nil(t, err)
	assert.Equal(t, ownerId, newCard.ownerId)
}

func TestAssignOwnerFail(t *testing.T) {
	card := New()
	ownerId := uuid.New()
	card.AssignOwner(ownerId)

	newCard, err := card.AssignOwner(ownerId)

	assert.NotNil(t, err)
	assert.Equal(t, ownerId, newCard.ownerId)
}

func TestUnassignOwner(t *testing.T) {
	card := New()
	ownerId := uuid.New()
	card.AssignOwner(ownerId)

	newCard, err := card.UnassignOwner()

	assert.Nil(t, err)
	assert.Equal(t, uuid.Nil, newCard.ownerId)
}

func TestUnassignOwnerFail(t *testing.T) {
	card := New()
	uuid.New()

	newCard, err := card.UnassignOwner()

	assert.NotNil(t, err)
	assert.Equal(t, uuid.Nil, newCard.ownerId)
}
