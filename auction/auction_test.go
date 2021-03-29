package auction

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuctionStart(t *testing.T) {
	auction := New()

	startedAuction, err := auction.Start()

	assert.Nil(t, err)
	assert.NotNil(t, startedAuction.startTime)
}

func TestAuctionStartMoreThanOnce(t *testing.T) {
	auction := New()
	auction.Start()

	startedAuction, err := auction.Start()

	assert.NotNil(t, err)
	assert.NotNil(t, startedAuction.startTime)
}

func TestAuctionEnd(t *testing.T) {
	auction := New()
	auction.Start()

	endedAuction, err := auction.End()

	assert.Nil(t, err)
	assert.NotNil(t, endedAuction.endTime)
}

func TestAuctionEndButNotStartedYet(t *testing.T) {
	auction := New()

	endedAuction, err := auction.End()

	assert.NotNil(t, err)
	assert.Nil(t, endedAuction.endTime)
}

func TestAuctionEndMoreThanOnce(t *testing.T) {
	auction := New()
	auction.Start()
	auction.End()

	endedAuction, err := auction.End()

	assert.NotNil(t, err)
	assert.NotNil(t, endedAuction.endTime)
}

func TestBidPlaced(t *testing.T) {
	auction := New()
	auction.Start()

	_, err := auction.PlaceBid(uuid.New(), uuid.New(), 10)

	assert.Nil(t, err)
}

func TestBidPlacedButAuctionHasNotStarted(t *testing.T) {
	auction := New()

	_, err := auction.PlaceBid(uuid.New(), uuid.New(), 10)

	assert.NotNil(t, err)

}

func TestBidPlacedButAuctionHasEnded(t *testing.T) {
	auction := New()
	auction.Start()
	auction.End()

	_, err := auction.PlaceBid(uuid.New(), uuid.New(), 10)

	assert.NotNil(t, err)
}
