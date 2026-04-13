package domain

import (
	"time"

	"github.com/google/uuid"
)

func NewBase() Base {
	now := time.Now()
	return Base{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// future helpers go here too
func NewID() uuid.UUID {
	return uuid.New()
}

func NowUTC() time.Time {
	return time.Now().UTC()
}
