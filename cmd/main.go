package main

import (
	"log"
	"net/http"

	"ekatr/internal/application/user"
	"ekatr/internal/infrastructure/persistence/postgresql"
	httpInterface "ekatr/internal/interfaces/http"
	"ekatr/internal/logger"
)

func main() {
	logger.Init()
	
	dataSourceName := "user=postgres password=komar123 dbname=ekatrdb host=localhost sslmode=disable"
	

	db, err := postgresql.NewDB(dataSourceName)
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	userRepo := postgresql.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := httpInterface.NewUserHandler(userService)

	router := httpInterface.NewRouter(userHandler)

	logger.InfoLogger.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.ErrorLogger.Fatalf("could not start the server: %v", err)
	}
}
