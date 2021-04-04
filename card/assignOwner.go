package card

import (
	"log"

	"github.com/google/uuid"
)

type AssignOwner interface {
	Execute(name string) (Card, error)
}

type AssignOwnerUseCase struct {
	Repository Repository
}

func (uc AssignOwnerUseCase) Execute(cardId uuid.UUID, ownerId uuid.UUID) (Card, error) {
	log.Println("[assign owner use-case]", "assigning owner to card...")
	card, err := uc.Repository.GetCard(cardId)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	err = card.AssignOwner(ownerId)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	card, err = uc.Repository.UpdateCard(card)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	return card, nil
}
