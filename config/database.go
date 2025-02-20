package config

import (
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDatabase initializes a MySQL connection using GORM
func NewDatabase(logger *zap.SugaredLogger) (*gorm.DB, error) {
	dbUser := GetEnv("DB_USER", "root")
	dbPass := GetEnv("DB_PASS", "password")
	dbHost := GetEnv("DB_HOST", "127.0.0.1")
	dbPort := GetEnv("DB_PORT", "3306")
	dbName := GetEnv("DB_NAME", "mydatabase")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Info("❌ Failed to connect to MySQL: ", err)
		return nil, err
	}

	logger.Info("✅ Connected to MySQL database!")
	return db, nil
}

// Fx Module for Database
var DatabaseModule = fx.Module("database", fx.Provide(NewDatabase))
