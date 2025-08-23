package usecase

import (
	"fmt"
	"happy_backend/internal/entities"
	"happy_backend/internal/repository"
)

type CheckoutUseCase struct {
	repo repository.CheckoutRepository
}

func NewCheckoutUseCase(repo repository.CheckoutRepository) *CheckoutUseCase {
	return &CheckoutUseCase{
		repo: repo,
	}
}

func (uc *CheckoutUseCase) GetCheckoutSummaryUseCase(
	userID string,
	items *[]entities.CartItem,
	deliveryCharge float32,
	taxPercent float32,
) (*entities.CheckoutSummary, error) {

	summary, err := uc.repo.GetCheckoutSummary(userID, items, deliveryCharge, taxPercent)
	if err != nil {
		return nil, fmt.Errorf("failed to generate checkout summary: %w", err)
	}

	return summary, nil
}
