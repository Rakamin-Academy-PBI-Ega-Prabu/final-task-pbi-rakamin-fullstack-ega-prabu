package helpers

import (
	"errors"
	"userapp/initializers"
	"userapp/models"

	"github.com/gin-gonic/gin"
)

func GetUserFromUri(c *gin.Context) (*models.User, error) {
	var uri struct {
		ID string `uri:"user_id" binding:"required"`
	}

	err := c.ShouldBindUri(&uri)

	if err != nil {
		return nil, err
	}

	var user models.User

	initializers.DB.First(&user, uri.ID)

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
