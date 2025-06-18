package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"cushon/pkg/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter(mockService DepositService) *gin.Engine {
	r := gin.Default()
	RegisterDepositRoutes(r, mockService)
	return r
}

func TestCreateDeposit_Success(t *testing.T) {
	mockService := NewMockDepositService(t)

	customerId := uuid.New()
	fundId := uuid.New()
	input := dto.DepositInput{
		Amount:     100.0,
		CustomerId: customerId,
		FundId:     fundId,
	}

	mockService.On("Deposit", mock.Anything, input.Amount, input.FundId, input.CustomerId).Return(nil)

	router := setupRouter(mockService)
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/deposit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	mockService.AssertExpectations(t)
}

func TestCreateDeposit_Failure(t *testing.T) {
	mockService := NewMockDepositService(t)

	customerId := uuid.New()
	fundId := uuid.New()
	input := dto.DepositInput{
		Amount:     100.0,
		CustomerId: customerId,
		FundId:     fundId,
	}

	err := errors.New("something")
	mockService.On("Deposit", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(err)

	router := setupRouter(mockService)
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/deposit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	mockService.AssertExpectations(t)
}
