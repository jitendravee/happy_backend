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

func (u *ProductUseCase) GetProductByID(productId string) (*entities.Product, error) {
	product, err := u.repo.GetByID(productId)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}
func (u *ProductUseCase) GetAllProductsUseCase() (*[]entities.Product, error) {
	products, err := u.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	if products == nil {
		return nil, errors.New("products list not found")
	}
	return products, nil
}
func (u *ProductUseCase) UpdateProductByID(id string, product *entities.Product) (*entities.Product, error) {
	product, err := u.repo.UpdateProductByID(id, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (u *ProductUseCase) DeleteProductByID(id string) error {
	err := u.repo.DeleteProductByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *ProductUseCase) AddProductColors(id string, colors *[]entities.Color) (*[]entities.Color, error) {
	addedColors, err := u.repo.AddColorsToProduct(id, colors)
	if err != nil {
		return nil, err
	}

	return addedColors, nil
}

func (u *ProductUseCase) UpdateProductColor(id string, colorId string, color *entities.Color) (*entities.Color, error) {
	updatedColor, err := u.repo.UpdateProductColor(id, colorId, color)
	if err != nil {
		return nil, err
	}
	return updatedColor, nil
}
