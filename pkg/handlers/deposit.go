package handlers

import (
	"cushon/pkg/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DepositService interface {
	Deposit(ctx *gin.Context, amount float64, fundId, customerId uuid.UUID) error
}

func RegisterDepositRoutes(r *gin.Engine, svc DepositService) {
	customersGroup := r.Group("/deposit")
	{
		customersGroup.POST("", createDeposit(svc))
	}
}

// PostDeposit godoc
// @Summary Add a deposit
// @Schemes
// @Tags Deposits
// @Description Add a deposit
// @Param        deposit  body      dto.DepositInput  true  "Deposit data"
// @Success 200 {object} dto.Deposit
// @Router /deposit [post]
func createDeposit(svc DepositService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := dto.DepositInput{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := svc.Deposit(ctx, input.Amount, input.FundId, input.CustomerId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	}
}
