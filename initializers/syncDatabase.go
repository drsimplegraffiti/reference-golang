package initializers

import (
	"github.com/drsimplegraffiti/ref/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Booking{})
}