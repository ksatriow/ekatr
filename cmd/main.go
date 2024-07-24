package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"database/sql"

	"ekatr/internal/application/user"
	"ekatr/internal/application/product"
	"ekatr/internal/infrastructure/persistence/postgresql"
	httpInterface "ekatr/internal/interfaces/http"
	"ekatr/internal/logger"
)

func main() {
	logger.Init()
	
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbSSLMode := os.Getenv("DB_SSL_MODE")

    connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSSLMode
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

	userRepo := postgresql.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := httpInterface.NewUserHandler(userService)

    productRepo := postgresql.NewProductRepository(db)
    productService := product.NewProductService(productRepo)
    productHandler := httpInterface.NewProductHandler(productService)

	router := httpInterface.NewRouter(userHandler, productHandler)

	logger.InfoLogger.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.ErrorLogger.Fatalf("could not start the server: %v", err)
	}
}
