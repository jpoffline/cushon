package dto

import (
	"time"

	"github.com/google/uuid"
)

type Fund struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Customer struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Deposit struct {
	ID         uuid.UUID `json:"id"`
	Amount     float64   `json:"amount"`
	CustomerId uuid.UUID `json:"customer_id"`
	FundId     uuid.UUID `json:"fund_id"`
	CreatedAt  time.Time `json:"created_at"`
}
