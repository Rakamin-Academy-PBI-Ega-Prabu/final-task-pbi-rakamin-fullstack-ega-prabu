package controllers

import (
	"net/http"
	"os"
	"userapp/helpers"
	"userapp/initializers"
	"userapp/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PhotoBody struct {
	Title   string `form:"title" validate:"required,max=100"`
	Caption string `form:"caption" validate:"required,max=255"`
}

func PhotoIndex(c *gin.Context) {
	assoc, err := helpers.GetPhotoAssociation(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
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
	assoc, errAssoc := helpers.GetPhotoAssociation(c)

	if errAssoc != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errAssoc.Error(),
		})

		return
	}

	var body PhotoBody

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	errValidation := initializers.Validate.Struct(body)
	if errValidation != nil {
		var validError []string
		for _, err := range errValidation.(validator.ValidationErrors) {
			validError = append(validError, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validError,
		})

		return
	}

	file, errFile := helpers.GetFile(c)

	if errFile != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errFile.Error(),
		})

		return
	}

	photoUrl, errUrl := helpers.GetPhotoUrl(c, file)

	if errUrl != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errUrl.Error(),
		})

		return
	}

	photo := models.Photo{Title: body.Title, Caption: body.Caption, PhotoUrl: photoUrl}

	assoc.Append(&photo)

	if err := c.SaveUploadedFile(file, photoUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "upload file err: " + err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, photo)
}

func PhotoGet(c *gin.Context) {
	photo, err := helpers.GetPhoto(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, photo)
}

func PhotoEdit(c *gin.Context) {
	photo, err := helpers.GetPhoto(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var body PhotoBody

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can not read body",
		})

		return
	}

	errValidation := initializers.Validate.Struct(body)
	if errValidation != nil {
		var validError []string
		for _, err := range errValidation.(validator.ValidationErrors) {
			validError = append(validError, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validError,
		})

		return
	}

	initializers.DB.Model(&photo).Select("Title", "Caption").Updates(models.Photo{Title: body.Title, Caption: body.Caption})

	c.JSON(http.StatusOK, gin.H{
		"message": photo,
	})
}

func PhotoChange(c *gin.Context) {
	photo, errPhoto := helpers.GetPhoto(c)

	if errPhoto != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errPhoto.Error(),
		})

		return
	}

	file, errFile := helpers.GetFile(c)

	if errFile != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errFile.Error(),
		})

		return
	}

	photoUrl, errUrl := helpers.GetPhotoUrl(c, file)

	if errUrl != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errUrl.Error(),
		})
		return
	}

	// Change old file name
	os.Rename(photo.PhotoUrl, photo.PhotoUrl+"(old)")

	initializers.DB.Model(&photo).Update("photo_url", photoUrl)

	if err := c.SaveUploadedFile(file, photoUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "upload file err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": photo,
	})
}

func PhotoDelete(c *gin.Context) {
	photo, errPhoto := helpers.GetPhoto(c)

	if errPhoto != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errPhoto.Error(),
		})

		return
	}

	os.Rename(photo.PhotoUrl, photo.PhotoUrl+"(deleted)")

	user, _ := c.Get("user")

	initializers.DB.Model(&user).Association("Photos").Delete(photo)

	c.JSON(http.StatusOK, gin.H{})

}
