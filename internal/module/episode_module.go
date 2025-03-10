package module

import (
	"github.com/wildanasyrof/golang-stream/api/handler"
	"github.com/wildanasyrof/golang-stream/internal/repository"
	"github.com/wildanasyrof/golang-stream/internal/service"
	"github.com/wildanasyrof/golang-stream/pkg/validation"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var EpisodeModule = fx.Module("episode",
	fx.Provide(
		repository.NewEpisodeRepository,
		service.NewEpisodeService,
		func(episodeService service.EpisodeService, logger *zap.SugaredLogger, validation validation.Validator) *handler.EpisodeHandler {
			return handler.NewEpisodeHandler(episodeService, logger, validation)
		},
	),
)
