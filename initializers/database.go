package initializers

import (
	"fmt"
	"os"
	"userapp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Current_User models.User

func ConnectToDatabase() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to Database")
	}

	fmt.Println(DB)
}

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}
