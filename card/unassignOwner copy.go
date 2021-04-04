package card

import (
	"log"

	"github.com/google/uuid"
)

type UnassignOwner interface {
	Execute(cardId uuid.UUID) (Card, error)
}

type UnassignOwnerUseCase struct {
	Repository Repository
}

func (uc UnassignOwnerUseCase) Execute(cardId uuid.UUID) (Card, error) {
	log.Println("[unassign owner use-case]", "unassigning owner from card...")
	card, err := uc.Repository.GetCard(cardId)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	err = card.UnassignOwner()
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	card, err = uc.Repository.UpdateCard(card)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	return card, nil
}
