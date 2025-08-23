package usecase

type Usecases struct {
	User          *UserUsecase
	Product       *ProductUseCase
	TrendingColor *TrendingColorUseCase
	CommonColor   *CommonColorUseCase
	Cart          *CartUseCase
	Address       *AddressUseCase
	Checkout      *CheckoutUseCase
}
