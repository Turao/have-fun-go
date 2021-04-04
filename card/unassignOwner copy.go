package card

import (
	"log"

	"github.com/google/uuid"
)

type UnassignOwner interface {
	Execute(cardId uuid.UUID) (*card, error)
}

type UnassignOwnerUseCase struct {
	Repository Repository
}

func (uc UnassignOwnerUseCase) Execute(cardId uuid.UUID) (*card, error) {
	log.Println("[UnassignOwner]", "Unassigning Owner from card...")
	card, err := uc.Repository.GetCard(cardId)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	card, err = card.UnassignOwner()
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	card, err = uc.Repository.UpdateCard(card)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	return card, nil
}
