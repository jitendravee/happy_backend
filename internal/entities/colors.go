package entities

import "github.com/google/uuid"

type Color struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProductID uuid.UUID `json:"-" gorm:"type:uuid;index"`
	Name      string    `json:"name" gorm:"not null"`
	HexCode   string    `json:"hex_code"`
	Stock     int       `json:"stock" gorm:"default:0"`
}
