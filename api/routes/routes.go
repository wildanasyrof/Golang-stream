package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/golang-stream/api/handler"
	"github.com/wildanasyrof/golang-stream/api/middleware"
	"github.com/wildanasyrof/golang-stream/pkg/auth"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RegisterRoutes sets up all API routes
func RegisterRoutes(app *fiber.App, logger *zap.SugaredLogger, authService *auth.AuthService, userHandler *handler.UserHandler, categoryHandler *handler.CategoryHandler, animeHandler *handler.AnimeHandler, episodeHandler *handler.EpisodeHandler) {
	// Global Middleware
	app.Use(middleware.LoggerMiddleware(logger))

	// Health Check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Fiber API is running with Uber Fx!")
	})

	// Auth Routes
	authRoutes := app.Group("/auth")
	authRoutes.Post("/register", userHandler.RegisterUser)
	authRoutes.Post("/login", userHandler.Login)

	// User Routes (Requires Authentication)
	userRoutes := app.Group("/user", middleware.AuthMiddleware(authService))
	userRoutes.Get("/", userHandler.GetProfile)
	userRoutes.Put("/", userHandler.UpdateProfile)

	// Public Category Routes
	categoryRoutes := app.Group("/category")
	categoryRoutes.Get("/", categoryHandler.GetAll) // Public

	animeRoutes := app.Group("/anime")
	animeRoutes.Get("/", animeHandler.GetAll)
	animeRoutes.Get("/:id", animeHandler.GetByID)

	episodeRoutes := animeRoutes.Group("/:anime_id/episode")
	episodeRoutes.Get("/", episodeHandler.Get)

	// Admin-Only Category Routes
	adminCategoryRoutes := categoryRoutes.Group("/", middleware.AuthMiddleware(authService, "ADMIN"))
	adminCategoryRoutes.Post("/", categoryHandler.Create)
	adminCategoryRoutes.Delete("/:id", categoryHandler.Destroy)

	adminAnimeRoutes := animeRoutes.Group("/", middleware.AuthMiddleware(authService, "ADMIN"))
	adminAnimeRoutes.Post("/", animeHandler.Create)
	adminAnimeRoutes.Put("/:id", animeHandler.Update)
	adminAnimeRoutes.Delete("/:id", animeHandler.Delete)

	adminEpisodeRoutes := episodeRoutes.Group("/", middleware.AuthMiddleware(authService, "ADMIN"))
	adminEpisodeRoutes.Post("/", episodeHandler.Create)
	adminEpisodeRoutes.Put("/:eps_number", episodeHandler.Update)
	adminEpisodeRoutes.Delete("/:eps_number", episodeHandler.Delete)
}

// Fx Module for Routes
var Module = fx.Module("routes", fx.Invoke(RegisterRoutes))
