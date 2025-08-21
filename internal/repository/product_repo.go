package repository

import (
	"context"
	"fmt"
	"happy_backend/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepo struct {
	coll *mongo.Collection
}

func NewMongoProductRepo(db *mongo.Database) *MongoProductRepo {
	return &MongoProductRepo{coll: db.Collection("products")}
}
func (r *MongoProductRepo) Create(product *entities.Product) error {
	productOID := primitive.NewObjectID()
	product.ID = productOID.Hex()

	for i := range product.Colors {
		colorOID := primitive.NewObjectID()
		product.Colors[i].ID = colorOID.Hex()
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
