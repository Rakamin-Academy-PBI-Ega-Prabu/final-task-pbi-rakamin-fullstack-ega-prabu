package initializers

import (
	"math/rand"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitializeApp() {
	// New random seed
	rand.New(rand.NewSource(time.Now().UnixMilli()))

	// Initialize defined upload folder
	// Chech if directory can be created. If there's no directory, create one
	errDir := os.Mkdir(BasePath, 0750)
	if errDir != nil && !os.IsExist(errDir) {
		panic(errDir.Error())

	}

	// use a single instance of Validate, it caches struct info
	Validate = validator.New(validator.WithRequiredStructEnabled())

	// Load .env varibles
	LoadEnvVariables()

	// Initialize Database Used
	ConnectToDatabase()
	SyncDatabase()
}
