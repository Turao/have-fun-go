package card

import (
	"errors"
	"log"
	"sync"

	"github.com/google/uuid"
)

type Repository interface {
	GetCard(uuid.UUID) (*card, error)
	GetCards() ([]*card, error)
	CreateCard(card *card) (*card, error)
	UpdateCard(card *card) (*card, error)
}

type InMemoryRepository struct {
	mutex sync.Mutex
	cards map[uuid.UUID]*card
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{mutex: sync.Mutex{}, cards: make(map[uuid.UUID]*card)}
}

func (r *InMemoryRepository) GetCard(cardId uuid.UUID) (*card, error) {
	log.Println("[in-memory repository]", "Getting Card...")
	r.mutex.Lock()
	defer r.mutex.Unlock()

	card, found := r.cards[cardId]
	if !found {
		return nil, errors.New("card does not exist")
	}
	return card, nil
}

func (r *InMemoryRepository) GetCards() ([]*card, error) {
	log.Println("[in-memory repository]", "Getting Cards...")
	r.mutex.Lock()
	defer r.mutex.Unlock()

	cards := make([]*card, 0, len(r.cards))

	for _, u := range r.cards {
		cards = append(cards, u)
	}

	return cards, nil
}

func (r *InMemoryRepository) CreateCard(card *card) (*card, error) {
	log.Println("[in-memory repository]", "Creating card...")
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, found := r.cards[card.id]
	if found {
		return nil, errors.New("card has already been stored")
	}
	r.cards[card.id] = card
	return card, nil
}

func (r *InMemoryRepository) UpdateCard(card *card) (*card, error) {
	log.Println("[in-memory repository]", "Updating card...")
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, found := r.cards[card.id]
	if !found {
		return nil, errors.New("card does not exist")
	}
	r.cards[card.id] = card
	return card, nil
}
