package entities

import "github.com/google/uuid"

type Address struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	AddressBookID uuid.UUID `json:"-" gorm:"type:uuid;not null;index"` // maps to UserID of AddressBook
	Line1         string    `json:"line1" gorm:"not null"`
	Line2         string    `json:"line2"`
	City          string    `json:"city" gorm:"not null"`
	State         string    `json:"state" gorm:"not null"`
	ZipCode       string    `json:"zip_code" gorm:"not null"`
	Country       string    `json:"country" gorm:"not null"`
	IsDefault     bool      `json:"is_default" gorm:"default:false"`
}
