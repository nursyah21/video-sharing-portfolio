package main

import (
	"video-go/database"
	"video-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDB()
	defer database.CloseDB()

	routes.SetupRoutes(router)

	router.Run(":5000")
}
