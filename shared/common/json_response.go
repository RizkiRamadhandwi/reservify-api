package common

import (
	"booking-room-app/shared/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendCreateResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusCreated, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusCreated,
			Message: message,
		},
		Data: data,
	})
}

func SendSingleResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

func SendErrorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, &model.Status{
		Code:    code,
		Message: message,
	})
}

func SendNoContentResponse(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}