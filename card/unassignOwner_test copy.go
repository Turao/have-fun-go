package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnassignOwnerUseCase(t *testing.T) {
	repository := NewInMemoryRepository()

	unassignOwner := UnassignOwnerUseCase{Repository: repository}

	card, err := unassignOwner.Execute(New().Id())
	assert.Nil(t, err)
	assert.NotNil(t, card)
}
