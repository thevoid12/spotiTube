package main

import (
	"spotiTube/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.GET("/", handler.IndexPageHandler)
	router.GET("/login", handler.LoginHandler)
	router.GET("/home", handler.HomePageHandler)

	// Start the server
	router.Run(":8080")
}
