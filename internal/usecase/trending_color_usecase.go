package usecase

import (
	"errors"
	"happy_backend/internal/entities"
	"happy_backend/internal/repository"
)

type TrendingColorUseCase struct {
	repo repository.TrendingColorRepository
}

func NewTrendingColorUseCase(r repository.TrendingColorRepository) *TrendingColorUseCase {
	return &TrendingColorUseCase{repo: r}
}

func (u *TrendingColorUseCase) GetAllTredingColorsUseCase() (*[]entities.TrendingColor, error) {
	colors, err := u.repo.GetTrendingColors()
	if err != nil {
		return nil, err
	}
	if colors == nil {
		return nil, errors.New("trending colors list not found")
	}
	return colors, nil
}
func (u *TrendingColorUseCase) AddTrendingColorUseCase(colorsData *[]entities.TrendingColor) (*[]entities.TrendingColor, error) {
	err := u.repo.AddTrendingColors(colorsData)
	if err != nil {
		return nil, err
	}
	return colorsData, nil

}
func (u *TrendingColorUseCase) UpdateTrendingColorUseCase(id string, color *entities.TrendingColor) (*entities.TrendingColor, error) {
	updatedColor, err := u.repo.UpdateTrendingColor(id, color)
	if err != nil {
		return nil, err
	}
	return updatedColor, nil
}
func (u *TrendingColorUseCase) DeleteTrendingColorByIDUseCase(id string) error {
	err := u.repo.DeleteTrendingColorByID(id)
	if err != nil {
		return err
	}
	return nil
}
