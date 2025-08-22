package db

import (
	"fmt"
	"happy_backend/config"
	"happy_backend/internal/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresDatabase creates a GORM DB connection for Postgres
func NewPostgresDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres: %w", err)
	}

	return db, nil
}

// Migrate runs GORM auto migrations
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.User{},
		&entities.Product{},
		&entities.Color{},
		&entities.Composition{},
	)
}
