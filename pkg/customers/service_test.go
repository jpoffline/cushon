package customers

import (
	"cushon/pkg/dto"
	"testing"

	"github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

func TestCreateCustomer(t *testing.T) {
	mockStore := NewMockCustomerStore(t)
	svc := New(mockStore)

	customerInput := &dto.CustomerInput{
		Name: "Test Customer",
	}

	mockStore.On("CreateCustomer", mock.Anything, customerInput).Return(&dto.Customer{
		ID:   uuid.New(),
		Name: customerInput.Name,
	}, nil)

	customer, err := svc.CreateCustomer(nil, customerInput)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if customer == nil || customer.Name != customerInput.Name {
		t.Errorf("expected customer with name %s, got %v", customerInput.Name, customer)
	}

	mockStore.AssertExpectations(t)
}

func TestGetCustomers(t *testing.T) {
	mockStore := NewMockCustomerStore(t)
	svc := New(mockStore)

	expectedCustomers := []*dto.Customer{
		{ID: uuid.New(), Name: "Customer 1"},
		{ID: uuid.New(), Name: "Customer 2"},
	}

	mockStore.On("GetCustomers", mock.Anything).Return(expectedCustomers, nil)

	customers, err := svc.GetCustomers(nil)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(customers) != len(expectedCustomers) {
		t.Errorf("expected %d customers, got %d", len(expectedCustomers), len(customers))
	}

	mockStore.AssertExpectations(t)
}
func TestGetDepositsByCustomerId(t *testing.T) {
	mockStore := NewMockCustomerStore(t)
	svc := New(mockStore)

	customerId := uuid.New()
	expectedDeposits := []*dto.Deposit{
		{ID: uuid.New(), Amount: 100.0, CustomerId: customerId},
		{ID: uuid.New(), Amount: 200.0, CustomerId: customerId},
	}

	mockStore.On("GetDepositsByCustomerId", mock.Anything, customerId).Return(expectedDeposits, nil)

	deposits, err := svc.GetDepositsByCustomerId(nil, customerId)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(deposits) != len(expectedDeposits) {
		t.Errorf("expected %d deposits, got %d", len(expectedDeposits), len(deposits))
	}

	mockStore.AssertExpectations(t)
}
