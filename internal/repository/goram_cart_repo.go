package repository

import (
	"fmt"
	"happy_backend/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GoramCartRepo struct {
	db *gorm.DB
}

func NewGoramCartRepo(db *gorm.DB) *GoramCartRepo {
	return &GoramCartRepo{db: db}
}

func (r *GoramCartRepo) GetUserCartRepo(userId string) (*entities.Cart, error) {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	var cart entities.Cart
	if err := r.db.Preload("Items").First(&cart, "user_id = ?", uid).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *GoramCartRepo) CreateUserCart(userId string) (*entities.Cart, error) {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}
	cart := &entities.Cart{UserID: uid}
	if err := r.db.Create(cart).Error; err != nil {
		return nil, fmt.Errorf("failed to create cart: %w", err)

	}
	return cart, nil
}

func (r *GoramCartRepo) AddCartItemRepo(userId string, cartItem *entities.CartItem) error {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}

	cartItem.CartID = uid
	cartItem.TotalAmount = cartItem.ProductPrice * float32(cartItem.Quantity)

	if err := r.db.Create(cartItem).Error; err != nil {
		return fmt.Errorf("failed to add item to cart: %w", err)
	}

	return nil
}
func (r *GoramCartRepo) UpdateCartItemRepo(itemId string, cartItem *entities.CartItem) (*entities.CartItem, error) {
	uid, err := uuid.Parse(itemId)
	if err != nil {
		return nil, fmt.Errorf("invalid item id: %w", err)
	}

	cartItem.ID = uid
	cartItem.TotalAmount = cartItem.ProductPrice * float32(cartItem.Quantity)

	if err := r.db.Model(&entities.CartItem{}).
		Where("id = ?", uid).
		Updates(cartItem).Error; err != nil {
		return nil, fmt.Errorf("failed to update cart item: %w", err)
	}

	var updatedItem entities.CartItem
	if err := r.db.First(&updatedItem, "id = ?", uid).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch updated cart item: %w", err)
	}

	return &updatedItem, nil
}

func (r *GoramCartRepo) DeleteCartItemRepo(itemId string) error {
	uid, err := uuid.Parse(itemId)
	if err != nil {
		return fmt.Errorf("invalid item id: %w", err)
	}

	if err := r.db.Delete(&entities.CartItem{}, "id = ?", uid).Error; err != nil {
		return fmt.Errorf("failed to delete cart item: %w", err)
	}

	return nil
}
func (r *GoramCartRepo) GetCartItemByID(itemId string) (*entities.CartItem, error) {
	uid, err := uuid.Parse(itemId)
	if err != nil {
		return nil, fmt.Errorf("invalid item id: %w", err)
	}

	var item entities.CartItem
	if err := r.db.First(&item, "id = ?", uid).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch cart item: %w", err)
	}

	return &item, nil
}
