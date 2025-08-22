package usecase

import (
	"errors"
	"happy_backend/internal/entities"
	"happy_backend/internal/repository"
)

type CommonColorUseCase struct {
	repo repository.CommonColorRepository
}

func NewCommonColorUseCase(r repository.CommonColorRepository) *CommonColorUseCase {
	return &CommonColorUseCase{repo: r}
}

func (u *CommonColorUseCase) GetAllTredingColorsUseCase() (*[]entities.CommonColor, error) {
	colors, err := u.repo.GetCommonColors()
	if err != nil {
		return nil, err
	}
	if colors == nil {
		return nil, errors.New("Common colors list not found")
	}
	return colors, nil
}
func (u *CommonColorUseCase) AddCommonColorUseCase(colorsData *[]entities.CommonColor) (*[]entities.CommonColor, error) {
	err := u.repo.AddCommonColors(colorsData)
	if err != nil {
		return nil, err
	}
	return colorsData, nil

}
func (u *CommonColorUseCase) UpdateCommonColorUseCase(id string, color *entities.CommonColor) (*entities.CommonColor, error) {
	updatedColor, err := u.repo.UpdateCommonColor(id, color)
	if err != nil {
		return nil, err
	}
	return updatedColor, nil
}
func (u *CommonColorUseCase) DeleteCommonColorByIDUseCase(id string) error {
	err := u.repo.DeleteCommonColorByID(id)
	if err != nil {
		return err
	}
	return nil
}
