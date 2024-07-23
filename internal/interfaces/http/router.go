package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(userHandler *UserHandler) *mux.Router {
	router := mux.NewRouter()
	
	router.Handle("/register", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.RegisterUser)))).Methods("POST")
	router.HandleFunc("/login", userHandler.LoginUser).Methods("POST")
	router.Handle("/user", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.GetUserByID)))).Methods("GET")
	router.Handle("/users", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.GetAllUsers)))).Methods("GET")
	return router
}