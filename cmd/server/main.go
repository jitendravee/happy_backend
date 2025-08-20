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
	cfg := config.LoadConfig()

	mongoDB := db.NewMongoDatabase(cfg)

	repos := &repository.Repositories{
		User: repository.NewMongoUserRepo(mongoDB),
	}
	jwtExpiry, _ := time.ParseDuration(cfg.JWTExpiry)
	jwtRefreshExpiry, _ := time.ParseDuration(cfg.JWTRefreshExpiry)
	ucs := &usecase.Usecases{
		User: usecase.NewUserUsecase(repos.User, cfg.JWTSecret,
			cfg.JWTRefreshSecret,
			jwtExpiry,
			jwtRefreshExpiry),
	}

	r := httpDelivery.NewServer(cfg, ucs)
	log.Println("🚀 Server running on :" + cfg.Port)
	r.Run(":" + cfg.Port)
}
