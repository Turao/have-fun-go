package card

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAssignOwnerUseCase(t *testing.T) {
	repository := NewInMemoryRepository()
	card := givenCardExists(t, repository)

	assignOwner := AssignOwnerUseCase{Repository: repository}

	ownerID := uuid.New()
	card, err := assignOwner.Execute(card.ID(), ownerID)

	assert.Nil(t, err)
	assert.NotNil(t, card)
	assert.True(t, card.HasOwner())
}
