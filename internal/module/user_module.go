package module

import (
	"github.com/wildanasyrof/golang-stream/api/handler"
	"github.com/wildanasyrof/golang-stream/internal/repository"
	"github.com/wildanasyrof/golang-stream/internal/service"
	"github.com/wildanasyrof/golang-stream/pkg/auth"
	"github.com/wildanasyrof/golang-stream/pkg/validation"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// UserModule combines repository, usecase, and handler in one module
var UserModule = fx.Module("user",
	fx.Provide(
		repository.NewUserRepository,
		service.NewUserService,
		func(userService service.UserService, validation validation.Validator, auth *auth.AuthService, logger *zap.SugaredLogger) *handler.UserHandler {
			return handler.NewUserHandler(validation, userService, auth, logger)
		},
	),
)
