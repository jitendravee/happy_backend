package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Postgres DB config
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string

	// App config
	Port             string
	JWTSecret        string
	JWTRefreshSecret string
	JWTExpiry        string
	JWTRefreshExpiry string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env") // load from project root
	wd, _ := os.Getwd()
	log.Println("Working directory:", wd)

	if err != nil {
		err = godotenv.Load("config/.env") // fallback
	}

	if err != nil {
		log.Println("⚠️ No .env file found, falling back to system environment variables")
	} else {
		log.Println("✅ Loaded env file successfully")
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", ""),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		DBPort:     getEnv("DB_PORT", "5432"),

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
