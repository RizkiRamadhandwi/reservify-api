package controller

import (
	"booking-room-app/config"
	"booking-room-app/delivery/middleware"
	"booking-room-app/entity"
	"booking-room-app/shared/common"
	"booking-room-app/usecase"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionsController struct {
	transactionUC usecase.TransactionsUsecase
	rg *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (t *TransactionsController) createHandler(ctx *gin.Context) {
	var payload entity.Transaction
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	transactions, err := t.transactionUC.RequestNewBookingRooms(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, transactions, "Created")
}

func (t *TransactionsController) listHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))

	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")

	if page == 0 || size == 0 {
		page = 1
		size = 5
	}
	
	if startDate == "" {
		startDate = "1000-01-01" 
	}

	if endDate == "" {
		endDate = "3000-12-31" 
	}

	startDateTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid startDate format")
		return
	}

	endDateTime, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid endDate format")
		return
	}

	transactions, paging, err := t.transactionUC.FindAllTransactions(page, size, startDateTime, endDateTime)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var response []interface{}
	for _, v := range transactions {
		response = append(response, v)
	}
	common.SendPagedResponse(ctx, response, paging, "Ok")
}

func (t *TransactionsController) getTransactionById(ctx *gin.Context) {
	id := ctx.Param("id")
	transactions, err := t.transactionUC.FindTransactionsById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "transaction with transaction ID "+id+" not found")
		return
	}

	common.SendSingleResponse(ctx, transactions, "Ok")
}

func (t *TransactionsController) getTransactionByEmployeeId(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	employeeId := ctx.Param("employeeId")

	if page == 0 || size == 0 {
		page = 1
		size = 5
	}

	transactions, paging, err := t.transactionUC.FindTransactionsByEmployeeId(employeeId, page, size)
	if err != nil {
		fmt.Println(employeeId)
		common.SendErrorResponse(ctx, http.StatusNotFound, "transaction with employee ID "+employeeId+" not found")
		return
	}

	var response []interface{}
	for _, v := range transactions {
		response = append(response, v)
	}

	common.SendPagedResponse(ctx, response, paging, "Ok")
}

func (t *TransactionsController) updateStatusHandler(ctx *gin.Context) {
	var payload entity.Transaction
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	transactions, err := t.transactionUC.AccStatusBooking(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, transactions, "Updated")
}

func (t *TransactionsController) Route() {
	t.rg.GET(config.TransactionList, t.authMiddleware.RequireToken("admin", "ga"), t.listHandler)
	t.rg.GET(config.TransactionGetById, t.authMiddleware.RequireToken("admin", "ga", "employee"), t.getTransactionById)
	t.rg.GET(config.TransactionGetByEmpId, t.authMiddleware.RequireToken("admin", "employee"), t.getTransactionByEmployeeId)
	t.rg.POST(config.TransactionCreate, t.authMiddleware.RequireToken("admin", "employee"), t.createHandler)
	t.rg.PUT(config.TransactionUpdatePerm, t.authMiddleware.RequireToken("admin", "ga"), t.updateStatusHandler)
}

func NewTransactionsController(transactionUC usecase.TransactionsUsecase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware,) *TransactionsController {
	return &TransactionsController{
		transactionUC: transactionUC,
		rg:         rg,
		authMiddleware: authMiddleware,
	}
}