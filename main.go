package main

import (
	"log"
	"net/http"

	_ "example.com/project/docs"

	"example.com/project/config"
	"example.com/project/models"
	"example.com/project/routes"
)

// @title Go Backend API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
func main() {
	config.ConnectDatabase()

	// Auto migrate tables
	config.DB.AutoMigrate(&models.User{})

	router := routes.RegisterRoutes()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
