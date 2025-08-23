package repository

import "happy_backend/internal/entities"

type UserRepository interface {
	Create(user *entities.User) error
	GetByEmail(email string) (*entities.User, error)
	GetByID(id string) (*entities.User, error)
}
type ProductRepository interface {
	Create(product *entities.Product) error
	GetByName(productName string) (*entities.Product, error)
	GetByID(productID string) (*entities.Product, error)
	GetAllProducts() (*[]entities.Product, error)
	UpdateProductByID(id string, product *entities.Product) (*entities.Product, error)
	DeleteProductByID(id string) error
	AddColorsToProduct(productID string, colors *[]entities.Color) (*[]entities.Color, error)
	UpdateProductColor(id string, colorId string, color *entities.Color) (*entities.Color, error)
}

type TrendingColorRepository interface {
	GetTrendingColors() (*[]entities.TrendingColor, error)
	AddTrendingColors(colors *[]entities.TrendingColor) error
	UpdateTrendingColor(colorID string, color *entities.TrendingColor) (*entities.TrendingColor, error)
	DeleteTrendingColorByID(colorID string) error
}
type CommonColorRepository interface {
	GetCommonColors() (*[]entities.CommonColor, error)
	AddCommonColors(colors *[]entities.CommonColor) error
	UpdateCommonColor(colorID string, color *entities.CommonColor) (*entities.CommonColor, error)
	DeleteCommonColorByID(colorID string) error
}
type CartRepository interface {
	GetUserCartRepo(userId string) (*entities.Cart, error)
	CreateUserCart(userId string) (*entities.Cart, error)
	AddCartItemRepo(userId string, cartItem *entities.CartItem) error
	UpdateCartItemRepo(itemId string, cartItem *entities.CartItem) (*entities.CartItem, error)
	DeleteCartItemRepo(itemId string) error
	GetCartItemByID(itemId string) (*entities.CartItem, error)
}
type AddressRepository interface {
	CreateAddress(userID string, address *entities.Address) (*entities.Address, error)
	GetAllAddresses(userID string) ([]*entities.Address, error)
	GetAddressByID(userID, addressID string) (*entities.Address, error)
	UpdateAddress(userID, addressID string, updated *entities.Address) (*entities.Address, error)
	DeleteAddress(userID, addressID string) error
}
type CheckoutRepository interface {
	GetCheckoutSummary(userID string, items *[]entities.CartItem, deliveryCharge, taxPercent float32) (*entities.CheckoutSummary, error)
}
