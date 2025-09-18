package main

import (
	"runtime"
	"github.com/minhanhbb/ecom-golang/config"
	"github.com/minhanhbb/ecom-golang/database"
	"github.com/minhanhbb/ecom-golang/router"
)

func main() {
	// Tối ưu sử dụng CPU
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	database.Connect(cfg)

	// Setup router
	r := router.SetupRouter()

	// Start the application
	r.Run() // listens and serves on 0.0.0.0:8080 (for windows "localhost:8080")
}
