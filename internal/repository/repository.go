package repository

import "happy_backend/internal/entities"

type UserRepository interface {
	Create(user *entities.User) error
	GetByEmail(email string) (*entities.User, error)
	GetByID(id string) (*entities.User, error)
}
