package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	//add more fields here
	// FirstName string `gorm:"not null"`
	LastName string
	FirstName string
	Age int
}

