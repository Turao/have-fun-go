package auctions

import (
	"testing"

	"github.com/google/uuid"
)

func TestAuctionStart(t *testing.T) {
	auction := New()
	startedAuction, err := auction.Start()
	if err != nil {
		t.Error(err.Error())
	}

	if startedAuction.startTime == nil {
		t.Error("Auction has no start timestamp. Should have")
	}
}

func TestAuctionStartMoreThanOnce(t *testing.T) {
	auction := New()
	auction.Start()
	startedAuction, err := auction.Start()
	if err == nil {
		t.Error("No error when starting an auction more than once.")
	}

	if startedAuction.startTime == nil {
		t.Error("Auction has no start timestamp. Should have")
	}
}

func TestAuctionEnd(t *testing.T) {
	auction := New()
	auction.Start()

	endedAuction, err := auction.End()
	if err != nil {
		t.Error(err.Error())
	}

	if endedAuction.endTime == nil {
		t.Error("Auction has no end timestamp. Should have")
	}
}

func TestAuctionEndButNotStartedYet(t *testing.T) {
	auction := New()

	endedAuction, err := auction.End()
	if err == nil {
		t.Error("No error when ending an auction more than once.")
	}

	if endedAuction.endTime != nil {
		t.Error("Auction has no end timestamp. Should have")
	}
}

func TestAuctionEndMoreThanOnce(t *testing.T) {
	auction := New()
	auction.Start()
	auction.End()

	endedAuction, err := auction.End()
	if err == nil {
		t.Error("No error when ending an auction more than once.")
	}

	if endedAuction.endTime == nil {
		t.Error("Auction has no end timestamp. Should have")
	}
}

func TestBidPlaced(t *testing.T) {
	auction := New()
	auction.Start()
	_, err := auction.PlaceBid(uuid.New(), uuid.New(), uuid.New(), 10)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestBidPlacedButAuctionHasNotStarted(t *testing.T) {
	auction := New()
	_, err := auction.PlaceBid(uuid.New(), uuid.New(), uuid.New(), 10)
	if err == nil {
		t.Error("No errors while placing bid at auction that has not started")
	}

}

func TestBidPlacedButAuctionHasEnded(t *testing.T) {
	auction := New()
	auction.Start()
	auction.End()

	_, err := auction.PlaceBid(uuid.New(), uuid.New(), uuid.New(), 10)
	if err == nil {
		t.Error("No errors while placing bid at auction that has not started")
	}
}
