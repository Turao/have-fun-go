package auction

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Bid interface {
	ID() uuid.UUID
	BidderId() uuid.UUID
	ItemId() uuid.UUID
	Price() uint
}

type bid struct {
	id       uuid.UUID
	bidderId uuid.UUID
	itemId   uuid.UUID
	price    uint
	status   status
}

type status string

const (
	OPEN = "open"
	WON  = "won"
	LOST = "lost"
)

func (bid *bid) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id       uuid.UUID `json:"id"`
		BidderId uuid.UUID `json:"bidderId"`
		ItemId   uuid.UUID `json:"itemId"`
		Price    uint      `json:"price"`
	}{
		bid.id,
		bid.bidderId,
		bid.itemId,
		bid.price,
	})
}

func newBid(bidderId uuid.UUID, itemId uuid.UUID, price uint) *bid {
	return &bid{uuid.New(), bidderId, itemId, price, OPEN}
}

// WHAT THE HELL IS THIS THING?!
// WHY CAN'T GO PROVIDE BASIC TYPES LIKE ENUMS?!
func (bid *bid) ClosePosition(status status) (*bid, error) {
	switch status {
	case OPEN:
		return bid, errors.New("bids must not be closed with status OPEN")

	case WON:
		bid.status = WON
		return bid, nil

	case LOST:
		bid.status = LOST
		return bid, nil

	default:
		return bid, errors.New("unable to close bid position. Invalid Status received")
	}
}
