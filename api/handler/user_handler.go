package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/golang-stream/internal/dto"
	"github.com/wildanasyrof/golang-stream/internal/service"
	"github.com/wildanasyrof/golang-stream/pkg/auth"
	"github.com/wildanasyrof/golang-stream/pkg/response"
	"github.com/wildanasyrof/golang-stream/pkg/validation"
	"go.uber.org/zap"
)

type UserHandler struct {
	validation  validation.Validator
	userService service.UserService
	AuthService *auth.AuthService
	logger      *zap.SugaredLogger
}

func NewUserHandler(v validation.Validator, userService service.UserService, auth *auth.AuthService, logger *zap.SugaredLogger) *UserHandler {
	return &UserHandler{validation: v, userService: userService, AuthService: auth, logger: logger}
}

// RegisterUserHandler is a handler for registering a new user
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var req dto.RegisterRequest

	if err := h.validation.ValidateBody(c, &req); err != nil {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Validation error", err)
	}

	user, err := h.userService.RegisterUser(req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Register Failed!", err.Error())
	}

	return response.Created(c, "User registered successfully", user)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest

	if err := h.validation.ValidateBody(c, &req); err != nil {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Validation error", err)
	}

	user, err := h.userService.Login(req)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Login Failed!", err.Error())
	}

	token, err := h.AuthService.GenerateToken(user.ID, user.Role)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to generate token", err)
	}

	return response.Success(c, "Login successful", fiber.Map{
		"token": token,
		"user":  user,
	})
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := uint(c.Locals("userID").(float64))

	h.logger.Info("User ID = ", userID)
	user, err := h.userService.GetProfile(userID)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "User not found", err)
	}

	return response.Success(c, "Succes get user data", user)
}
