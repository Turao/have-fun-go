package user

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	repo := NewInMemoryRepository()
	user := New("dummy")

	stored, err := repo.CreateUser(user)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user, stored)
}

func TestCreateUserAlreadyExists(t *testing.T) {
	repo := NewInMemoryRepository()
	user := New("dummy")

	repo.CreateUser(user)
	stored, err := repo.CreateUser(user)

	assert.NotNil(t, err)
	assert.Nil(t, stored)
}

func TestGetUser(t *testing.T) {
	repo := NewInMemoryRepository()
	user := New("dummy")
	repo.CreateUser(user)

	found, err := repo.GetUser(user.id)
	assert.Nil(t, err)
	assert.Equal(t, user, found)
}

func TestGetUserDoesNotExist(t *testing.T) {
	repo := NewInMemoryRepository()
	user := New("dummy")

	found, err := repo.GetUser(user.id)
	assert.NotNil(t, err)
	assert.Nil(t, found)
}

func TestGetUsers(t *testing.T) {
	repo := NewInMemoryRepository()
	user0 := New("dummy-0")
	repo.CreateUser(user0)

	user1 := New("dummy-1")
	repo.CreateUser(user1)

	users, err := repo.GetUsers()
	log.Println(users)
	assert.Nil(t, err)

	assert.Contains(t, users, *user0)
	assert.Contains(t, users, *user1)
}

func TestUpdateUser(t *testing.T) {
	repo := NewInMemoryRepository()

	user, _ := repo.CreateUser(New("dummy-0"))
	user.name = "dummy-1"
	updated, err := repo.UpdateUser(user)

	assert.Nil(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, "dummy-1", updated.name)
}

func TestUpdateUserDoesNotExist(t *testing.T) {
	repo := NewInMemoryRepository()
	user := New("dummy")

	updated, err := repo.UpdateUser(user)

	assert.NotNil(t, err)
	assert.Nil(t, updated)
}
