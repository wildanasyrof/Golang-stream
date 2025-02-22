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

type EpisodeHandler struct {
	episodeService service.EpisodeService
	validation     validation.Validator
	logger         *zap.SugaredLogger
}

func NewEpisodeHandler(episodeService service.EpisodeService, logger *zap.SugaredLogger, validation validation.Validator) *EpisodeHandler {
	return &EpisodeHandler{episodeService: episodeService, logger: logger, validation: validation}
}

func (h *EpisodeHandler) Create(c *fiber.Ctx) error {
	animeId := c.Params("anime_id")
	id, err := strconv.Atoi(animeId)
	if err != nil {
		h.logger.Warn("Invalid Anime ID:", id)
		return response.Error(c, fiber.StatusBadRequest, "Validation Error", "Invalid Anime ID")
	}

	var req dto.CreateEpisodeRequest

	if err := h.validation.ValidateBody(c, &req); err != nil {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Validation error", err)
	}

	episode, err := h.episodeService.CreateEpisode(uint(id), req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Create Episode Failed!", err.Error())
	}

	return response.Created(c, "Episode created successfully", episode)
}
