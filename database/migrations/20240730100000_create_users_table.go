package main

import (
	"go-auth/database"
	"go-auth/app/Models"
	"go-auth/config"
	"log"
)

func main() {
	config := config.LoadConfig()
	database.Connect(config)

	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	log.Println("Migration successful.")
}
