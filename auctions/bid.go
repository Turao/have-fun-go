package auctions

import (
	"errors"

	"github.com/google/uuid"
)

type Bid interface {
	Id() uuid.UUID
	BidderId() uuid.UUID
	ItemId() uuid.UUID
	Price() int
}

type bid struct {
	id       uuid.UUID
	bidderId uuid.UUID
	itemId   uuid.UUID
	price    int
	status   status
}

type status string

const (
	OPEN = "open"
	WON  = "won"
	LOST = "lost"
)

func newBid(bidderId uuid.UUID, itemId uuid.UUID, price int) *bid {
	return &bid{uuid.New(), bidderId, itemId, price, OPEN}
}

// WHAT THE HELL IS THIS THING?!
// WHY CAN'T GO PROVIDE BASIC TYPES LIKE ENUMS?!
func (bid *bid) ClosePosition(status status) (*bid, error) {
	switch status {
	case OPEN:
		return bid, errors.New("Bids must not be closed with status OPEN")

	case WON:
		bid.status = WON
		return bid, nil

	case LOST:
		bid.status = LOST
		return bid, nil

	default:
		return bid, errors.New("Unable to close bid position. Invalid Status received")
	}
}
