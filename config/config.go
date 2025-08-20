package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURI            string
	DBName           string
	Port             string
	JWTSecret        string
	JWTRefreshSecret string
	JWTExpiry        string
	JWTRefreshExpiry string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, falling back to system environment variables")
	}

	return &Config{
		DBURI:            getEnv("MONGO_URI", ""),
		DBName:           getEnv("MONGO_DBNAME", "happydb"),
		Port:             getEnv("PORT", "8080"),
		JWTSecret:        getEnv("JWT_SECRET", "defaultsecret"),
		JWTRefreshSecret: getEnv("JWT_REFRESH_SECRET", "defaultrefresh"),
		JWTExpiry:        getEnv("JWT_EXPIRY", "15m"),
		JWTRefreshExpiry: getEnv("JWT_REFRESH_EXPIRY", "168h"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
