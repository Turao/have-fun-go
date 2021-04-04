package user

import (
	"errors"
	"log"
	"sync"

	"github.com/google/uuid"
)

type Repository interface {
	GetUser(uuid.UUID) (User, error)
	GetUsers() ([]User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
}

type InMemoryRepository struct {
	mutex sync.Mutex
	users map[uuid.UUID]User
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{mutex: sync.Mutex{}, users: make(map[uuid.UUID]User)}
}

func (r *InMemoryRepository) GetUser(userId uuid.UUID) (User, error) {
	log.Println("[in-memory repository]", "Getting User...")
	r.mutex.Lock()
	defer r.mutex.Unlock()

	user, found := r.users[userId]
	if !found {
		return nil, errors.New("user does not exist")
	}
	return user, nil
}

func (r *InMemoryRepository) GetUsers() ([]User, error) {
	log.Println("[in-memory repository]", "Getting Users...")
	r.mutex.Lock()
	defer r.mutex.Unlock()

	users := make([]User, 0, len(r.users))

	for _, u := range r.users {
		users = append(users, u)
	}

	return users, nil
}

func (r *InMemoryRepository) CreateUser(user User) (User, error) {
	log.Println("[in-memory repository]", "Creating user...")
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, found := r.users[user.Id()]
	if found {
		return nil, errors.New("user has already been stored")
	}
	r.users[user.Id()] = user
	return user, nil
}

func (r *InMemoryRepository) UpdateUser(user User) (User, error) {
	log.Println("[in-memory repository]", "Updating user...")
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, found := r.users[user.Id()]
	if !found {
		return nil, errors.New("user does not exist")
	}
	r.users[user.Id()] = user
	return user, nil
}
