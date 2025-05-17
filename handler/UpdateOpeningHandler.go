package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "PUT Opening"})
}
