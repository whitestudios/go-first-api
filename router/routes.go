package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/whitestudios/go-first-api/docs"
	"github.com/whitestudios/go-first-api/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	BasePath := "/api/v1"

	handler.InitializeHandler()

	docs.SwaggerInfo.BasePath = BasePath

	v1 := router.Group(BasePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
