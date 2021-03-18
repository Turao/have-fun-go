package event

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	id        uuid.UUID
	timestamp time.Time
}

func New() Event {
	return Event{uuid.New(), time.Now()}
}
