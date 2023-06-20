package main

import (
	"os"

	"github.com/drsimplegraffiti/ref/initializers"
	"github.com/drsimplegraffiti/ref/routes"
	"github.com/gin-gonic/gin"
)

func init(){
	// Load environment variables
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()	
}

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()
	routes.InitializeRoutes(router)
	router.Run(":" + os.Getenv("PORT"))
}
