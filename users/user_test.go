package users

import (
	"testing"

	"github.com/google/uuid"
)

func TestAddCard(t *testing.T) {
	user := New("dummy")
	cardId := uuid.New()
	newUser, err := user.AddCard(cardId)
	if err != nil {
		t.Errorf("Error while trying to add card to user")
	}

	found := false
	for _, card := range newUser.cards {
		if card == cardId {
			found = true
		}
	}

	if !found {
		t.Error("Card was not found after being added. It should have been")
	}
}

func TestAddCardMoreThanOnce(t *testing.T) {
	user := New("dummy")
	cardId := uuid.New()
	user.AddCard(cardId)

	_, err := user.AddCard(cardId)
	if err == nil {
		t.Errorf("No errors after trying to add already existing card")
	}
}
