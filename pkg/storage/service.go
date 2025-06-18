package storage

import (
	"cushon/ent"
	"cushon/ent/customer"
	"cushon/ent/deposit"
	"cushon/ent/fund"
	"cushon/pkg/config"
	"cushon/pkg/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct {
	client *ent.Client
}

func New(cfg *config.Config) *Service {
	return &Service{
		client: newClient(cfg),
	}
}

func (s *Service) Close() error {
	return s.client.Close()
}

func (s *Service) GetFunds(ctx *gin.Context) ([]*dto.Fund, error) {
	funds, err := s.client.Fund.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var result []*dto.Fund
	for _, fund := range funds {
		result = append(result, FundFromEnt(fund))
	}
	return result, nil
}

func (s *Service) CreateFund(ctx *gin.Context, fund *dto.FundInput) (*dto.Fund, error) {
	f, err := s.client.Fund.Create().
		SetName(fund.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return FundFromEnt(f), nil

}

func (s *Service) GetCustomers(ctx *gin.Context) ([]*dto.Customer, error) {
	customers, err := s.client.Customer.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var result []*dto.Customer
	for _, customer := range customers {
		result = append(result, CustomerFromEnt(customer))
	}
	return result, nil
}

func (s *Service) GetDepositsByCustomerId(ctx *gin.Context, customerId uuid.UUID) ([]*dto.Deposit, error) {
	deposits, err := s.client.Deposit.
		Query().
		Where(deposit.HasCustomerWith(customer.IDEQ(customerId))).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var result []*dto.Deposit
	for _, deposit := range deposits {
		result = append(result, DepositFromEnt(deposit))
	}
	return result, nil
}

func (s *Service) CreateCustomer(ctx *gin.Context, customer *dto.CustomerInput) (*dto.Customer, error) {
	cust, err := s.client.Customer.Create().
		SetName(customer.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return CustomerFromEnt(cust), nil
}

func (s *Service) IsCustomer(ctx *gin.Context, customerId uuid.UUID) (bool, error) {
	count, err := s.client.Customer.Query().
		Where(customer.IDEQ(customerId)).
		Count(ctx)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *Service) IsFund(ctx *gin.Context, fundId uuid.UUID) (bool, error) {
	count, err := s.client.Fund.Query().
		Where(fund.IDEQ(fundId)).
		Count(ctx)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *Service) CreateDeposit(ctx *gin.Context, input dto.DepositInput) (*dto.Deposit, error) {
	deposit, err := s.client.Deposit.Create().
		SetAmount(input.Amount).
		AddCustomerIDs(input.CustomerId).
		AddFundIDs(input.FundId).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return DepositFromEnt(deposit), nil
}
