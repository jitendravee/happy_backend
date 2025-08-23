package entities

import "github.com/google/uuid"

type Cart struct {
	UserID uuid.UUID  `json:"user_id" gorm:"type:uuid;primaryKey"`
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID;references:UserID"`
}

type CartItem struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CartID       uuid.UUID `json:"cart_id" gorm:"type:uuid;not null"`
	ProductID    uuid.UUID `json:"product_id" gorm:"type:uuid;not null" binding:"required"`
	ProductName  string    `json:"product_name" gorm:"not null" binding:"required"`
	ColorCode    string    `json:"color_code" gorm:"not null" binding:"required"`
	ProductPrice float32   `json:"product_price" gorm:"not null" binding:"required"`
	Quantity     int       `json:"quantity" gorm:"not null;check:quantity >= 1" binding:"required,min=1"`
	TotalAmount  float32   `json:"total_amount" gorm:"not null;default:0"`
}
