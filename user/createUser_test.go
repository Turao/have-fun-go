package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCase(t *testing.T) {
	repository := NewInMemoryRepository() // todo: mock this repository (how?)
	usecase := CreateUserUseCase{Repository: repository}

	user, err := usecase.Execute("dummy")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "dummy", user.Name())

	// todo: assert on repo calls
}
