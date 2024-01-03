package controller

import (
	"booking-room-app/config"
	"booking-room-app/delivery/middleware"
	"booking-room-app/entity"
	"booking-room-app/shared/common"
	"booking-room-app/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FacilitiesController struct {
	facilitiesUC   usecase.FacilitiesUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (f *FacilitiesController) updateHandler(ctx *gin.Context) {
	var payload entity.Facilities
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	facility, err := f.facilitiesUC.EditFacilities(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "not found ID "+payload.ID)
		return
	}
	common.SendCreateResponse(ctx, facility, "Updated")
}

func (f *FacilitiesController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	facility, err := f.facilitiesUC.FindFacilitiesById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "not found ID "+id)
		return
	}

	common.SendSingleResponse(ctx, facility, "ok")
}

func (f *FacilitiesController) listHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "5"))

	facilities, paging, err := f.facilitiesUC.FindAllFacilities(page, size)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var response []interface{}
	for _, f := range facilities {
		response = append(response, f)
	}

	common.SendPagedResponse(ctx, response, paging, "ok")

}

func (f *FacilitiesController) createHandler(ctx *gin.Context) {
	var payload entity.Facilities
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	facility, err := f.facilitiesUC.RegisterNewFacilities(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, facility, "Created")
}

func (f *FacilitiesController) Route() {
	f.rg.POST(config.FacilitiesCreate, f.authMiddleware.RequireToken("admin"), f.createHandler)
	f.rg.GET(config.FacilitiesList, f.authMiddleware.RequireToken("admin", "employee", "ga"), f.listHandler)
	f.rg.GET(config.FacilitiesGetById, f.authMiddleware.RequireToken("admin", "employee", "ga"), f.getHandler)
	f.rg.PUT(config.FacilitiesUpdate, f.authMiddleware.RequireToken("admin"), f.updateHandler)
}

func NewFacilitiesController(facilitiesUC usecase.FacilitiesUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *FacilitiesController {
	return &FacilitiesController{
		facilitiesUC:   facilitiesUC,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}
