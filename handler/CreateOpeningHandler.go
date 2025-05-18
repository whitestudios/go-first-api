package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/go-first-api/schemas"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param opening body CreateOpeningRequest true "Opening body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(c *gin.Context) {
	var request CreateOpeningRequest

	// Return a DTO
	c.BindJSON(&request)

	// Validate request (a opening schema DTO )
	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Create a real schema struct with DTO's fields
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Link:     request.Link,
		Location: request.Location,
		Salary:   request.Salary,
		Remote:   *request.Remote,
	}

	// Create in database the opening
	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("error creating opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	// Return Opening
	sendSuccess(c, "create-opening", opening)

	return
}
