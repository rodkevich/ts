package models

import (
	"time"

	"github.com/google/uuid"
)

// TicketTag ...
type TicketTag struct {
	TicketID  uuid.UUID
	TagID     uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
