package user

import (
	"log"

	"github.com/google/uuid"
)

type AddCard interface {
	Execute(userID uuid.UUID, cardID uuid.UUID) (User, error)
}

type AddCardUseCase struct {
	Repository Repository
}

func (uc AddCardUseCase) Execute(userID uuid.UUID, cardID uuid.UUID) (User, error) {
	log.Println("[add card use-case]", "adding card to user...", cardID, userID)

	user, err := uc.Repository.GetUser(userID)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	err = user.AddCard(cardID)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	user, err = uc.Repository.UpdateUser(user)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	log.Println("[add card use-case]", "card added")
	return user, nil
}
