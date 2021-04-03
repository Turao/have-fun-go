package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListUsersUseCase(t *testing.T) {
	repository := NewInMemoryRepository() // todo: mock this repository (how?)
	usecase := ListUsersUseCase{Repository: repository}

	users, err := usecase.Execute()
	assert.Nil(t, err)
	assert.NotNil(t, users)

	// todo: assert on repo calls
}
