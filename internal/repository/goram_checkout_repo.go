package repository

import (
	"fmt"
	"happy_backend/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckoutRepo struct {
	db *gorm.DB
}

func NewCheckoutRepo(db *gorm.DB) *CheckoutRepo {
	return &CheckoutRepo{db: db}
}

func (r *CheckoutRepo) GetCheckoutSummary(userID string, items *[]entities.CartItem, deliveryCharge, taxPercent float32) (*entities.CheckoutSummary, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	var totalPrice float32
	for i := range *items {
		(*items)[i].TotalAmount = (*items)[i].ProductPrice*float32((*items)[i].Quantity) +
			((*items)[i].ProductPrice * float32((*items)[i].Quantity) * taxPercent / 100)
		totalPrice += (*items)[i].TotalAmount
	}

	grandTotal := totalPrice + deliveryCharge

	return &entities.CheckoutSummary{
		UserID:         uid,
		Items:          *items,
		DeliveryCharge: deliveryCharge,
		TotalPrice:     totalPrice,
		GrandTotal:     grandTotal,
		TaxPercent:     taxPercent,
	}, nil
}
