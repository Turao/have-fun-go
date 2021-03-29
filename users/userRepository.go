package users

import (
	"errors"

	"github.com/google/uuid"
)

type Repository struct {
	inMemoryDatabase map[uuid.UUID]*user
}

func NewRepository() *Repository {
	return &Repository{make(map[uuid.UUID]*user)}
}

func (r *Repository) GetUser(userId uuid.UUID) (*user, error) {
	user, found := r.inMemoryDatabase[userId]
	if !found {
		return nil, errors.New("user does not exist")
	}
	return user, nil
}

func (r *Repository) CreateUser(user *user) (*user, error) {
	_, found := r.inMemoryDatabase[user.id]
	if found {
		return nil, errors.New("User has already been stored")
	}
	r.inMemoryDatabase[user.id] = user
	return user, nil
}
