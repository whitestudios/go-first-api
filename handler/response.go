package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

func sendSuccess(c *gin.Context, op string, data any) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}
