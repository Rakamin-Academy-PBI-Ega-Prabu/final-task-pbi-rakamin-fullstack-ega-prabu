package controllers

import (
	"net/http"
	"userapp/initializers"
	"userapp/models"

	"github.com/gin-gonic/gin"
)

func PhotoIndex(c *gin.Context) {
	assoc := initializers.DB.Model(&initializers.Current_User).Association("Photos")

	if assoc.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can not retrive photos",
		})
	}

	var photos []models.Photo
	assoc.Find(&photos)

	if photos == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "No photos yet",
		})

		return
	}

	c.JSON(http.StatusOK, photos)
}

func PhotoCreate(c *gin.Context) {
	assoc := initializers.DB.Model(&initializers.Current_User).Association("Photos")

	if assoc.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can not retrive photos",
		})
	}

	var body struct {
		TItle   string `form:"title" binding:"required"`
		Caption string `form:"caption" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	file, _ := c.FormFile("file")

	c.JSON(http.StatusAccepted, file)
}
