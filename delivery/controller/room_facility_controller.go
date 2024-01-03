package controller

import (
	"booking-room-app/config"
	"booking-room-app/delivery/middleware"
	"booking-room-app/entity"
	"booking-room-app/shared/common"
	"booking-room-app/usecase"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type RoomFacilityController struct {
	transactionUC  usecase.RoomFacilityUsecase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (t *RoomFacilityController) createRoomFacilityHandler(ctx *gin.Context) {
	var payload entity.RoomFacility
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// check required field
	if payload.RoomId == "" || payload.FacilityId == "" || payload.Quantity == 0 {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "field roomId, facilityId, and quantity are required")
		return
	}

	// check valid uuid
	_, err = uuid.Parse(payload.RoomId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid uuid for roomId field")
		return
	}
	_, err = uuid.Parse(payload.FacilityId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid uuid for facilityId field")
		return
	}

	transactions, statusCode, err := t.transactionUC.AddRoomFacilityTransaction(payload)
	if err != nil {
		common.SendErrorResponse(ctx, statusCode, err.Error())
		return
	}
	common.SendCreateResponse(ctx, transactions, "Created")
}

func (t *RoomFacilityController) listRoomFacilityHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))

	transactions, paging, err := t.transactionUC.FindAllRoomFacility(page, size)
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

func (t *RoomFacilityController) getRoomFacilityById(ctx *gin.Context) {
	id := ctx.Param("id")
	transactions, _, err := t.transactionUC.FindRoomFacilityById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "transaction with transaction ID "+id+" not found")
		return
	}

	common.SendSingleResponse(ctx, transactions, "Ok")
}

func (t *RoomFacilityController) updateRoomFacilityHandler(ctx *gin.Context) {
	var payload entity.RoomFacility
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if payload.FacilityId == "" && payload.RoomId == "" && payload.Quantity == 0 {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "required at least one of folowing field to update: roomId, facilityId, quantity")
		return
	}

	transactions, statusCode, err := t.transactionUC.UpdateRoomFacilityTransaction(payload)
	if err != nil {
		common.SendErrorResponse(ctx, statusCode, err.Error())
		return
	}
	common.SendCreateResponse(ctx, transactions, "Updated")
}

func (t *RoomFacilityController) Route() {
	t.rg.GET(config.RoomFacilityList, t.authMiddleware.RequireToken("admin"), t.listRoomFacilityHandler)
	t.rg.GET(config.RoomFacilityGetById, t.authMiddleware.RequireToken("admin"), t.getRoomFacilityById)
	t.rg.POST(config.RoomFacilityCreate, t.authMiddleware.RequireToken("admin"), t.createRoomFacilityHandler)
	t.rg.PUT(config.RoomFacilityUpdate, t.authMiddleware.RequireToken("admin"), t.updateRoomFacilityHandler)
}

func NewRoomFacilityController(transactionUC usecase.RoomFacilityUsecase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *RoomFacilityController {
	return &RoomFacilityController{
		transactionUC:  transactionUC,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}

