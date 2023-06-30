package initializers

import (
	"log"

	"github.com/joho/godotenv"
)


func LoadEnvVariables(){
    err := godotenv.Load()
    if err != nil { // error must return nil
        log.Fatal("Error loading .env file" + err.Error())
    }
}
