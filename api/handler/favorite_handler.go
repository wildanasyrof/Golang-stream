package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/golang-stream/internal/dto"
	"github.com/wildanasyrof/golang-stream/internal/service"
	"github.com/wildanasyrof/golang-stream/pkg/response"
	"github.com/wildanasyrof/golang-stream/pkg/validation"
	"go.uber.org/zap"
)

type FavoriteHandler struct {
	favoriteService service.FavoriteService
	validation      validation.Validator
	logger          *zap.SugaredLogger
}

func NewFavoriteHandler(favoriteService service.FavoriteService, validation validation.Validator, logger *zap.SugaredLogger) *FavoriteHandler {
	return &FavoriteHandler{
		favoriteService: favoriteService, validation: validation, logger: logger,
	}
}

func (h *FavoriteHandler) Create(c *fiber.Ctx) error {
	userID := uint(c.Locals("userID").(float64))
	var req dto.FavoriteRequest

	if err := h.validation.ValidateBody(c, &req); err != nil {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Validation error", err)
	}

	err := h.favoriteService.AddFavorite(userID, req.AnimeID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to add favorite", err.Error())
	}

	return response.Success(c, "Success add anime to favorite", nil)
}

func (h *FavoriteHandler) GetAll(c *fiber.Ctx) error {
	userID := uint(c.Locals("userID").(float64))

	limit := c.QueryInt("limit", 10)  // Default limit = 10
	offset := c.QueryInt("offset", 0) // Default offset = 0

	favorites, total, err := h.favoriteService.GetUserFavorites(userID, limit, offset)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to get Favorites", err.Error())
	}

	return response.Success(c, "Success get User Favorites", fiber.Map{
		"total":     total,
		"page":      offset + 1,
		"favorites": favorites,
	})
}

func (h *FavoriteHandler) Delete(c *fiber.Ctx) error {
	userID := uint(c.Locals("userID").(float64))

	var req dto.FavoriteRequest

	if err := h.validation.ValidateBody(c, &req); err != nil {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Validation error", err)
	}

	err := h.favoriteService.RemoveFavorite(userID, req.AnimeID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to remove Favorite", err.Error())
	}

	return response.Success(c, "Succes remove User Favorite", nil)
}
