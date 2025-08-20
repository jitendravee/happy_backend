// internal/delivery/http/server.go
package http

import (
	"happy_backend/config"
	"happy_backend/internal/usecase"
	"happy_backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func NewServer(cfg *config.Config, ucs *usecase.Usecases) *gin.Engine {
	r := gin.Default()

	NewUserHandler(r, ucs.User)

	r.GET("/health", func(c *gin.Context) {
		response.Success(c, 200, "Service running", gin.H{"service": "ok"})
	})

	return r
}
