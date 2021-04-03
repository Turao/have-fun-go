package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func givenUserExists(t *testing.T, repo Repository) *user {
	user, err := repo.CreateUser(New("dummy"))
	if err != nil {
		t.Fatal("failed on helper function")
	}
	return user
}

func TestAddCardUseCase(t *testing.T) {
	repository := NewInMemoryRepository() // todo: mock this repository (how?)
	user := givenUserExists(t, repository)

	usecase := AddCardUseCase{Repository: repository}

	cardId := uuid.New()

	user, err := usecase.Execute(user.id, cardId)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Contains(t, user.cards, cardId)

	// todo: assert on repo calls
}
