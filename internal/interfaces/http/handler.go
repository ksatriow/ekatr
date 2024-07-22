package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ekatr/internal/application/user"
	"ekatr/internal/logger"
)

type UserHandler struct {
	service *user.UserService
}

func NewUserHandler(service *user.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var dto user.RegisterUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		logger.ErrorLogger.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.RegisterUser(dto)
	if err != nil {
		logger.ErrorLogger.Printf("Error registering user: %v", err)
		if err.Error() == "email already exists" || err.Error() == "username already exists" {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	logger.InfoLogger.Printf("User registered successfully: %v", dto.Username)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var dto user.LoginUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		logger.ErrorLogger.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.service.LoginUser(dto)
	if err != nil {
		logger.ErrorLogger.Printf("Error logging in user: %v", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	logger.InfoLogger.Printf("User logged in successfully: %v", u.Username)
	json.NewEncoder(w).Encode(u)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing user ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		logger.ErrorLogger.Printf("Error getting user by ID: %v", err)
		if err.Error() == "user not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
