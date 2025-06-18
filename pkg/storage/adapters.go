package storage

import (
	"cushon/ent"
	"cushon/pkg/dto"
)

func DepositFromEnt(deposit *ent.Deposit) *dto.Deposit {
	return &dto.Deposit{
		ID:        deposit.ID,
		Amount:    deposit.Amount,
		CreatedAt: deposit.CreatedAt,
	}
}

func CustomerFromEnt(customer *ent.Customer) *dto.Customer {
	return &dto.Customer{
		ID:   customer.ID,
		Name: customer.Name,
	}
}

func FundFromEnt(fund *ent.Fund) *dto.Fund {
	return &dto.Fund{
		ID:   fund.ID,
		Name: fund.Name,
	}
}
