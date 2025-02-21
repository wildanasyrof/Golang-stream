package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/golang-stream/api/routes"
	"github.com/wildanasyrof/golang-stream/config"
	"github.com/wildanasyrof/golang-stream/internal/module"
	"github.com/wildanasyrof/golang-stream/pkg/auth"
	"github.com/wildanasyrof/golang-stream/pkg/logger"
	"github.com/wildanasyrof/golang-stream/pkg/validation"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewFiberApp() *fiber.App {
	return fiber.New()
}

func StartServer(l *zap.SugaredLogger, app *fiber.App, db *gorm.DB) {
	port := config.GetEnv("PORT", "3000")
	l.Info("Server running on port:", port)
	log.Fatal(app.Listen(":" + port))

	app.Listen(":" + port)
}

var FiberModule = fx.Module("fiber", fx.Provide(NewFiberApp))

func main() {
	app := fx.New(
		config.Module,
		config.DatabaseModule,
		config.MigrationModule,
		logger.Module,
		auth.Module,
		routes.Module,
		validation.Module,
		module.UserModule,
		module.CategoryModule,
		module.AnimeModule,
		FiberModule,
		fx.Invoke(StartServer),
	)

	app.Run()
}
