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
		t.Errorf(err.Error())
	}

	found := newUser.cards[cardId]
	if found == false {
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

func TestRemoveCard(t *testing.T) {
	user := New("dummy")
	cardId := uuid.New()
	user.AddCard(cardId)

	newUser, err := user.RemoveCard(cardId)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(newUser.cards) > 0 {
		t.Errorf("User still has cards. Should have none")
	}
}

func TestRemoveCardMoreThanOnce(t *testing.T) {
	user := New("dummy")
	cardId := uuid.New()
	user.AddCard(cardId)

	user.RemoveCard(cardId)
	_, err := user.RemoveCard(cardId)
	if err == nil {
		t.Errorf("No errors after trying to remove card")
	}
}
