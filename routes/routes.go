package routes

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"example.com/project/handlers"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	
	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
