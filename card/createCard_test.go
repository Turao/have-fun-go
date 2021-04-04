package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCardUseCase(t *testing.T) {
	repository := NewInMemoryRepository() // todo: mock this repository (how?)
	usecase := CreateCardUseCase{Repository: repository}

	card, err := usecase.Execute()
	assert.Nil(t, err)
	assert.NotNil(t, card)

	// todo: assert on repo calls
}
