package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/fiber-rest-api/api/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RegisterRoutes sets up all routes
func RegisterRoutes(app *fiber.App, logger *zap.SugaredLogger) {
	app.Use(middleware.LoggerMiddleware(logger))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Fiber API is running with Uber Fx!")
	})
}

// Fx Module for Routes
var Module = fx.Module("routes", fx.Invoke(RegisterRoutes))
