package http

import (
	"happy_backend/config"
	"happy_backend/internal/usecase"
	"happy_backend/pkg/middleware"
	"happy_backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func NewServer(cfg *config.Config, ucs *usecase.Usecases) *gin.Engine {
	r := gin.Default()

	registerPublicRoutes(r, ucs)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(ucs.User.Secret()))
	registerProtectedRoutes(api, ucs)

	r.GET("/health", func(c *gin.Context) {
		response.Success(c, 200, "Service running", gin.H{"service": "ok"})
	})

	return r
}

// Public endpoints like signup, signin, refresh
func registerPublicRoutes(r *gin.Engine, ucs *usecase.Usecases) {
	userHandler := NewUserHandler(ucs.User)
	r.POST("/signup", userHandler.SignUp)
	r.POST("/signin", userHandler.SignIn)
	r.POST("/refresh", userHandler.Refresh)
}

// Protected endpoints
func registerProtectedRoutes(api *gin.RouterGroup, ucs *usecase.Usecases) {
	registerAuthRoutes(api, ucs)
	registerProductRoutes(api, ucs)
	registerTrendingColorRoutes(api, ucs)
	registerCommonColorRoutes(api, ucs)
	registerCartRoutes(api, ucs)
	registerAddressRoutes(api, ucs)
	registerCheckoutRoutes(api, ucs)
	// registerCartRoutes(api, ucs) // Uncomment when cart is ready
}

// Auth routes
func registerAuthRoutes(api *gin.RouterGroup, ucs *usecase.Usecases) {
	userHandler := NewUserHandler(ucs.User)
	authGroup := api.Group("/auth")
	authGroup.GET("/me", userHandler.Me)
}

// Product routes
func registerProductRoutes(api *gin.RouterGroup, ucs *usecase.Usecases) {
	productHandler := NewProductHandler(ucs.Product)
	productGroup := api.Group("/products")
	{
		productGroup.POST("", productHandler.AddProduct)
		productGroup.GET("/:id", productHandler.GetProductById)
		productGroup.GET("", productHandler.GetProductsList)
		productGroup.PATCH("/:id", productHandler.UpdateTheProductById)
		productGroup.DELETE("/:id", productHandler.DeleteProductByID)

		// Color routes for product
		productGroup.POST("/:id/colors", productHandler.AddNewColorToProduct)
		productGroup.PATCH("/:id/colors/:color_id", productHandler.UpdateProductColor)
	}
}

// Trending colors
func registerTrendingColorRoutes(api *gin.RouterGroup, ucs *usecase.Usecases) {
	trendingColorHandler := NewTrendingColorHandler(ucs.TrendingColor)
	trendingColorGroup := api.Group("/colors/trending")
	{
		trendingColorGroup.POST("", trendingColorHandler.AddTrendingColorHandler)
		trendingColorGroup.DELETE("/:trending_id", trendingColorHandler.DeleteTrendingColorByIDHandler)
		trendingColorGroup.GET("", trendingColorHandler.GetAllTrendingColorsHandler)
		trendingColorGroup.PATCH("/:trending_id", trendingColorHandler.UpdateTrendingColorHandler)
	}
}
func registerCartRoutes(api *gin.RouterGroup, ucs *usecase.Usecases) {
	cartHandler := NewCartHandler(ucs.Cart)
	cartHandlerGroup := api.Group("/cart")
	{
		cartHandlerGroup.GET("", cartHandler.GetCartDetailsHandler)
		cartHandlerGroup.POST("", cartHandler.AddCartItemHandler)
		cartHandlerGroup.GET("/:cart_item_id", cartHandler.GetCartItemByIdHandler)
		cartHandlerGroup.PATCH("/:cart_item_id", cartHandler.UpdateCartItemHandler)
		cartHandlerGroup.DELETE("/:cart_item_id", cartHandler.DeleteCartItemByIdHandler)
	}
}

// Common colors
func registerCommonColorRoutes(api *gin.RouterGroup, ucs *usecase.Usecases) {
	commonColorHandler := NewCommonColorHandler(ucs.CommonColor)
	commonColorGroup := api.Group("/colors/daily")
	{
		commonColorGroup.POST("", commonColorHandler.AddCommonColorHandler)
		commonColorGroup.DELETE("/:common_id", commonColorHandler.DeleteCommonColorByIDHandler)
		commonColorGroup.GET("", commonColorHandler.GetAllCommonColorsHandler)
		commonColorGroup.PATCH("/:common_id", commonColorHandler.UpdateCommonColorHandler)
	}
}
func registerAddressRoutes(api *gin.RouterGroup, ucs *usecase.Usecases) {
	addressHandler := NewAddressHandler(ucs.Address)
	addressGroup := api.Group("/addresses")
	{
		addressGroup.POST("", addressHandler.CreateAddressHandler)
		addressGroup.GET("", addressHandler.GetAllAddressesHandler)
		addressGroup.GET("/:address_id", addressHandler.GetAddressByIDHandler)
		addressGroup.PATCH("/:address_id", addressHandler.UpdateAddressHandler)
		addressGroup.DELETE("/:address_id", addressHandler.DeleteAddressHandler)
	}
}

func registerCheckoutRoutes(api *gin.RouterGroup, ucs *usecase.Usecases) {
	checkoutHandler := NewCheckoutHandler(ucs.Checkout)
	checkoutGroup := api.Group("/checkout")
	{
		checkoutGroup.POST("/summary", checkoutHandler.GetCheckoutSummaryHandler)
	}
}
