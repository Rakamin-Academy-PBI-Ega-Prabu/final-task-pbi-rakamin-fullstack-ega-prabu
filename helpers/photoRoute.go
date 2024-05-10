package helpers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"userapp/initializers"
	"userapp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPhotoAssociation(c *gin.Context) (*gorm.Association, bool) {
	u, _ := c.Get("user")

	user := u.(models.User)

	assoc := initializers.DB.Model(&user).Association("Photos")

	if assoc.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can not retrive photos",
		})

		return nil, false
	}

	return assoc, true
}

func GetPhoto(c *gin.Context) (models.Photo, bool) {
	var uri struct {
		ID string `uri:"photo_id" binding:"required"`
	}

	if c.ShouldBindUri(&uri) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read Photo ID",
		})

		return models.Photo{}, false
	}

	u, _ := c.Get("user")

	user := u.(models.User)

	var photo models.Photo

	initializers.DB.Debug().Model(&user).Where("ID = ?", uri.ID).Association("Photos").Find(&photo)

	if photo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not found",
		})

		return models.Photo{}, false
	}

	return photo, true
}

func GetFile(c *gin.Context) (*multipart.FileHeader, bool) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Get Form File Error: %s", err.Error()),
		})
		return nil, false
	}

	return file, true
}

func GetPhotoUrl(c *gin.Context, file *multipart.FileHeader) (string, bool) {
	user, _ := c.Get("user")

	userFilePath := filepath.Clean(GetUserFilePath(user.(models.User).Username))

	// Chech if directory can be created. If there's no directory, create one
	errDir := os.MkdirAll(userFilePath, 0750)
	if errDir != nil && !os.IsExist(errDir) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Can not make Directory for User: %s", errDir.Error()),
		})

		return "", false
	}

	photoUrl := filepath.Join(userFilePath, SetFileName(file.Filename))

	return photoUrl, true
}
