package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUserUseCase(t *testing.T) {
	repository := NewInMemoryRepository() // todo: mock this repository (how?)
	usecase := GetUserUseCase{Repository: repository}

	userId := uuid.New()
	user, err := usecase.Execute(userId)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	// assert.Equal(t, "dummy", user.name)

	// todo: assert on repo calls
}

func TestGetUserUseCaseNotExists(t *testing.T) {
	repository := NewInMemoryRepository() // todo: mock this repository (how?)
	usecase := GetUserUseCase{Repository: repository}

	userId := uuid.New()
	user, err := usecase.Execute(userId)
	assert.NotNil(t, err)
	assert.Nil(t, user)

	// todo: assert on repo calls
}