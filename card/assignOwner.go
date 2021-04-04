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

func (uc AssignOwnerUseCase) Execute(cardID uuid.UUID, ownerID uuid.UUID) (Card, error) {
	log.Println("[assign owner use-case]", "assigning owner to card...")
	card, err := uc.Repository.GetCard(cardID)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	err = card.AssignOwner(ownerID)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	card, err = uc.Repository.UpdateCard(card)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}

	return card, nil
}
