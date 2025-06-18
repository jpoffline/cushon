package funds

import (
	"cushon/pkg/dto"

	"github.com/gin-gonic/gin"
)

type FundStore interface {
	CreateFund(ctx *gin.Context, fund *dto.FundInput) (*dto.Fund, error)
	GetFunds(ctx *gin.Context) ([]*dto.Fund, error)
}

type FundService struct {
	store FundStore
}

func New(store FundStore) *FundService {
	return &FundService{store: store}
}

func (s *FundService) CreateFund(ctx *gin.Context, fund *dto.FundInput) (*dto.Fund, error) {
	return s.store.CreateFund(ctx, fund)
}

func (s *FundService) GetFunds(ctx *gin.Context) ([]*dto.Fund, error) {
	return s.store.GetFunds(ctx)
}
