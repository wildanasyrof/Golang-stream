package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// LoggerMiddleware logs incoming requests using Zap
func LoggerMiddleware(logger *zap.SugaredLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next() // Call the next handler

		logger.Infow("Incoming request",
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"latency", time.Since(start),
		)

		return err
	}
}
