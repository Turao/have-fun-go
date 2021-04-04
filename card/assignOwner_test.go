package card

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func givenCardExists(t *testing.T, repository Repository) *card {
	card := New()
	card, err := repository.CreateCard(card)
	if err != nil {
		t.Fatal("failed on helper function")
	}
	return card
}

func TestAssignOwnerUseCase(t *testing.T) {
	repository := NewInMemoryRepository()
	card := givenCardExists(t, repository)

	assignOwner := AssignOwnerUseCase{Repository: repository}

	card, err := assignOwner.Execute(uuid.New(), card.Id())
	assert.Nil(t, err)
	assert.NotNil(t, card)
}
