package user

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

type Repository interface {
	GetUser(uuid.UUID) (*user, error)
	GetUsers() ([]user, error)
	CreateUser(user *user) (*user, error)
	UpdateUser(user *user) (*user, error)
}

type InMemoryRepository struct {
	inMemoryDatabase map[uuid.UUID]*user
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{make(map[uuid.UUID]*user)}
}

func (r InMemoryRepository) GetUser(userId uuid.UUID) (*user, error) {
	log.Println("[in-memory repository]", "Getting User...")

	user, found := r.inMemoryDatabase[userId]
	if !found {
		return nil, errors.New("user does not exist")
	}
	return user, nil
}

func (r InMemoryRepository) GetUsers() ([]user, error) {
	log.Println("[in-memory repository]", "Getting Users...")

	users := make([]user, len(r.inMemoryDatabase))

	for _, u := range r.inMemoryDatabase {
		users = append(users, *u)
	}

	return users, nil
}

func (r *InMemoryRepository) CreateUser(user *user) (*user, error) {
	log.Println("[in-memory repository]", "Creating user...")

	_, found := r.inMemoryDatabase[user.id]
	if found {
		return nil, errors.New("user has already been stored")
	}
	r.inMemoryDatabase[user.id] = user
	return user, nil
}

func (r *InMemoryRepository) UpdateUser(user *user) (*user, error) {
	log.Println("[in-memory repository]", "Updating user...")

	_, found := r.inMemoryDatabase[user.id]
	if !found {
		return nil, errors.New("user does not exist")
	}
	r.inMemoryDatabase[user.id] = user
	return user, nil
}
