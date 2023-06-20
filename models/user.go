package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	LastName string
	FirstName string
	Age int
	ProfilePic string
}

