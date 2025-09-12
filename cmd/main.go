package main

import (
	"go-auth/config"
	"go-auth/database"
	"go-auth/app/Controller/Auth"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	database.Connect(cfg)

	// Set up the router
	r := gin.Default()
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	r.POST("/logout", auth.Logout)

	// Start the application
	r.Run() // listens and serves on 0.0.0.0:8080 (for windows "localhost:8080")
}
