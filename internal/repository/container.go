package repository

type Repositories struct {
	User          UserRepository
	Product       ProductRepository
	TrendingColor TrendingColorRepository
	CommonColor   CommonColorRepository
}
