package main

import (
	"log"
	"net/http"

	_ "example.com/project/docs"

	"example.com/project/config"
	"example.com/project/internal/user"
	"example.com/project/routes"
)

// @title Go Backend API
// @version 1.0
// @host localhost:8080
// @BasePath /

func main() {

	config.ConnectDatabase()

	config.DB.AutoMigrate(&user.User{})

	// Manual Dependency Injection
	userRepo := user.NewRepository(config.DB)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	router := routes.RegisterRoutes(userHandler)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
