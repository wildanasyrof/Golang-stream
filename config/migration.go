package config

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"

	"go.uber.org/fx"
)

// RunMigrations applies database migrations
func RunMigrations(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get SQL DB from GORM: %v", err)
	}

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		log.Fatalf("Migration driver error: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration", // ✅ Migration folder
		"mysql", driver,
	)
	if err != nil {
		log.Fatalf("Migration instance error: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("✅ Database migrated successfully!")
}

// Fx Module for Migrations
var MigrationModule = fx.Module("migration", fx.Invoke(RunMigrations))
