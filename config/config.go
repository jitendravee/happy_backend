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
	// Try multiple possible locations
	paths := []string{".env", "config/.env"}

	var err error
	for _, path := range paths {
		if _, statErr := os.Stat(path); statErr == nil {
			err = godotenv.Load(path)
			if err == nil {
				log.Printf("✅ Loaded env file from %s\n", path)
				break
			}
		}
	}

	if err != nil {
		log.Println("⚠️ No .env file found, falling back to system environment variables")
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "ep-mute-violet-adslhi5i-pooler.c-2.us-east-1.aws.neon.tech"),
		DBUser:     getEnv("DB_USER", "neondb_owner"),
		DBPassword: getEnv("DB_PASSWORD", "npg_dlV1wk4eMRWa"),
		DBName:     getEnv("DB_NAME", "neondb"),
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
