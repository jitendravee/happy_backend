package repository

type Repositories struct {
	User          UserRepository
	Product       ProductRepository
	TrendingColor TrendingColorRepository
	CommonColor   CommonColorRepository
	Cart          CartRepository
	Address       AddressRepository
	Checkout      CheckoutRepository
}
