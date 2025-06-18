package handlers

import (
	"net/http"

	"cushon/pkg/dto"

	"github.com/gin-gonic/gin"
)

type FundService interface {
	CreateFund(ctx *gin.Context, fund *dto.FundInput) (*dto.Fund, error)
	GetFunds(ctx *gin.Context) ([]*dto.Fund, error)
}

func RegisterFundRoutes(r *gin.Engine, svc FundService) {
	fundsGroup := r.Group("/fund")
	{
		fundsGroup.GET("/list", getFunds(svc))
		fundsGroup.POST("", createFund(svc))
	}
}

// GetFunds godoc
// @Summary Get the list of funds
// @Schemes
// @Description Get the list of funds
// @Tags Funds
// @Accept json
// @Produce json
// @Success 200 {object} []dto.Fund
// @Router /fund/list [get]
func getFunds(svc FundService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		funds, err := svc.GetFunds(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, funds)
	}
}

// PostFunds godoc
// @Summary Add a fund
// @Schemes
// @Description Add a fund
// @Tags Funds
// @Param        fund  body      dto.FundInput  true  "Fund data"
// @Success 200 {object} dto.Fund
// @Router /fund [post]
func createFund(svc FundService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := dto.FundInput{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fund, err := svc.CreateFund(ctx, &input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, fund)
	}
}
