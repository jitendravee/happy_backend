package entities

import "github.com/google/uuid"

type Cart struct {
	UserID uuid.UUID  `json:"user_id" gorm:"type:uuid;primaryKey"`
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID;references:UserID"`
}

type CartItem struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CartID    uuid.UUID `json:"cart_id" gorm:"type:uuid;not null"`
	ProductID string    `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null;check:quantity >= 1"`
}
