package entities

type Color struct {
	ID      string `json:"id" bson:"-"`
	Name    string `json:"name" bson:"name"`
	HexCode string `json:"hex_code" bson:"hex_code"`
	Stock   int    `json:"stock" bson:"stock"`
}
