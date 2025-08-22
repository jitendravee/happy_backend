// cmd/server/main.go
package main

import (
	"happy_backend/config"
	httpDelivery "happy_backend/internal/delivery/http"
	"happy_backend/internal/infrastructure/db"
	"happy_backend/internal/repository"
	"happy_backend/internal/usecase"
	"log"
	"time"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize Postgres (Neon) with GORM
	postgresDB, err := db.NewPostgresDatabase(cfg) // <- this replaces mongo
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Run migrations (auto-create tables for entities)
	if err := db.Migrate(postgresDB); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	// Initialize repositories
	repos := &repository.Repositories{
		Product:       repository.NewGormProductRepo(postgresDB),
		User:          repository.NewGormUserRepo(postgresDB),
		TrendingColor: repository.NewGoramTrendingColorRepo(postgresDB),
		CommonColor:   repository.NewGoramCommonColorRepo(postgresDB),
	}

	// Parse JWT expiry durations
	jwtExpiry, _ := time.ParseDuration(cfg.JWTExpiry)
	jwtRefreshExpiry, _ := time.ParseDuration(cfg.JWTRefreshExpiry)

	// Initialize use cases
	ucs := &usecase.Usecases{
		User: usecase.NewUserUsecase(
			repos.User,
			cfg.JWTSecret,
			cfg.JWTRefreshSecret,
			jwtExpiry,
			jwtRefreshExpiry,
		),
		Product:       usecase.NewProductUseCase(repos.Product),
		TrendingColor: usecase.NewTrendingColorUseCase(repos.TrendingColor),
		CommonColor:   usecase.NewCommonColorUseCase(repos.CommonColor),
	}

	// Start HTTP server
	r := httpDelivery.NewServer(cfg, ucs)
	log.Println("🚀 Server running on :" + cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
