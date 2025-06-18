package deposit

import (
	"cushon/pkg/dto"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestDepositFundNotFound(t *testing.T) {
	mockStore := NewMockDepositStore(t)
	svc := New(mockStore)
	mockStore.On("IsFund", mock.Anything, mock.Anything).Return(false, nil)
	err := svc.Deposit(nil, 100.0, uuid.New(), uuid.New())
	if err == nil {
		t.Error("expected error for non-existent fund, got nil")
	}
	assert.Error(t, err, ErrFundNotFound)
	mockStore.AssertExpectations(t)
}

func TestDepositCustomerNotFound(t *testing.T) {
	mockStore := NewMockDepositStore(t)
	svc := New(mockStore)
	mockStore.On("IsFund", mock.Anything, mock.Anything).Return(true, nil)
	mockStore.On("IsCustomer", mock.Anything, mock.Anything).Return(false, nil)
	err := svc.Deposit(nil, 100.0, uuid.New(), uuid.New())
	if err == nil {
		t.Error("expected error for non-existent customer, got nil")
	}
	assert.Error(t, err, ErrCustomerNotFound)
	mockStore.AssertExpectations(t)
}

func TestDepositOk(t *testing.T) {
	mockStore := NewMockDepositStore(t)
	svc := New(mockStore)
	fundId := uuid.New()
	customerId := uuid.New()

	mockStore.On("IsFund", mock.Anything, fundId).Return(true, nil)
	mockStore.On("IsCustomer", mock.Anything, customerId).Return(true, nil)
	mockStore.On("CreateDeposit", mock.Anything, mock.AnythingOfType("dto.DepositInput")).Return(&dto.Deposit{}, nil)

	err := svc.Deposit(nil, 100.0, fundId, customerId)
	assert.NoError(t, err)
}
