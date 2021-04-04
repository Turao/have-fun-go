package card

import (
	"errors"
	"log"
	"sync"

	"github.com/google/uuid"
)

type Repository interface {
	GetCard(uuid.UUID) (Card, error)
	GetCards() ([]Card, error)
	CreateCard(card Card) (Card, error)
	UpdateCard(card Card) (Card, error)
}

type InMemoryRepository struct {
	mutex sync.Mutex
	cards map[uuid.UUID]Card
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{mutex: sync.Mutex{}, cards: make(map[uuid.UUID]Card)}
}

func (r *InMemoryRepository) GetCard(cardId uuid.UUID) (Card, error) {
	log.Println("[in-memory repository]", "getting card...")
	r.mutex.Lock()
	defer r.mutex.Unlock()

	card, found := r.cards[cardId]
	if !found {
		return nil, errors.New("card does not exist")
	}
	return card, nil
}

func (r *InMemoryRepository) GetCards() ([]Card, error) {
	log.Println("[in-memory repository]", "getting cards...")
	r.mutex.Lock()
	defer r.mutex.Unlock()

	cards := make([]Card, 0, len(r.cards))

	for _, u := range r.cards {
		cards = append(cards, u)
	}

	return cards, nil
}

func (r *InMemoryRepository) CreateCard(card Card) (Card, error) {
	log.Println("[in-memory repository]", "creating card...")
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, found := r.cards[card.Id()]
	if found {
		return nil, errors.New("card has already been stored")
	}
	r.cards[card.Id()] = card
	return card, nil
}

func (r *InMemoryRepository) UpdateCard(card Card) (Card, error) {
	log.Println("[in-memory repository]", "updating card...")
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, found := r.cards[card.Id()]
	if !found {
		return nil, errors.New("card does not exist")
	}
	r.cards[card.Id()] = card
	return card, nil
}
