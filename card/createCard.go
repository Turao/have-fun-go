package card

import "log"

type CreateCard interface {
	Execute() (*card, error)
}

type CreateCardUseCase struct {
	Repository Repository
}

func (uc CreateCardUseCase) Execute() (*card, error) {
	log.Println("[createCard]", "Creating Card...")
	card := New()

	created, err := uc.Repository.CreateCard(card)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}
	return created, nil
}
