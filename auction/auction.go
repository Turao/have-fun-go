package auction

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Auction interface {
	ID()
	StartTime() time.Time
	EndTime() time.Time
	Start() (Auction, error)
	End() (Auction, error)
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
		Bids      []bid      `json:"bids"`
	}{
		auction.id,
		auction.startTime,
		auction.endTime,
		auction.bids,
	})
}

func New() *auction {
	return &auction{uuid.New(), nil, nil, make([]bid, 0)}
}

func (auction auction) ID() uuid.UUID {
	return auction.id
}

func (auction *auction) Start() (*auction, error) {
	if auction.startTime != nil {
		return auction, errors.New("auction has already started")
	}

	now := time.Now()
	auction.startTime = &now
	return auction, nil
}

func (auction *auction) End() (*auction, error) {
	if auction.startTime == nil {
		return auction, errors.New("auction has not started yet")
	}

	if auction.endTime != nil {
		return auction, errors.New("auction has already ended")
	}

	now := time.Now()
	auction.endTime = &now
	return auction, nil
}

func (auction *auction) PlaceBid(bidderID uuid.UUID, itemID uuid.UUID, price uint) (*auction, error) {
	if auction.startTime == nil {
		return auction, errors.New("auction has not started yet")
	}

	if auction.endTime != nil {
		return auction, errors.New("auction has already ended")
	}

	newBid := newBid(bidderID, itemID, price)
	auction.bids = append(auction.bids, *newBid)
	return auction, nil
}
