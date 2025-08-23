package repository

// import (
// 	"context"
// 	"fmt"
// 	"happy_backend/internal/entities"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type MongoUserRepo struct {
// 	coll *mongo.Collection
// }

// func NewMongoUserRepo(db *mongo.Database) *MongoUserRepo {
// 	return &MongoUserRepo{coll: db.Collection("users")}
// }

// func (r *MongoUserRepo) Create(user *entities.User) error {
// 	// Insert
// 	res, err := r.coll.InsertOne(context.Background(), bson.M{
// 		"name":     user.Name,
// 		"email":    user.Email,
// 		"password": user.Password,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("failed to insert user: %w", err)
// 	}

// 	// Convert inserted ID to string
// 	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
// 		user.ID = oid.Hex()
// 	}

// 	return nil
// }

// func (r *MongoUserRepo) GetByEmail(email string) (*entities.User, error) {
// 	var doc struct {
// 		ID       primitive.ObjectID `bson:"_id"`
// 		Name     string             `bson:"name"`
// 		Email    string             `bson:"email"`
// 		Password string             `bson:"password"`
// 	}

// 	err := r.coll.FindOne(context.Background(), bson.M{"email": email}).Decode(&doc)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, nil
// 		}
// 		return nil, fmt.Errorf("failed to fetch user by email: %w", err)
// 	}

// 	// Convert to entity
// 	return &entities.User{
// 		ID:       doc.ID.Hex(),
// 		Name:     doc.Name,
// 		Email:    doc.Email,
// 		Password: doc.Password,
// 	}, nil
// }
// func (r *MongoUserRepo) GetByID(id string) (*entities.User, error) {
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid id format: %w", err)
// 	}

// 	var doc struct {
// 		ID       primitive.ObjectID `bson:"_id"`
// 		Name     string             `bson:"name"`
// 		Email    string             `bson:"email"`
// 		Password string             `bson:"password"`
// 	}

// 	err = r.coll.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&doc)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, nil
// 		}
// 		return nil, fmt.Errorf("failed to fetch user by id: %w", err)
// 	}

// 	return &entities.User{
// 		ID:       doc.ID.Hex(),
// 		Name:     doc.Name,
// 		Email:    doc.Email,
// 		Password: doc.Password,
// 	}, nil
// }
