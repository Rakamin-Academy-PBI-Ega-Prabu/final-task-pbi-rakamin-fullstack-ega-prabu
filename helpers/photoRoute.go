package helpers

import (
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
	"userapp/initializers"
	"userapp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPhotoAssociation(c *gin.Context) (*gorm.Association, error) {
	u, _ := c.Get("user")

	user := u.(models.User)

	assoc := initializers.DB.Model(&user).Association("Photos")

	if assoc.Error != nil {
		return nil, assoc.Error
	}

	return assoc, nil
}

func GetPhoto(c *gin.Context) (*models.Photo, error) {
	var uri struct {
		ID string `uri:"photo_id" binding:"required"`
	}

	err := c.ShouldBindUri(&uri)

	if err != nil {
		return nil, errors.New("photo not found")
	}

	u, _ := c.Get("user")

	user := u.(models.User)

	var photo models.Photo

	initializers.DB.Model(&user).Where("ID = ?", uri.ID).Association("Photos").Find(&photo)

	if photo.ID == 0 {
		return nil, errors.New("photo not found")
	}

	return &photo, nil
}

func GetFile(c *gin.Context) (*multipart.FileHeader, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, errors.New("get form file error: " + err.Error())
	}

	errImage := IsFileImage(file)
	if errImage != nil {
		return nil, errImage
	}

	return file, nil
}

func GetPhotoUrl(c *gin.Context, file *multipart.FileHeader) (string, error) {
	user, _ := c.Get("user")

	userFilePath := filepath.Clean(GetUserFilePath(user.(models.User).Username))

	// Chech if directory can be created. If there's no directory, create one
	errDir := os.MkdirAll(userFilePath, 0750)
	if errDir != nil && !os.IsExist(errDir) {
		return "", errors.New("Can not make Directory for User: " + errDir.Error())
	}

	photoUrl := filepath.Join(userFilePath, SetFileName(file.Filename))

	return photoUrl, nil
}
