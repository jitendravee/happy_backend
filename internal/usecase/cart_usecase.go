package usecase

import (
	"errors"
	"fmt"
	"happy_backend/internal/entities"
	"happy_backend/internal/repository"

	"gorm.io/gorm"
)

type CartUseCase struct {
	repo repository.CartRepository
}

func NewCartUseCase(repo repository.CartRepository) *CartUseCase {
	return &CartUseCase{
		repo: repo,
	}
}

func (uc *CartUseCase) GetCartDetailsUseCase(userID string) (*entities.Cart, error) {

	cart, err := uc.repo.GetUserCartRepo(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cart, err = uc.repo.CreateUserCart(userID)
			if err != nil {
				return nil, fmt.Errorf("failed to create cart: %w", err)
			}
			return cart, nil
		}
		return nil, fmt.Errorf("failed to fetch cart: %w", err)
	}

	return cart, nil
}

func (uc *CartUseCase) AddItemToCartUseCase(userId string, cartItem *entities.CartItem) error {
	err := uc.repo.AddCartItemRepo(userId, cartItem)
	if err != nil {
		return err
	}
	return nil
}
func (uc *CartUseCase) UpdateTheCartItemUseCase(itemId string, cartItem *entities.CartItem) (*entities.CartItem, error) {
	updatedItem, err := uc.repo.UpdateCartItemRepo(itemId, cartItem)
	if err != nil {
		return nil, err
	}
	return updatedItem, nil
}
func (uc *CartUseCase) GetCartItemById(itemId string) (*entities.CartItem, error) {
	cartItem, err := uc.repo.GetCartItemByID(itemId)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}
func (uc *CartUseCase) DeleteCartItemById(itemId string) error {
	err := uc.repo.DeleteCartItemRepo(itemId)
	if err != nil {
		return err
	}
	return nil
}
