package handler

import (
	"strconv"

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

func (h *CategoryHandler) GetAll(c *fiber.Ctx) error {
	categories, err := h.CategoryService.GetAllCategories()
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed get Categories!", err.Error())
	}

	return response.Success(c, "Succes get All Categories", categories)
}

func (h *CategoryHandler) Destroy(c *fiber.Ctx) error {
	catId := c.Params("id")
	id, err := strconv.Atoi(catId)
	if err != nil {
		h.logger.Warn("Invalid category ID:", id)
		return response.Error(c, fiber.StatusBadRequest, "Validation Error", "Invalid category ID")
	}

	category, err := h.CategoryService.Destroy(uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "Failed delete Category!", err.Error())
	}

	return response.Success(c, "Category deleted successfully", category)
}
