package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/go-first-api/schemas"
)

func DeleteOpeningHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsValid("id", "query-parameter").Error())
		return
	}

	opening := schemas.Opening{}
	// On Gorm, we first find the entity (opening) and then, we delete

	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opening with id: %v not found", id))
		return
	}

	// Delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %v", id))
		return
	}

	sendSuccess(c, "delete-opening", opening)
}
