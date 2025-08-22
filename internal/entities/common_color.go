package entities

import "github.com/google/uuid"

type CommonColor struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name    string    `json:"name" gorm:"not null"`
	HexCode string    `json:"hex_code" gorm:"not null;unique"`
}
