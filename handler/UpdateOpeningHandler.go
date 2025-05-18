package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/go-first-api/schemas"
)

func UpdateOpeningHandler(c *gin.Context) {
	request := UpdateOpeningRequest{}
	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsValid("id", "query-parameter").Error())
		return
	}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "opening not found")
	}
	// update opening
	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// Save Opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errf("Error saving opening's update: %s", err)
		sendError(c, http.StatusInternalServerError, "Error saving opening's update")
		return
	}
	sendSuccess(c, "update-opening", opening)
}
