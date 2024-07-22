package user

import (
	"errors"

	"ekatr/internal/domain/user"
	"ekatr/internal/logger"
	"ekatr/internal/utils"
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(dto RegisterUserDTO) error {
	existingUserByEmail, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		logger.ErrorLogger.Printf("Error finding user by email: %v", err)
		return errors.New("internal server error")
	}
	if existingUserByEmail != nil {
		logger.InfoLogger.Println("Email already exists")
		return errors.New("email already exists")
	}

	existingUserByUsername, err := s.repo.FindByUsername(dto.Username)
	if err != nil {
		logger.ErrorLogger.Printf("Error finding user by username: %v", err)
		return errors.New("internal server error")
	}
	if existingUserByUsername != nil {
		logger.InfoLogger.Println("Username already exists")
		return errors.New("username already exists")
	}

	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		logger.ErrorLogger.Printf("Error hashing password: %v", err)
		return errors.New("internal server error")
	}

	newUser := user.NewUser(dto.Username, hashedPassword, dto.Email, user.UserType(dto.Type))
	err = s.repo.Save(newUser)
	if err != nil {
		logger.ErrorLogger.Printf("Error saving user: %v", err)
		return errors.New("internal server error")
	}
	return nil
}

func (s *UserService) LoginUser(dto LoginUserDTO) (*user.User, error) {
	existingUser, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		logger.ErrorLogger.Printf("Error finding user by email: %v", err)
		return nil, errors.New("internal server error")
	}
	if existingUser == nil {
		logger.InfoLogger.Println("User not found")
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(dto.Password, existingUser.Password) {
		logger.InfoLogger.Println("Invalid password")
		return nil, errors.New("invalid credentials")
	}

	return existingUser, nil
}

func (s *UserService) GetUserByID(id int) (*user.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		logger.ErrorLogger.Printf("Error finding user by ID: %v", err)
		return nil, errors.New("internal server error")
	}
	if user == nil {
		logger.InfoLogger.Printf("User with ID %d not found", id)
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]*user.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		logger.ErrorLogger.Printf("Error finding all users: %v", err)
		return nil, errors.New("internal server error")
	}
	return users, nil
}