package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnassignOwnerUseCase(t *testing.T) {
	repository := NewInMemoryRepository()

	unassignOwner := UnassignOwnerUseCase{Repository: repository}

	card := New()
	card, err := unassignOwner.Execute(card.Id())
	assert.Nil(t, err)
	assert.NotNil(t, card)
}
