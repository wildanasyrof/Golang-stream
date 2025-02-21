package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/golang-stream/pkg/auth"
	"github.com/wildanasyrof/golang-stream/pkg/response"
)

// AuthMiddleware protects routes with JWT authentication
func AuthMiddleware(authService *auth.AuthService, allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from Authorization header
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return response.Error(c, fiber.StatusUnauthorized, "Unauthorized", "Missing Authorization header")
		}

		// Remove "Bearer " prefix if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// Validate token
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			return response.Error(c, fiber.StatusUnauthorized, "Unauthorized", "Invalid or expired token")
		}

		// Extract role from claims
		role, ok := claims["role"].(string)
		if !ok {
			return response.Error(c, fiber.StatusForbidden, "Unauthorized", "No role found")
		}

		c.Locals("userID", claims["user_id"])
		c.Locals("role", claims["role"])

		// If no specific roles are required, allow all authenticated users
		if len(allowedRoles) == 0 {
			return c.Next()
		}

		// Check if user has an allowed role
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				return c.Next() // Proceed to the next handler
			}
		}

		// If role is not allowed, return forbidden
		return response.Error(c, fiber.StatusForbidden, "Unauthorized", "Forbidden")
	}
}
