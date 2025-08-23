package entities

import "github.com/google/uuid"

type CheckoutSummary struct {
	UserID         uuid.UUID  `json:"user_id"`
	Items          []CartItem `json:"items"`
	DeliveryCharge float32    `json:"delivery_charge"`
	TotalPrice     float32    `json:"total_price"`
	GrandTotal     float32    `json:"grand_total"`
	TaxPercent     float32    `json:"tax_percent"`
}
