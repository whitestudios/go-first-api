package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/go-first-api/schemas"
)

func ListOpeningsHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(c, "list-openings", openings)
}
