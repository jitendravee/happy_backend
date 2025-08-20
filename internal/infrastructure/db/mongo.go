package db

import (
	"context"
	"log"
	"time"

	"happy_backend/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoDatabase initializes a new MongoDB connection
func NewMongoDatabase(cfg *config.Config) *mongo.Database {
	clientOpts := options.Client().ApplyURI(cfg.DBURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("❌ Mongo connection failed:", err)
	}

	// Ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("❌ Mongo ping failed:", err)
	}

	log.Println("✅ Connected to MongoDB")

	return client.Database(cfg.DBName)
}
