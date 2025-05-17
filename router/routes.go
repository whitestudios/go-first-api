package router

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/opening")
		v1.POST("/opening")
		v1.DELETE("/opening")
		v1.PUT("/opening")
		v1.GET("/openings")

	}
}
