package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(userHandler *UserHandler, productHandler *ProductHandler) *mux.Router {
	router := mux.NewRouter()
	
	// Auth routes
	router.Handle("/register", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.RegisterUser)))).Methods(http.MethodPost)
	router.HandleFunc("/login", userHandler.LoginUser).Methods(http.MethodPost)
	router.Handle("/user", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.GetUserByID)))).Methods(http.MethodGet)
	router.Handle("/users", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.GetAllUsers)))).Methods(http.MethodGet)
	router.Handle("/user", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.DeleteUserByID)))).Methods(http.MethodDelete)
	router.Handle("/user", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(userHandler.UpdateUserByID)))).Methods(http.MethodPut)
	
    // Product routes
    router.Handle("/products", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(productHandler.CreateProduct)))).Methods(http.MethodPost)
    router.Handle("/product", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(productHandler.GetProductByID)))).Methods(http.MethodGet)
    router.Handle("/products", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(productHandler.GetAllProducts)))).Methods(http.MethodGet)
	router.Handle("/product", AuthMiddleware(RoleMiddleware(ownerRole, http.HandlerFunc(productHandler.UpdateProduct)))).Methods(http.MethodPut)
	return router
}