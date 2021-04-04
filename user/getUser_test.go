package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// ! this should (and does) fail since it does not store
// ! an user before retrieving it (mock later)
// func TestGetUserUseCase(t *testing.T) {
// 	repository := NewInMemoryRepository() // todo: mock this repository (how?)
// 	usecase := GetUserUseCase{Repository: repository}

// 	userID := uuid.New()
// 	user, err := usecase.Execute(userID)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, user)
// 	// assert.Equal(t, "dummy", user.name)

// 	// todo: assert on repo calls
// }

func TestGetUserUseCaseNotExists(t *testing.T) {
	repository := NewInMemoryRepository() // todo: mock this repository (how?)
	usecase := GetUserUseCase{Repository: repository}

	userID := uuid.New()
	user, err := usecase.Execute(userID)
	assert.NotNil(t, err)
	assert.Nil(t, user)

	// todo: assert on repo calls
}
