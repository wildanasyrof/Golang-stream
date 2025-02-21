package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/golang-stream/internal/dto"
	"github.com/wildanasyrof/golang-stream/internal/service"
	"github.com/wildanasyrof/golang-stream/pkg/response"
	"github.com/wildanasyrof/golang-stream/pkg/validation"
	"go.uber.org/zap"
)

type AnimeHandler struct {
	animeService service.AnimeService
	validation   validation.Validator
	logger       *zap.SugaredLogger
}

func NewAnimeHandler(animeService service.AnimeService, logger *zap.SugaredLogger, validation validation.Validator) *AnimeHandler {
	return &AnimeHandler{animeService: animeService, logger: logger, validation: validation}
}

func (h *AnimeHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateAnimeRequest

	if err := h.validation.ValidateBody(c, &req); err != nil {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Validation error", err)
	}

	anime, err := h.animeService.CreateAnime(req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Create Anime Failed!", err.Error())
	}

	return response.Created(c, "Anime created successfully", anime)
}

func (h *AnimeHandler) GetAll(c *fiber.Ctx) error {
	// Get query parameters
	limit := c.QueryInt("limit", 10) // Default to 10 per page
	page := c.QueryInt("page", 1)    // Default to page 1

	// Extract filtering parameters
	filters := make(map[string]string)
	if title := c.Query("title"); title != "" {
		filters["title"] = title
	}
	if year := c.Query("year"); year != "" {
		filters["year"] = year
	}
	if studio := c.Query("studio"); studio != "" {
		filters["studio"] = studio
	}
	if category := c.Query("category"); category != "" {
		filters["category"] = category
	}

	// Fetch anime with filters & pagination
	animes, total, err := h.animeService.GetAllAnime(limit, page, filters)
	if err != nil {
		h.logger.Error("Get all anime failed: ", err)
		return response.Error(c, fiber.StatusInternalServerError, "Get All Anime Failed!", err.Error())
	}

	// Return paginated response
	return response.Success(c, "Success get all anime", fiber.Map{
		"total": total,
		"page":  page,
		"limit": limit,
		"data":  animes,
	})
}
