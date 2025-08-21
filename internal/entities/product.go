package entities

type Product struct {
	ID          string        `json:"id" bson:"-"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Price       float32       `json:"price" bson:"price"`
	FabricType  string        `json:"fabric_type" bson:"fabric_type"`
	ImagesUrl   []string      `json:"images_url" bson:"images_url"`
	Colors      []Color       `json:"colors" bson:"colors"`
	Composition []Composition `json:"composition" bson:"composition"`
	Active      bool          `json:"active" bson:"active"`
}
type Composition struct {
	Label      string `json:"label" bson:"label"`
	Percentage string `json:"percentage" bson:"percentage"`
}
