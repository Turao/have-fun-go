package auctions

import (
	"testing"

	"github.com/google/uuid"
)

func TestClosePositionAsOpen(t *testing.T) {
	bid := newBid(uuid.New(), uuid.New(), 10)
	_, err := bid.ClosePosition(OPEN)
	if err == nil {
		t.Error(err.Error())
	}
}

func TestClosePositionAsWon(t *testing.T) {
	bid := newBid(uuid.New(), uuid.New(), 10)
	_, err := bid.ClosePosition(WON)
	if err != nil {
		t.Error(err.Error())
	}

	if bid.status != WON {
		t.Error("Closed Bid should have status WON")
	}
}

func TestClosePositionAsLost(t *testing.T) {
	bid := newBid(uuid.New(), uuid.New(), 10)
	_, err := bid.ClosePosition(LOST)
	if err != nil {
		t.Error(err.Error())
	}

	if bid.status != LOST {
		t.Error("Closed Bid should have status LOST")
	}
}

func TestClosePositionAsInvalidStatus(t *testing.T) {
	bid := newBid(uuid.New(), uuid.New(), 10)
	_, err := bid.ClosePosition("why-cant-go-provide-enums")
	if err == nil {
		t.Error(err.Error())
	}
}
