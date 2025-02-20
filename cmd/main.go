package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/fiber-rest-api/api/routes"
	"github.com/wildanasyrof/fiber-rest-api/config"
	"github.com/wildanasyrof/fiber-rest-api/pkg/logger"
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
		FiberModule,
		logger.Module,
		routes.Module,
		fx.Invoke(StartServer),
	)

	app.Run()
}
