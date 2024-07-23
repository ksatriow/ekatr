package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(userHandler *UserHandler) *mux.Router {
	router := mux.NewRouter()
	
	router.Handle("/register", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.RegisterUser)))).Methods(http.MethodPost)
	router.HandleFunc("/login", userHandler.LoginUser).Methods(http.MethodPost)
	router.Handle("/user", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.GetUserByID)))).Methods(http.MethodGet)
	router.Handle("/users", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.GetAllUsers)))).Methods(http.MethodGet)
	router.Handle("/user", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.DeleteUserByID)))).Methods(http.MethodDelete)
	router.Handle("/user", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.UpdateUserByID)))).Methods(http.MethodPut)
	return router
}