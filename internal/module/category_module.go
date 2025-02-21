package module

import (
	"github.com/wildanasyrof/golang-stream/api/handler"
	"github.com/wildanasyrof/golang-stream/internal/repository"
	"github.com/wildanasyrof/golang-stream/internal/service"
	"github.com/wildanasyrof/golang-stream/pkg/validation"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var CategoryModule = fx.Module("category",
	fx.Provide(
		repository.NewCategoryRepository,
		service.NewCategoryService,
		func(categoryService service.CategoryService, logger *zap.SugaredLogger, validation validation.Validator) *handler.CategoryHandler {
			return handler.NewCategoryHandler(categoryService, logger, validation)
		},
	),
)
