package controllers

import (
	"fmt"
	"net/http"

	"github.com/drsimplegraffiti/ref/initializers"
	"github.com/drsimplegraffiti/ref/models"
	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	var booking models.Booking

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + ` Fields are empty`})
		return
	}

	// check if body is empty or not
	if booking.Description == ""  {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
	}

	// Get the user ID from the token
	userID := c.MustGet("userID").(uint)

	booking.UserID = userID
	initializers.DB.Create(&booking)

	c.JSON(http.StatusOK, gin.H{"message": "Booking created successfully"})
}

func GetAllBookings(c *gin.Context) {
	var bookings []models.Booking

	// Get the user ID from the token
	userID := c.MustGet("userID").(uint)
	fmt.Println(userID)
	

	initializers.DB.Where("user_id = ?", userID).Find(&bookings)

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

func GetBooking(c *gin.Context) {
	var booking models.Booking

	bookingID := c.Param("id")

	initializers.DB.Where("id = ?", bookingID).First(&booking)

	c.JSON(http.StatusOK, gin.H{"booking": booking})
}