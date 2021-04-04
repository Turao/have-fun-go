package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddCard(t *testing.T) {
	user := New("dummy")
	cardId := uuid.New()

	err := user.AddCard(cardId)

	assert.Nil(t, err)
	assert.Contains(t, user.cards, cardId)
}

func TestAddCardMoreThanOnce(t *testing.T) {
	user := New("dummy")
	cardId := uuid.New()
	user.AddCard(cardId)

	err := user.AddCard(cardId)

	assert.NotNil(t, err)
}

func TestRemoveCard(t *testing.T) {
	user := New("dummy")
	cardId := uuid.New()
	user.AddCard(cardId)

	err := user.RemoveCard(cardId)

	assert.Nil(t, err)
	assert.Empty(t, user.cards)
}

func TestRemoveCardMoreThanOnce(t *testing.T) {
	user := New("dummy")
	cardId := uuid.New()
	user.AddCard(cardId)
	user.RemoveCard(cardId)

	err := user.RemoveCard(cardId)

	assert.NotNil(t, err)
}
