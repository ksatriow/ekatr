package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ekatr/internal/application/user"
	"ekatr/internal/logger"
	"ekatr/internal/utils"
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

	// Convert UserType to string
	role := string(u.Type)

	// Generate token
	token, err := utils.GenerateToken(u.Username, role)
	if err != nil {
		logger.ErrorLogger.Printf("Error generating token: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": token}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
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

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		logger.ErrorLogger.Printf("Error getting all users: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.DeleteUserByID(id)
	if err != nil {
		logger.ErrorLogger.Printf("Error deleting user by ID: %v", err)
		if err.Error() == "user not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Return a success message with status 200 OK
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "User deleted successfully"}
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
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

	var dto user.UpdateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		logger.ErrorLogger.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateUser(id, dto)
	if err != nil {
		logger.ErrorLogger.Printf("Error updating user: %v", err)
		if err.Error() == "user not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Return a success message with status 200 OK
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "User updated successfully"}
	json.NewEncoder(w).Encode(response)
}