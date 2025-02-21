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
	animes, err := h.animeService.GetAllAnime()
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Get All Anime Failed!", err.Error())
	}

	return response.Success(c, "Success get all anime", animes)
}
