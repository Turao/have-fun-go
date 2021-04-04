package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddCard(t *testing.T) {
	user := New("dummy")
	cardID := uuid.New()

	err := user.AddCard(cardID)

	assert.Nil(t, err)
	assert.Contains(t, user.cards, cardID)
}

func TestAddCardMoreThanOnce(t *testing.T) {
	user := New("dummy")
	cardID := uuid.New()
	user.AddCard(cardID)

	err := user.AddCard(cardID)

	assert.NotNil(t, err)
}

func TestRemoveCard(t *testing.T) {
	user := New("dummy")
	cardID := uuid.New()
	user.AddCard(cardID)

	err := user.RemoveCard(cardID)

	assert.Nil(t, err)
	assert.Empty(t, user.cards)
}

func TestRemoveCardMoreThanOnce(t *testing.T) {
	user := New("dummy")
	cardID := uuid.New()
	user.AddCard(cardID)
	user.RemoveCard(cardID)

	err := user.RemoveCard(cardID)

	assert.NotNil(t, err)
}
