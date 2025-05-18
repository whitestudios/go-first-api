package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/go-first-api/schemas"
)

func ShowOpeningHandler(c *gin.Context) {
	opening := schemas.Opening{}
	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsValid("id", "query-parameter").Error())
		return
	}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(c, "show-opening", opening)
}
