package repository

import (
	"fmt"
	"happy_backend/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormProductRepo struct {
	db *gorm.DB
}

func NewGormProductRepo(db *gorm.DB) *GormProductRepo {
	return &GormProductRepo{db: db}
}

// Create inserts a new product
func (r *GormProductRepo) Create(product *entities.Product) error {
	// GORM will auto-generate IDs because we used gorm.Model/primary keys in entities
	if err := r.db.Create(product).Error; err != nil {
		return fmt.Errorf("failed to insert product: %w", err)
	}
	return nil
}

// GetByID finds a product by its ID
func (r *GormProductRepo) GetByID(id string) (*entities.Product, error) {
	var product entities.Product
	if err := r.db.Preload("Colors").Preload("Composition").
		First(&product, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch product by id: %w", err)
	}
	return &product, nil
}

// GetByName finds a product by its name
func (r *GormProductRepo) GetByName(productName string) (*entities.Product, error) {
	var product entities.Product
	if err := r.db.Preload("Colors").Preload("Composition").
		First(&product, "name = ?", productName).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch product by name: %w", err)
	}
	return &product, nil
}

// get all the products from the table
func (r *GormProductRepo) GetAllProducts() (*[]entities.Product, error) {
	var products []entities.Product
	if err := r.db.Preload("Colors").Preload("Composition").Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch all products: %w", err)

	}
	return &products, nil
}

func (r *GormProductRepo) UpdateProductByID(id string, product *entities.Product) (*entities.Product, error) {
	if err := r.db.Model(&entities.Product{}).Where("id = ?", id).Updates(product).Error; err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}
	var updatedProduct entities.Product
	if err := r.db.Preload("Colors").Preload("Composition").First(&updatedProduct, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch updated product: %w", err)
	}

	return &updatedProduct, nil
}

func (r *GormProductRepo) DeleteProductByID(id string) error {
	if err := r.db.Delete(&entities.Product{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

func (r *GormProductRepo) UpdateProductColor(idStr string, colorIdStr string, color *entities.Color) (*entities.Color, error) {
	productID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid product id: %w", err)
	}

	colorID, err := uuid.Parse(colorIdStr)
	if err != nil {
		return nil, fmt.Errorf("invalid color id: %w", err)
	}

	var existingColor entities.Color
	if err := r.db.First(&existingColor, "id = ?", colorID).Error; err != nil {
		return nil, fmt.Errorf("color not found: %w", err)
	}

	var otherColors []entities.Color
	if err := r.db.Model(&entities.Product{ID: productID}).Association("Colors").Find(&otherColors); err == nil {
		for _, c := range otherColors {
			if c.HexCode == color.HexCode && c.ID != colorID {
				return nil, fmt.Errorf("color with hex code %s already exists", color.HexCode)
			}
		}
	}

	if err := r.db.Model(&existingColor).Updates(color).Error; err != nil {
		return nil, fmt.Errorf("failed to update color: %w", err)
	}

	return &existingColor, nil
}
func (r *GormProductRepo) AddColorsToProduct(productIDStr string, colors *[]entities.Color) (*[]entities.Color, error) {
	// Parse ProductID string to uuid.UUID
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid product id: %w", err)
	}

	// Fetch existing colors
	var existingColors []entities.Color
	if err := r.db.Model(&entities.Product{ID: productID}).Association("Colors").Find(&existingColors); err != nil {
		return nil, fmt.Errorf("failed to fetch existing colors: %w", err)
	}

	var addedColors []entities.Color
	for i := range *colors {
		newColor := &(*colors)[i]

		exists := false
		for _, c := range existingColors {
			if c.HexCode == newColor.HexCode {
				exists = true
				break
			}
		}
		if exists {
			return nil, fmt.Errorf("color with hex code %s already exists", newColor.HexCode)
		}

		if newColor.ID == uuid.Nil {
			newColor.ID = uuid.New()
		}

		newColor.ProductID = productID

		if err := r.db.Create(newColor).Error; err != nil {
			return nil, fmt.Errorf("failed to create color: %w", err)
		}

		if err := r.db.Model(&entities.Product{ID: productID}).Association("Colors").Append(newColor); err != nil {
			return nil, fmt.Errorf("failed to associate color: %w", err)
		}

		addedColors = append(addedColors, *newColor)
	}

	return &addedColors, nil
}
