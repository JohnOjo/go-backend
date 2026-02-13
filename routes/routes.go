package routes

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"

	"example.com/project/internal/user"
)

func RegisterRoutes(userHandler *user.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
