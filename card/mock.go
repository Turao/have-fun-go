package card

import (
	"testing"

	"github.com/google/uuid"
)

func givenCardExists(t *testing.T, repository Repository) Card {
	card, err := repository.CreateCard(New())
	if err != nil {
		t.Fatal("failed on helper function")
	}
	return card
}

func givenCardHasOwner(t *testing.T, repository Repository, card Card) Card {
	err := card.AssignOwner(uuid.New())
	if err != nil {
		t.Fatal("failed on helper function")
	}
	card, err = repository.UpdateCard(card)
	if err != nil {
		t.Fatal("failed on helper function")
	}
	return card
}
