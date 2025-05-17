package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "POST Opening"})
}

func ShowOpeningHandler(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "GET Opening"})
}

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "DELETE Opening"})
}

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "PUT Opening"})
}

func ListOpeningsHandler(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "GET Openings"})
}
