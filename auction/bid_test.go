package auction

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestClosePositionAsOpen(t *testing.T) {
	bid := newBid(uuid.New(), uuid.New(), 10)

	_, err := bid.ClosePosition(OPEN)

	assert.NotNil(t, err)
}

func TestClosePositionAsWon(t *testing.T) {
	bid := newBid(uuid.New(), uuid.New(), 10)

	_, err := bid.ClosePosition(WON)

	assert.Nil(t, err)
	assert.Equal(t, status(WON), bid.status)
}

func TestClosePositionAsLost(t *testing.T) {
	bid := newBid(uuid.New(), uuid.New(), 10)

	_, err := bid.ClosePosition(LOST)

	assert.Nil(t, err)
	assert.Equal(t, status(LOST), bid.status)
}

func TestClosePositionAsInvalidStatus(t *testing.T) {
	bid := newBid(uuid.New(), uuid.New(), 10)

	_, err := bid.ClosePosition("why-cant-go-provide-enums")

	assert.NotNil(t, err)
}
