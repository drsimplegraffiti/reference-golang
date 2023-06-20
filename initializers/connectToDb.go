package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Assign value to the existing DB variable
	fmt.Println("Connection Opened to Database")
	if err != nil {
		panic("failed to connect database")
	}

}
