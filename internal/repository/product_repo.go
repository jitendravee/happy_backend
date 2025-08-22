package repository

import (
	"context"
	"fmt"
	"happy_backend/internal/entities"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepo struct {
	coll *mongo.Collection
}

func NewMongoProductRepo(db *mongo.Database) *MongoProductRepo {
	return &MongoProductRepo{coll: db.Collection("products")}
}
func (r *MongoProductRepo) GetByID(id string) (*entities.Product, error) {
	productUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid product id: %w", err)
	}

	var product entities.Product
	err = r.coll.FindOne(context.Background(), bson.M{"id": productUUID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch product by id: %w", err)
	}

	return &product, nil
}

// func(r *MongoProductRepo)
func (r *MongoProductRepo) Create(product *entities.Product) error {
	product.ID = uuid.New()

	// Generate UUIDs for each color
	for i := range product.Colors {
		product.Colors[i].ID = uuid.New()
	}

	_, err := r.coll.InsertOne(context.Background(), product)
	if err != nil {
		return fmt.Errorf("failed to insert product: %w", err)
	}
	return nil
}

func (r *MongoProductRepo) GetByName(productName string) (*entities.Product, error) {
	var product entities.Product
	err := r.coll.FindOne(context.Background(), bson.M{"name": productName}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch product by name: %w", err)
	}
	return &product, nil
}
