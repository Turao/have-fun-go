package users

import (
	"errors"

	"github.com/google/uuid"
)

type Repository interface {
	GetUser(uuid.UUID) (*user, error)
	CreateUser(user *user) (*user, error)
}

type InMemoryRepository struct {
	inMemoryDatabase map[uuid.UUID]*user
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{make(map[uuid.UUID]*user)}
}

func (r *InMemoryRepository) GetUser(userId uuid.UUID) (*user, error) {
	user, found := r.inMemoryDatabase[userId]
	if !found {
		return nil, errors.New("user does not exist")
	}
	return user, nil
}

func (r *InMemoryRepository) CreateUser(user *user) (*user, error) {
	_, found := r.inMemoryDatabase[user.id]
	if found {
		return nil, errors.New("User has already been stored")
	}
	r.inMemoryDatabase[user.id] = user
	return user, nil
}