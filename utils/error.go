package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func SendErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, ErrorResponse{
		Message: message,
		Status:  status,
	})
}
