package repository

import (
	"fmt"
	"happy_backend/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GoramTrendingColorRepo struct {
	db *gorm.DB
}

func NewGoramTrendingColorRepo(db *gorm.DB) *GoramTrendingColorRepo {
	return &GoramTrendingColorRepo{
		db: db,
	}
}
func (r *GoramTrendingColorRepo) GetTrendingColors() (*[]entities.TrendingColor, error) {
	var colors []entities.TrendingColor
	if err := r.db.Find(&colors).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch trending colors: %w", err)

	}
	return &colors, nil
}

func (r *GoramTrendingColorRepo) AddTrendingColors(colors *[]entities.TrendingColor) error {
	for i := range *colors {
		newColor := &(*colors)[i]

		if newColor.ID == uuid.Nil {
			newColor.ID = uuid.New()
		}

		var existing entities.TrendingColor
		if err := r.db.First(&existing, "hex_code = ?", newColor.HexCode).Error; err == nil {
			return fmt.Errorf("color with hex code %s already exists", newColor.HexCode)
		} else if err != gorm.ErrRecordNotFound {
			return fmt.Errorf("failed to check existing color: %w", err)
		}
	}

	if err := r.db.Create(colors).Error; err != nil {
		return fmt.Errorf("failed to add trending colors: %w", err)
	}

	return nil
}

func (r *GoramTrendingColorRepo) UpdateTrendingColor(colorID string, color *entities.TrendingColor) (*entities.TrendingColor, error) {
	var existing entities.TrendingColor
	if err := r.db.First(&existing, "id = ?", colorID).Error; err != nil {
		return nil, fmt.Errorf("color not found: %w", err)
	}
	if err := r.db.Model(&existing).Updates(color).Error; err != nil {
		return nil, fmt.Errorf("failed to update trending color: %w", err)
	}
	return &existing, nil
}
func (r *GoramTrendingColorRepo) DeleteTrendingColorByID(colorID string) error {
	if err := r.db.Delete(&entities.TrendingColor{}, "id = ?", colorID).Error; err != nil {
		return fmt.Errorf("failed to delete trending color: %w", err)
	}
	return nil
}
