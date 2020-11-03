package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"implementasi-mvc/app/model"
	"implementasi-mvc/app/utils"
	"net/http"
)

type TransactionController struct {
	DB *gorm.DB
}

func (ctrl TransactionController) Transfer (ctx *gin.Context) {
	transactionModel := model.Transaction{
		DB: ctrl.DB,
	}

	err := ctx.Bind(&transactionModel)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Transfer()
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}

func (ctrl TransactionController) Withdraw (ctx *gin.Context) {
	transactionModel := model.Transaction{
		DB: ctrl.DB,
	}

	err := ctx.Bind(&transactionModel)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Withdraw()
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}

func (ctrl TransactionController) Deposit (ctx *gin.Context) {
	transactionModel := model.Transaction{
		DB: ctrl.DB,
	}

	err := ctx.Bind(&transactionModel)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Deposit()
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}
