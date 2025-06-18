package deposit

import (
	"cushon/pkg/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DepositStore interface {
	IsFund(ctx *gin.Context, fundId uuid.UUID) (bool, error)
	IsCustomer(ctx *gin.Context, customerId uuid.UUID) (bool, error)
	CreateDeposit(ctx *gin.Context, input dto.DepositInput) (*dto.Deposit, error)
}
type DepositService struct {
	store DepositStore
}

func New(store DepositStore) *DepositService {
	return &DepositService{store: store}
}

func (s *DepositService) isFund(ctx *gin.Context, fundId uuid.UUID) error {
	isFund, err := s.store.IsFund(ctx, fundId)
	if err != nil {
		return err
	}
	if !isFund {
		return ErrFundNotFound
	}
	return nil
}

func (s *DepositService) isCustomer(ctx *gin.Context, customerId uuid.UUID) error {
	isCustomer, err := s.store.IsCustomer(ctx, customerId)
	if err != nil {
		return err
	}
	if !isCustomer {
		return ErrCustomerNotFound
	}
	return nil
}

func (s *DepositService) Deposit(ctx *gin.Context, amount float64, fundId, customerId uuid.UUID) error {

	if err := s.isFund(ctx, fundId); err != nil {
		return err
	}

	if err := s.isCustomer(ctx, customerId); err != nil {
		return err
	}

	if _, err := s.store.CreateDeposit(ctx, dto.DepositInput{
		Amount:     amount,
		CustomerId: customerId,
		FundId:     fundId,
	}); err != nil {
		return err
	}

	return nil
}
