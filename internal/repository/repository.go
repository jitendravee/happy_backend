package repository

import "happy_backend/internal/entities"

type UserRepository interface {
	Create(user *entities.User) error
	GetByEmail(email string) (*entities.User, error)
	GetByID(id string) (*entities.User, error)
}
type ProductRepository interface {
	Create(product *entities.Product) error
	GetByName(productName string) (*entities.Product, error)
}
