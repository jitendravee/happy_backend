package entities

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Product struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string         `json:"name" gorm:"unique;not null"`
	Description string         `json:"description"`
	Price       float32        `json:"price"`
	FabricType  string         `json:"fabric_type"`
	ImagesUrl   pq.StringArray `json:"images_url" gorm:"type:text[]"`
	Colors      []Color        `json:"colors" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	Composition []Composition  `json:"composition" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	Active      bool           `json:"active" gorm:"default:true"`
}

type Composition struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProductID  uuid.UUID `json:"-" gorm:"type:uuid;index"`
	Label      string    `json:"label"`
	Percentage string    `json:"percentage"`
}
