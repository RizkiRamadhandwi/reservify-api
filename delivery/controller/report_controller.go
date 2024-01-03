package controller

import (
	"booking-room-app/delivery/middleware"
	"booking-room-app/shared/common"
	"booking-room-app/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	reportUC       usecase.ReportUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (r *ReportController) downloadHandler(c *gin.Context) {
	rangeParam := c.Query("range")

	if rangeParam == "" || (rangeParam != "day" && rangeParam != "week" && rangeParam != "month" && rangeParam != "year") {
		common.SendErrorResponse(c, http.StatusBadRequest, "Invalid range parameter")
		return
	}

	_, err := r.reportUC.PrintAllReports(rangeParam)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "transaction.csv"))
	c.Header("Content-Type", "text/csv")
	c.File("public/transaction.csv")
}

func (r *ReportController) Route() {
	r.rg.GET("/reports/download", r.authMiddleware.RequireToken("admin"), r.downloadHandler)
}

func NewReportController(reportUC usecase.ReportUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *ReportController {
	return &ReportController{reportUC: reportUC, rg: rg, authMiddleware: authMiddleware}
}
