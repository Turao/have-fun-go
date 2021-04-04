package user

import (
	"log"

	"github.com/google/uuid"
)

type GetUser interface {
	Execute(userId uuid.UUID) (User, error)
}

type GetUserUseCase struct {
	Repository Repository
}

func (uc GetUserUseCase) Execute(userId uuid.UUID) (User, error) {
	log.Println("[get user use-case]", "getting user...")

	found, err := uc.Repository.GetUser(userId)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}
	return found, nil
}
