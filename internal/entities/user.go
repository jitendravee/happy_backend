package entities

type User struct {
	ID       string `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
