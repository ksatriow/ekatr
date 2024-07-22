package http

import (
	"github.com/gorilla/mux"
)

func NewRouter(userHandler *UserHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	router.HandleFunc("/login", userHandler.LoginUser).Methods("POST")
	router.HandleFunc("/user", userHandler.GetUserByID).Methods("GET")

	return router
}