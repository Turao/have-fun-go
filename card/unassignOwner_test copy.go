package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnassignOwnerUseCase(t *testing.T) {
	repository := NewInMemoryRepository()
	card := givenCardExists(t, repository)
	card = givenCardHasOwner(t, repository, card)

	unassignOwner := UnassignOwnerUseCase{Repository: repository}

	card, err := unassignOwner.Execute(card.Id())
	assert.Nil(t, err)
	assert.NotNil(t, card)
	assert.False(t, card.HasOwner())
}
