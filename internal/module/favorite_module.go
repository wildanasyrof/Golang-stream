package module

import (
	"github.com/wildanasyrof/golang-stream/api/handler"
	"github.com/wildanasyrof/golang-stream/internal/repository"
	"github.com/wildanasyrof/golang-stream/internal/service"
	"github.com/wildanasyrof/golang-stream/pkg/validation"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var FavoriteModule = fx.Module("favorite",
	fx.Provide(
		repository.NewFavoriteRepository,
		service.NewFavoriteService,
		func(favoriteService service.FavoriteService, validation validation.Validator, logger *zap.SugaredLogger) *handler.FavoriteHandler {
			return handler.NewFavoriteHandler(favoriteService, validation, logger)
		},
	),
)
