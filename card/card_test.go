package card

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAssignOwner_forCardWithoutOwner(t *testing.T) {
	card := New()
	ownerID := uuid.New()

	err := card.AssignOwner(ownerID)

	assert.Nil(t, err)
	assert.True(t, card.HasOwner())
	assert.Equal(t, ownerID, card.OwnerID())
}

func TestAssignOwner_forCardWithOwner(t *testing.T) {
	card := New()
	ownerID := uuid.New()
	card.AssignOwner(ownerID)

	err := card.AssignOwner(ownerID)

	assert.NotNil(t, err)
	assert.Equal(t, ownerID, card.OwnerID())
}

func TestUnassignOwner_forCardWithOwner(t *testing.T) {
	card := New()
	ownerID := uuid.New()
	card.AssignOwner(ownerID)

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
