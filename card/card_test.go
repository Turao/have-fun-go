package card

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAssignOwner(t *testing.T) {
	card := New()
	ownerId := uuid.New()

	err := card.AssignOwner(ownerId)

	assert.Nil(t, err)
	assert.Equal(t, ownerId, card.OwnerId())
}

func TestAssignOwnerFail(t *testing.T) {
	card := New()
	ownerId := uuid.New()
	card.AssignOwner(ownerId)

	err := card.AssignOwner(ownerId)

	assert.NotNil(t, err)
	assert.Equal(t, ownerId, card.OwnerId())
}

func TestUnassignOwner(t *testing.T) {
	card := New()
	ownerId := uuid.New()
	card.AssignOwner(ownerId)

	err := card.UnassignOwner()

	assert.Nil(t, err)
	assert.Equal(t, uuid.Nil, card.OwnerId())
}

func TestUnassignOwnerFail(t *testing.T) {
	card := New()
	uuid.New()

	err := card.UnassignOwner()

	assert.NotNil(t, err)
	assert.Equal(t, uuid.Nil, card.OwnerId())
}
