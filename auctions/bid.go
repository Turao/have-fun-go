package auctions

import "github.com/google/uuid"

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
}
