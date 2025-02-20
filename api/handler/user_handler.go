package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/golang-stream/internal/dto"
	"github.com/wildanasyrof/golang-stream/internal/service"
	"github.com/wildanasyrof/golang-stream/pkg/response"
	"github.com/wildanasyrof/golang-stream/pkg/validation"
)

type UserHandler struct {
	validation  validation.Validator
	userService service.UserService
}

func NewUserHandler(v validation.Validator, userService service.UserService) *UserHandler {
	return &UserHandler{validation: v, userService: userService}
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
