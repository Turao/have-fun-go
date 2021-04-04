package card

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAssignOwner_forCardWithoutOwner(t *testing.T) {
	card := New()
	ownerId := uuid.New()

	err := card.AssignOwner(ownerId)

	assert.Nil(t, err)
	assert.True(t, card.HasOwner())
	assert.Equal(t, ownerId, card.OwnerId())
}

func TestAssignOwner_forCardWithOwner(t *testing.T) {
	card := New()
	ownerId := uuid.New()
	card.AssignOwner(ownerId)

	err := card.AssignOwner(ownerId)

	assert.NotNil(t, err)
	assert.Equal(t, ownerId, card.OwnerId())
}

func TestUnassignOwner_forCardWithOwner(t *testing.T) {
	card := New()
	ownerId := uuid.New()
	card.AssignOwner(ownerId)

	err := card.UnassignOwner()

	assert.Nil(t, err)
	assert.False(t, card.HasOwner())
}

func TestUnassignOwner_forCardWithoutOwner(t *testing.T) {
	card := New()
	uuid.New()

	err := card.UnassignOwner()

	assert.NotNil(t, err)
	assert.False(t, card.HasOwner())
}
