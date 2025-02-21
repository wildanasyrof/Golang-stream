package module

import (
	"github.com/wildanasyrof/golang-stream/api/handler"
	"github.com/wildanasyrof/golang-stream/internal/repository"
	"github.com/wildanasyrof/golang-stream/internal/service"
	"github.com/wildanasyrof/golang-stream/pkg/validation"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var AnimeModule = fx.Module("anime",
	fx.Provide(
		repository.NewAnimeRepository,
		service.NewAnimeService,
		func(animeService service.AnimeService, logger *zap.SugaredLogger, validation validation.Validator) *handler.AnimeHandler {
			return handler.NewAnimeHandler(animeService, logger, validation)
		},
	),
)
