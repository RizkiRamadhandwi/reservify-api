package controller

import (
	"booking-room-app/config"
	"booking-room-app/delivery/middleware"
	"booking-room-app/entity"
	"booking-room-app/shared/common"
	"booking-room-app/shared/model"
	"booking-room-app/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	roomUC         usecase.RoomUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (r *RoomController) createHandler(c *gin.Context) {
	var payload entity.Room
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	room, err := r.roomUC.RegisterNewRoom(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(c, room, "Created")
}

func (r *RoomController) getHandler(c *gin.Context) {
	id := c.Param("id")
	room, err := r.roomUC.FindRoomByID(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusNotFound, fmt.Sprintf("Room with ID %s not found", id))
		return
	}
	common.SendSingleResponse(c, room, "Ok")
}

func (r *RoomController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	status := c.Query("status")

	var rooms []entity.Room
	var paging model.Paging
	var err error

	if status == "" {
		if page == 0 && size == 0 {
			rooms, paging, err = r.roomUC.FindAllRoom(1, 5)
		} else {
			rooms, paging, err = r.roomUC.FindAllRoom(page, size)
		}
	} else {
		if page == 0 && size == 0 {
			rooms, paging, err = r.roomUC.FindAllRoomStatus(status, 1, 5)
		} else {
			rooms, paging, err = r.roomUC.FindAllRoomStatus(status, page, size)
		}
	}

	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var response []interface{}
	for _, v := range rooms {
		response = append(response, v)
	}
	common.SendPagedResponse(c, response, paging, "Ok")
}

func (r *RoomController) updateDetailHandler(c *gin.Context) {
	var payload entity.Room
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	room, err := r.roomUC.UpdateRoomDetail(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(c, room, "Updated")
}
func (r *RoomController) updateStatusHandler(c *gin.Context) {
	var payload entity.Room
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	room, err := r.roomUC.UpdateRoomStatus(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(c, room, "Ok")
}

func (r *RoomController) Route() {
	r.rg.POST(config.RoomCreate, r.authMiddleware.RequireToken("admin"), r.createHandler)
	r.rg.GET(config.RoomList, r.authMiddleware.RequireToken("employee", "admin", "ga"), r.listHandler)
	r.rg.GET(config.RoomGetById, r.authMiddleware.RequireToken("employee", "admin", "ga"), r.getHandler)
	r.rg.PUT(config.RoomUpdate, r.authMiddleware.RequireToken("admin"), r.updateDetailHandler)
	r.rg.PUT(config.RoomUpdateStatus, r.authMiddleware.RequireToken("admin", "ga"), r.updateStatusHandler)
}

func NewRoomController(roomUC usecase.RoomUseCase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup) *RoomController {
	return &RoomController{roomUC: roomUC, authMiddleware: authMiddleware, rg: rg}
}
