package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/golang-stream/api/handler"
	"github.com/wildanasyrof/golang-stream/api/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RegisterRoutes sets up all routes
func RegisterRoutes(app *fiber.App, logger *zap.SugaredLogger, userHandler *handler.UserHandler) {
	app.Use(middleware.LoggerMiddleware(logger))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Fiber API is running with Uber Fx!")
	})

	app.Post("/register", userHandler.RegisterUser)
}

// Fx Module for Routes
var Module = fx.Module("routes", fx.Invoke(RegisterRoutes))
