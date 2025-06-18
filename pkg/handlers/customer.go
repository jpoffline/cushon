package handlers

import (
	"cushon/pkg/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomerService interface {
	CreateCustomer(ctx *gin.Context, customer *dto.CustomerInput) (*dto.Customer, error)
	GetCustomers(ctx *gin.Context) ([]*dto.Customer, error)
	GetDepositsByCustomerId(ctx *gin.Context, customerId uuid.UUID) ([]*dto.Deposit, error)
}

func RegisterCustomerRoutes(r *gin.Engine, svc CustomerService) {
	customersGroup := r.Group("/customer")
	{
		customersGroup.GET("/list", getCustomers(svc))
		customersGroup.POST("", createCustomer(svc))
		customersGroup.GET("/:customer_id/deposits", getDepositsByCustomerId(svc))
	}
}

// GetCustomersDeposits godoc
// @Summary Get deposits for a specific customer
// @Schemes
// @Description deposits for a specific customer
// @Tags Customers
// @Param customer_id path string true "Customer ID"
// @Produce json
// @Success 200 {object} []dto.Deposit
// @Router /customer/{customer_id}/deposits [get]
func getDepositsByCustomerId(svc CustomerService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerId, err := uuid.Parse(ctx.Param("customer_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
			return
		}

		deposits, err := svc.GetDepositsByCustomerId(ctx, customerId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, deposits)
	}
}

// GetCustomers godoc
// @Summary Get the list of customers
// @Schemes
// @Description Get the list of customers
// @Tags Customers
// @Accept json
// @Produce json
// @Success 200 {object} []dto.Customer
// @Router /customer/list [get]
func getCustomers(svc CustomerService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customers, err := svc.GetCustomers(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, customers)
	}
}

// PostCustomer godoc
// @Summary Add a customer
// @Schemes
// @Tags Customers
// @Description Add a customer
// @Param        customer  body      dto.CustomerInput  true  "Customer data"
// @Success 200 {object} dto.Customer
// @Router /customer [post]
func createCustomer(svc CustomerService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := dto.CustomerInput{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fund, err := svc.CreateCustomer(ctx, &input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, fund)
	}
}
