package card

import (
	"log"

	"github.com/google/uuid"
)

type AssignOwner interface {
	Execute(name string) (*card, error)
}

type AssignOwnerUseCase struct {
	Repository Repository
}

func (uc AssignOwnerUseCase) Execute(ownerId uuid.UUID, cardId uuid.UUID) (*card, error) {
	log.Println("[AssignOwner]", "Assigning Owner to card...")
	card, err := uc.Repository.GetCard(cardId)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	card, err = card.AssignOwner(ownerId)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	card, err = uc.Repository.UpdateCard(card)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	return card, nil
}
