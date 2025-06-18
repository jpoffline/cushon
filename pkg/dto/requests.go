package dto

import (
	"cushon/ent"

	"github.com/google/uuid"
)

type CustomerInput struct {
	Name string `json:"name"`
}

type FundInput struct {
	Name string `json:"name"`
}

type DepositInput struct {
	Amount     float64   `json:"amount"`
	CustomerId uuid.UUID `json:"customer_id"`
	FundId     uuid.UUID `json:"fund_id"`
}

func (f *FundInput) ToEnt() *ent.Fund {
	return &ent.Fund{
		Name: f.Name,
	}
}

func (c *CustomerInput) ToEnt() *ent.Customer {
	return &ent.Customer{
		Name: c.Name,
	}
}
