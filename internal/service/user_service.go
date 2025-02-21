package service

import (
	"errors"

	"github.com/wildanasyrof/golang-stream/internal/dto"
	"github.com/wildanasyrof/golang-stream/internal/entity"
	"github.com/wildanasyrof/golang-stream/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(req dto.RegisterRequest) (*entity.User, error)
	Login(req dto.LoginRequest) (*entity.User, error)
	GetProfile(id uint) (*entity.User, error)
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword, password string) error
	UpdateUser(userId uint, req dto.UpdateUserRequest) (*entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

// UpdateUser implements UserService.
func (u *userService) UpdateUser(userId uint, req dto.UpdateUserRequest) (*entity.User, error) {
	user, err := u.userRepo.FindByID(userId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update fields if provided
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.NewPassword != "" && req.LastPassword != "" {

		if err := u.VerifyPassword(user.Password, req.LastPassword); err != nil {
			return nil, errors.New("invalid last password")
		}

		// Hash password
		hashedPassword, err := u.HashPassword(req.NewPassword)
		if err != nil {
			return nil, err
		}

		user.Password = string(hashedPassword)
	}
	if req.ProfileImg != "" {
		user.ProfileImg = req.ProfileImg
	}

	err = u.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Get Profile
func (u *userService) GetProfile(id uint) (*entity.User, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login implements UserService.
func (u *userService) Login(req dto.LoginRequest) (*entity.User, error) {
	user, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err := u.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, err
	}

	return user, nil
}

// HashPassword implements UserService.
func (u *userService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	return string(hashedPassword), nil
}

// VerifyPassword implements UserService.
func (u *userService) VerifyPassword(hashedPassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return errors.New("invalid email or password")
	}

	return nil
}

func (u *userService) RegisterUser(req dto.RegisterRequest) (*entity.User, error) {
	// Check if user already exists
	existingUser, _ := u.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := u.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user entity
	user := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Save user
	if err := u.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
