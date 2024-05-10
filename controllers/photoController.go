package controllers

import (
	"fmt"
	"net/http"
	"os"
	"userapp/helpers"
	"userapp/initializers"
	"userapp/models"

	"github.com/gin-gonic/gin"
)

func PhotoIndex(c *gin.Context) {
	assoc, status := helpers.GetPhotoAssociation(c)

	if !status {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can not read photos",
		})

		return
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
	assoc, status := helpers.GetPhotoAssociation(c)

	if !status {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can not read photos",
		})

		return
	}

	var body struct {
		Title   string `form:"title" binding:"required"`
		Caption string `form:"caption" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	file, statusFile := helpers.GetFile(c)

	if !statusFile {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get file",
		})

		return
	}

	photoUrl, statusUrl := helpers.GetPhotoUrl(c, file)

	if !statusUrl {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get photo url",
		})

		return
	}

	photo := models.Photo{Title: body.Title, Caption: body.Caption, PhotoUrl: photoUrl}

	assoc.Append(&photo)

	if err := c.SaveUploadedFile(file, photoUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("upload file err: %s", err.Error()),
		})

		return
	}

	c.JSON(http.StatusAccepted, photo)
}

func PhotoGet(c *gin.Context) {
	var uri struct {
		ID string `uri:"photo_id" binding:"required"`
	}

	if c.ShouldBindUri(&uri) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read Photo ID",
		})

		return
	}

	user, _ := c.Get("user")

	var photo models.Photo

	initializers.DB.Model(&user).Where("ID = ?", uri.ID).Association("Photos").Find(&photo)

	if photo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not found",
		})

		return
	}

	c.JSON(http.StatusAccepted, photo)
}

func PhotoEdit(c *gin.Context) {
	photo, status := helpers.GetPhoto(c)

	if !status {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can not read photo",
		})
		return
	}

	var body struct {
		Title   string `form:"title"`
		Caption string `form:"caption"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can not read body",
		})

		return
	}

	initializers.DB.Model(&photo).Select("Title", "Caption").Updates(models.Photo{Title: body.Title, Caption: body.Caption})

	c.JSON(http.StatusOK, gin.H{
		"message": photo,
	})
}

func PhotoChange(c *gin.Context) {
	photo, status := helpers.GetPhoto(c)

	if !status || photo.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can not read photo",
		})
		return
	}

	file, statusFile := helpers.GetFile(c)

	if !statusFile {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get file",
		})

		return
	}

	photoUrl, statusUrl := helpers.GetPhotoUrl(c, file)

	if !statusUrl {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get photo url",
		})

		return
	}

	initializers.DB.Model(&photo).Update("photo_url", photoUrl)

	if err := c.SaveUploadedFile(file, photoUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("upload file err: %s", err.Error()),
		})
		return
	}

	// Change old file name
	os.Rename(photo.PhotoUrl, photo.PhotoUrl+"(old)")

	c.JSON(http.StatusOK, gin.H{
		"message": photo,
	})
}

func PhotoDelete(c *gin.Context) {
	var uri struct {
		ID string `uri:"photo_id" binding:"required"`
	}

	if c.ShouldBindUri(&uri) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read Photo ID",
		})

		return
	}

	user, _ := c.Get("user")

	var photo models.Photo

	initializers.DB.Model(&user).Where("ID = ?", uri.ID).Association("Photos").Find(&photo)

	if photo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not found",
		})

		return
	}

	os.Rename(photo.PhotoUrl, photo.PhotoUrl+"(deleted)")

	initializers.DB.Model(&user).Association("Photos").Delete(photo)

	c.JSON(http.StatusOK, gin.H{})

}
