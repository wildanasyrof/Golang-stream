package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/golang-stream/internal/dto"
	"github.com/wildanasyrof/golang-stream/internal/service"
	"github.com/wildanasyrof/golang-stream/pkg/response"
	"github.com/wildanasyrof/golang-stream/pkg/validation"
	"go.uber.org/zap"
)

type CategoryHandler struct {
	CategoryService service.CategoryService
	logger          *zap.SugaredLogger
	validation      validation.Validator
}

func NewCategoryHandler(categoryService service.CategoryService, logger *zap.SugaredLogger, validation validation.Validator) *CategoryHandler {
	return &CategoryHandler{CategoryService: categoryService, logger: logger, validation: validation}
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateCategoryRequest

	if err := h.validation.ValidateBody(c, &req); err != nil {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Validation error", err)
	}

	category, err := h.CategoryService.CreateCategory(req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Create Category Failed!", err.Error())
	}

	return response.Created(c, "Category created successfully", category)
}
