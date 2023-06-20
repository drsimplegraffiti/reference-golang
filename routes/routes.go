package routes

import (
	"github.com/drsimplegraffiti/ref/controllers"
	"github.com/drsimplegraffiti/ref/utils"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)

	openRoutes := router.Group("/api")
	{
		openRoutes.GET("/bookings/:id", controllers.GetBooking)
	}
	
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(utils.AuthMiddleware())
	{
		protectedRoutes.GET("/bookings", controllers.GetAllBookings)
		protectedRoutes.POST("/bookings", controllers.CreateBooking)
	}
}
