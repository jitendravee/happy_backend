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

	postgresDB, err := db.NewPostgresDatabase(cfg)
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
		Cart:          repository.NewGoramCartRepo(postgresDB),
		Address:       repository.NewGoramAddressRepo(postgresDB),
		Checkout:      repository.NewCheckoutRepo(postgresDB),
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
		Cart:          usecase.NewCartUseCase(repos.Cart),
		Address:       usecase.NewAddressUseCase(repos.Address),
		Checkout:      usecase.NewCheckoutUseCase(repos.Checkout),
	}

	// Start HTTP server
	r := httpDelivery.NewServer(cfg, ucs)
	log.Println("🚀 Server running on :" + cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
