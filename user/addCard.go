package user

import (
	"log"

	"github.com/google/uuid"
)

type AddCard interface {
	Execute(userId uuid.UUID, cardId uuid.UUID) (*user, error)
}

type AddCardUseCase struct {
	Repository Repository
}

func (uc AddCardUseCase) Execute(userId uuid.UUID, cardId uuid.UUID) (*user, error) {
	log.Println("[addCard]", "Adding card to user...", cardId, userId)

	user, err := uc.Repository.GetUser(userId)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	user, err = user.AddCard(cardId)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	user, err = uc.Repository.UpdateUser(user)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	log.Println("[addCard]", "Card added!")
	return user, nil
}
