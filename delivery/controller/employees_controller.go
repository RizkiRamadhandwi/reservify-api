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

type EmployeeController struct {
	employeeUC     usecase.EmployeesUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (e *EmployeeController) createHandler(ctx *gin.Context) {

	var payload entity.Employee
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	employee, err := e.employeeUC.RegisterNewEmployee(payload)

	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return

	}
	common.SendCreateResponse(ctx, employee, "Created")
}

// read by
func (e *EmployeeController) getByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	employee, err := e.employeeUC.FindEmployeesByID(id)
	if err != nil {

		common.SendErrorResponse(ctx, http.StatusNotFound, "Employee with ID "+id+" not found")
		return
	}
	common.SendSingleResponse(ctx, employee, "Ok")
}
func (e *EmployeeController) getByUsernameHandler(ctx *gin.Context) {
	username := ctx.Param("user")
	employee, err := e.employeeUC.FindEmployeesByUsername(username)
	if err != nil {

		common.SendErrorResponse(ctx, http.StatusNotFound, "Employee with Username "+username+" not found")
		return
	}
	common.SendSingleResponse(ctx, employee, "Ok")
}

// update

func (e *EmployeeController) putHandler(ctx *gin.Context) {
	var payload entity.Employee
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Failed to bind data")
		return
	}
	employee, err := e.employeeUC.UpdateEmployee(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	common.SendSingleResponse(ctx, employee, "Updated Successfully")

}

// pagination
func (e *EmployeeController) ListHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "5"))

	employees, paging, err := e.employeeUC.ListAll(page, size)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var response []interface{}

	for _, v := range employees {
		response = append(response, v)
	}
	common.SendPagedResponse(ctx, response, paging, "Ok")
}

// route
func (e *EmployeeController) Route() {
	e.rg.GET(config.EmployeesGetById, e.authMiddleware.RequireToken("admin", "employee", "ga"), e.getByIdHandler)
	e.rg.GET(config.EmployeesGetByUsername, e.authMiddleware.RequireToken("admin", "employee", "ga"), e.getByUsernameHandler)
	e.rg.POST(config.EmployeesCreate, e.authMiddleware.RequireToken("admin"), e.createHandler)
	e.rg.PUT(config.EmployeesUpdate, e.authMiddleware.RequireToken("admin"), e.putHandler)
	e.rg.GET(config.EmployeesList, e.authMiddleware.RequireToken("admin", "employee", "ga"), e.ListHandler)
}

func NewEmployeeController(employeeUC usecase.EmployeesUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *EmployeeController {
	return &EmployeeController{
		employeeUC:     employeeUC,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}
