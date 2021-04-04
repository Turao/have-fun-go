package card

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCard(t *testing.T) {
	repo := NewInMemoryRepository()
	card := New()

	stored, err := repo.CreateCard(card)
	assert.Nil(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, card, stored)
}

func TestCreateCardAlreadyExists(t *testing.T) {
	repo := NewInMemoryRepository()
	card := New()

	repo.CreateCard(card)
	stored, err := repo.CreateCard(card)

	assert.NotNil(t, err)
	assert.Nil(t, stored)
}

func TestGetCard(t *testing.T) {
	repo := NewInMemoryRepository()
	card := New()
	repo.CreateCard(card)

	found, err := repo.GetCard(card.id)
	assert.Nil(t, err)
	assert.Equal(t, card, found)
}

func TestGetCardDoesNotExist(t *testing.T) {
	repo := NewInMemoryRepository()
	card := New()

	found, err := repo.GetCard(card.id)
	assert.NotNil(t, err)
	assert.Nil(t, found)
}

func TestGetCards(t *testing.T) {
	repo := NewInMemoryRepository()
	card0 := New()
	repo.CreateCard(card0)

	card1 := New()
	repo.CreateCard(card1)

	cards, err := repo.GetCards()
	log.Println(cards)
	assert.Nil(t, err)

	assert.Contains(t, cards, card0)
	assert.Contains(t, cards, card1)
}

func TestUpdateCard(t *testing.T) {
	repository := NewInMemoryRepository()
	card := givenCardExists(t, repository)

	updated, err := repository.UpdateCard(card)

	assert.Nil(t, err)
	assert.NotNil(t, updated)
	// todo: can this be tested without handling the concrete implementation directly?
}

func TestUpdateCardDoesNotExist(t *testing.T) {
	repository := NewInMemoryRepository()
	card := New()

	updated, err := repository.UpdateCard(card)

	assert.NotNil(t, err)
	assert.Nil(t, updated)
}
