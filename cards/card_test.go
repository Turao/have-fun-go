package cards

import (
	"testing"

	"github.com/google/uuid"
)

func TestAssignOwner(t *testing.T) {
	card := New()
	ownerId := uuid.New()
	_, err := card.AssignOwner(ownerId)
	if err != nil {
		t.Error(err.Error())
	}

	if card.ownerId != ownerId {
		t.Error("Card OwnerId does not match assigned OwnerId")
	}
}

func TestAssignOwnerFail(t *testing.T) {
	card := New()
	ownerId := uuid.New()
	card.AssignOwner(ownerId)
	newCard, err := card.AssignOwner(ownerId)
	if err == nil {
		t.Error()
	}
	if newCard.ownerId != ownerId {
		t.Error("Card OwnerId does not match assigned OwnerId")
	}
}

func TestUnassignOwner(t *testing.T) {
	card := New()
	ownerId := uuid.New()
	card.AssignOwner(ownerId)
	newCard, err := card.UnassignOwner()
	if err != nil {
		t.Error(err.Error())
	}

	if newCard.ownerId != uuid.Nil {
		t.Error("Card still has an Owner Id assigned")
	}
}

func TestUnassignOwnerFail(t *testing.T) {
	card := New()
	uuid.New()
	newCard, err := card.UnassignOwner()
	if err == nil {
		t.Error()
	}
	if newCard.ownerId != uuid.Nil {
		t.Error("Card still has an Owner Id assigned")
	}
}