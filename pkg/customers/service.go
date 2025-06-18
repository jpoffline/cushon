package customers

import (
	"cushon/pkg/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomerStore interface {
	CreateCustomer(ctx *gin.Context, customer *dto.CustomerInput) (*dto.Customer, error)
	GetCustomers(ctx *gin.Context) ([]*dto.Customer, error)
	GetDepositsByCustomerId(ctx *gin.Context, customerId uuid.UUID) ([]*dto.Deposit, error)
}

type CustomerService struct {
	store CustomerStore
}

func New(store CustomerStore) *CustomerService {
	return &CustomerService{store: store}
}
func (s *CustomerService) CreateCustomer(ctx *gin.Context, customer *dto.CustomerInput) (*dto.Customer, error) {
	return s.store.CreateCustomer(ctx, customer)
}
func (s *CustomerService) GetCustomers(ctx *gin.Context) ([]*dto.Customer, error) {
	return s.store.GetCustomers(ctx)
}
func (s *CustomerService) GetDepositsByCustomerId(ctx *gin.Context, customerId uuid.UUID) ([]*dto.Deposit, error) {
	return s.store.GetDepositsByCustomerId(ctx, customerId)
}
