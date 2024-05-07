package main

import (
	"userapp/controllers"
	"userapp/initializers"
	"userapp/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/validate", middleware.RequireAuth, controllers.ValidateUser)

	userRoute := r.Group("/users")
	{
		// User Routes
		userRoute.GET("/", controllers.UserIndex)
		userRoute.POST("/register", controllers.UserSignUp)
		userRoute.POST("/login", controllers.UserLogin)
		userRoute.POST("/:user_id/edit", controllers.UserEdit)
		userRoute.DELETE("/:user_id/delete", controllers.UserDelete)
		userRoute.POST("/:user_id/change-password", controllers.UserChangePassword)

		// Photo Routes
		photoRoute := r.Group("/photos")

		photoRoute.Use(middleware.RequireAuth)
		{
			photoRoute.GET("/", controllers.PhotoIndex)
			photoRoute.POST("/", controllers.PhotoCreate)

		}
	}

	r.Run()
}
