package auctions

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Auction interface {
	Id()
	StartTime() time.Time
	EndTime() time.Time
	Start() (*Auction, error)
	End() (*Auction, error)
	Bids() []Bid
}

type auction struct {
	id        uuid.UUID
	startTime *time.Time
	endTime   *time.Time
	bids      []bid
}

func (auction auction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id        uuid.UUID  `json:"id"`
		StartTime *time.Time `json:"startTime"`
		EndTime   *time.Time `json:"endTime"`
	}{
		auction.id,
		auction.startTime,
		auction.endTime,
	})
}

func New() *auction {
	return &auction{uuid.New(), nil, nil, make([]bid, 0)}
}

func (auction auction) Id() uuid.UUID {
	return auction.id
}

func (auction *auction) Start() (*auction, error) {
	if auction.startTime != nil {
		return auction, errors.New("Auction has already started")
	}

	now := time.Now()
	auction.startTime = &now
	return auction, nil
}

func (auction *auction) End() (*auction, error) {
	if auction.startTime == nil {
		return auction, errors.New("Auction has not started yet")
	}

	if auction.endTime != nil {
		return auction, errors.New("Auction has already ended")
	}

	now := time.Now()
	auction.endTime = &now
	return auction, nil
}

func (auction *auction) PlaceBid(id uuid.UUID, bidderId uuid.UUID, itemId uuid.UUID, price int) (*auction, error) {
	if auction.startTime == nil {
		return auction, errors.New("Auction has not started yet")
	}

	if auction.endTime != nil {
		return auction, errors.New("Auction has already ended")
	}

	newBid := newBid(id, bidderId, itemId, price)
	auction.bids = append(auction.bids, *newBid)
	return auction, nil
}
