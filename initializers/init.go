package initializers

import (
	"math/rand"
	"os"
	"time"

	"github.com/asaskevich/govalidator"
)

func InitializeApp() {
	// New random seed
	rand.New(rand.NewSource(time.Now().UnixMilli()))

	// Initialize defined upload folder
	// Chech if directory can be created. If there's no directory, create one
	errDir := os.Mkdir(BasePath, 0750)
	if errDir != nil && !os.IsExist(errDir) {
		panic(errDir.Error())

	}

	// Set field required by default
	govalidator.SetFieldsRequiredByDefault(true)

	// Load .env varibles
	LoadEnvVariables()

	// Initialize Database Used
	ConnectToDatabase()
	SyncDatabase()
}
