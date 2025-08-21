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

	userHandler := NewUserHandler(ucs.User)
	r.POST("/signup", userHandler.SignUp)
	r.POST("/signin", userHandler.SignIn)
	r.POST("/refresh", userHandler.Refresh)

	// --- Protected API group (requires access_token cookie) ---
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(ucs.User.Secret())) // applies auth to ALL children
	{
		// Auth-specific routes
		authGroup := api.Group("/auth")
		{
			authGroup.GET("/me", userHandler.Me)
		}

		// Product routes (example)
		productGroup := api.Group("/products")
		productHandler := NewProductHandler(ucs.Product)
		{
			productGroup.POST("", productHandler.AddProduct)
			// productGroup.POST("/", productHandler.Create)
		}

		// Cart routes (example)
		// cartGroup := api.Group("/cart")
		// {
		// 	// cartGroup.GET("/", cartHandler.GetCart)
		// 	// cartGroup.POST("/add", cartHandler.AddItem)
		// }
	}

	r.GET("/health", func(c *gin.Context) {
		response.Success(c, 200, "Service running", gin.H{"service": "ok"})
	})

	return r
}
