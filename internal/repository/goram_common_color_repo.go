package repository

import (
	"fmt"
	"happy_backend/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GoramCommonColorRepo struct {
	db *gorm.DB
}

func NewGoramCommonColorRepo(db *gorm.DB) *GoramCommonColorRepo {
	return &GoramCommonColorRepo{
		db: db,
	}
}
func (r *GoramCommonColorRepo) GetCommonColors() (*[]entities.CommonColor, error) {
	var colors []entities.CommonColor
	if err := r.db.Find(&colors).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch Common colors: %w", err)

	}
	return &colors, nil
}

func (r *GoramCommonColorRepo) AddCommonColors(colors *[]entities.CommonColor) error {
	for i := range *colors {
		newColor := &(*colors)[i]

		if newColor.ID == uuid.Nil {
			newColor.ID = uuid.New()
		}

		var existing entities.CommonColor
		if err := r.db.First(&existing, "hex_code = ?", newColor.HexCode).Error; err == nil {
			return fmt.Errorf("color with hex code %s already exists", newColor.HexCode)
		} else if err != gorm.ErrRecordNotFound {
			return fmt.Errorf("failed to check existing color: %w", err)
		}
	}

	if err := r.db.Create(colors).Error; err != nil {
		return fmt.Errorf("failed to add Common colors: %w", err)
	}

	return nil
}

func (r *GoramCommonColorRepo) UpdateCommonColor(colorID string, color *entities.CommonColor) (*entities.CommonColor, error) {
	var existing entities.CommonColor
	if err := r.db.First(&existing, "id = ?", colorID).Error; err != nil {
		return nil, fmt.Errorf("color not found: %w", err)
	}
	if err := r.db.Model(&existing).Updates(color).Error; err != nil {
		return nil, fmt.Errorf("failed to update Common color: %w", err)
	}
	return &existing, nil
}
func (r *GoramCommonColorRepo) DeleteCommonColorByID(colorID string) error {
	if err := r.db.Delete(&entities.CommonColor{}, "id = ?", colorID).Error; err != nil {
		return fmt.Errorf("failed to delete Common color: %w", err)
	}
	return nil
}
