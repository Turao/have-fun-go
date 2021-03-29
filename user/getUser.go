package user

import (
	"log"

	"github.com/google/uuid"
)

type GetUser interface {
	Execute(userId uuid.UUID) (*user, error)
}

type GetUserUseCase struct {
	Repository Repository
}

func (uc *GetUserUseCase) Execute(userId uuid.UUID) (*user, error) {
	log.Println("[getUser] Getting User...")

	found, err := uc.Repository.GetUser(userId)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}
	return found, nil
}
