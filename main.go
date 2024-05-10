package main

import (
	"userapp/controllers"
	"userapp/initializers"
	"userapp/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.InitializeApp()
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
		userRoute.PUT("/:user_id/edit", controllers.UserEdit)
		userRoute.POST("/:user_id/change-password", controllers.UserChangePassword)
		userRoute.DELETE("/:user_id/delete", controllers.UserDelete)
	}

	// Photo Routes
	photoRoute := r.Group("/photos")

	photoRoute.Use(middleware.RequireAuth)
	{
		photoRoute.GET("/", controllers.PhotoIndex)
		photoRoute.POST("", controllers.PhotoCreate)
		photoRoute.GET("/:photo_id", controllers.PhotoGet)
		photoRoute.PUT("/:photo_id", controllers.PhotoEdit)
		photoRoute.PUT("/:photo_id/change-photo", controllers.PhotoChange)
		photoRoute.DELETE("/:photo_id", controllers.PhotoDelete)
	}

	r.Run()
}
