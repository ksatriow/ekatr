package user

import (
	"errors"
	"fmt"

	"ekatr/internal/domain/user"
	"ekatr/internal/infrastructure/persistence/postgresql"
	"ekatr/internal/utils"
)

type UserService struct {
	repo *postgresql.UserRepository
}

func NewUserService(repo *postgresql.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(dto RegisterUserDTO) error {
	// Check if email already exists
	existingUser, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		return fmt.Errorf("error finding user by email: %w", err)
	}
	if existingUser != nil {
		return errors.New("email already exists")
	}

	// Check if username already exists
	existingUser, err = s.repo.FindByUsername(dto.Username)
	if err != nil {
		return fmt.Errorf("error finding user by username: %w", err)
	}
	if existingUser != nil {
		return errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	// Create new user
	userType := user.UserType(dto.Type)
	newUser := user.NewUser(dto.Username, hashedPassword, dto.Email, userType, dto.ProfilePhoto)

	// Save user to the repository
	if err := s.repo.Save(newUser); err != nil {
		return fmt.Errorf("error saving user: %w", err)
	}

	return nil
}

func (s *UserService) LoginUser(dto LoginUserDTO) (*user.User, error) {
	u, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		return nil, fmt.Errorf("error finding user by email: %w", err)
	}
	if u == nil {
		return nil, errors.New("user not found")
	}

	// Check password
	if !utils.CheckPasswordHash(dto.Password, u.Password) {
		return nil, errors.New("invalid password")
	}

	return u, nil
}

func (s *UserService) GetUserByID(id int) (*user.User, error) {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("error finding user by ID: %w", err)
	}
	if u == nil {
		return nil, errors.New("user not found")
	}

	return u, nil
}

func (s *UserService) GetAllUsers() ([]*user.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error finding all users: %w", err)
	}
	return users, nil
}
