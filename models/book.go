package models

import (
	"gorm.io/gorm"
)


type Booking struct {
	gorm.Model
	UserID      uint
	Description string
}


