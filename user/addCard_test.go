package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddCardUseCase(t *testing.T) {
	repository := NewInMemoryRepository() // todo: mock this repository (how?)
	user := givenUserExists(t, repository)

	usecase := AddCardUseCase{Repository: repository}

	cardID := uuid.New()

	user, err := usecase.Execute(user.ID(), cardID)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Contains(t, user.Cards(), cardID)

	// todo: assert on repo calls
}
