package repository

import (
	"fmt"
	"happy_backend/internal/entities"

	"gorm.io/gorm"
)

type GormUserRepo struct {
	db *gorm.DB
}

func NewGormUserRepo(db *gorm.DB) *GormUserRepo {
	return &GormUserRepo{db: db}
}

// Create inserts a new user into the database
func (r *GormUserRepo) Create(user *entities.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

// GetByEmail fetches a user by email
func (r *GormUserRepo) GetByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch user by email: %w", err)
	}
	return &user, nil
}

// GetByID fetches a user by ID
func (r *GormUserRepo) GetByID(id string) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch user by id: %w", err)
	}
	return &user, nil
}
