package card

import "log"

type CreateCard interface {
	Execute() (Card, error)
}

type CreateCardUseCase struct {
	Repository Repository
}

func (uc CreateCardUseCase) Execute() (Card, error) {
	log.Println("[create card use-case]", "creating card...")
	card := New()

	created, err := uc.Repository.CreateCard(card)
	if err != nil {
		return nil, err // todo: do not return naked errors!
	}
	return created, nil
}
