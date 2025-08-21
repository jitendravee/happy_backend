package usecase

import (
	"errors"
	"happy_backend/internal/entities"
	"happy_backend/internal/repository"
)

type ProductUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(r repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: r}
}
func (u *ProductUseCase) AddProduct(product *entities.Product) (*entities.Product, error) {
	existingProduct, _ := u.repo.GetByName(product.Name)
	if existingProduct != nil {
		return existingProduct, errors.New("Product already exists")
	}
	if err := u.repo.Create(product); err != nil {
		return nil, err
	}
	return product, nil
}
