package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/golang-stream/pkg/auth"
	"github.com/wildanasyrof/golang-stream/pkg/response"
	"go.uber.org/zap"
)

// AuthMiddleware protects routes with JWT authentication
func AuthMiddleware(logger *zap.SugaredLogger, auth *auth.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		// Check if the token exists
		if authHeader == "" {
			logger.Warn("Missing Authorization header")
			return response.Error(c, fiber.StatusUnauthorized, "Unauthorized")
		}

		// Extract token from "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			logger.Warn("Invalid Authorization header format")
			return response.Error(c, fiber.StatusUnauthorized, "Invalid token format")
		}

		tokenString := tokenParts[1]

		// Parse and validate JWT using auth package
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			logger.Warn("Invalid or expired token")
			return response.Error(c, fiber.StatusUnauthorized, "Invalid or expired token")
		}

		// Attach user ID and role to the request context
		c.Locals("userID", claims["user_id"])
		c.Locals("role", claims["role"])

		logger.Info("Authenticated request", "userID", claims["user_id"])
		return c.Next()
	}
}
