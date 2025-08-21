// pkg/middleware/auth.go
package middleware

import (
	"happy_backend/pkg/jwt"
	"happy_backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("access_token")
		if err != nil || token == "" {
			response.ErrorWithData(c, http.StatusUnauthorized, "Unauthorized, please login", gin.H{"redirect": "/login"})
			// optionally include redirect info in Data
			c.Abort()
			return
		}

		claims, err := jwt.ValidateToken(token, secret)
		if err != nil {
			response.ErrorWithData(c, http.StatusUnauthorized, "Unauthorized, please login", gin.H{"redirect": "/login"})
			c.Abort()
			return
		}

		userID := claims["user_id"].(string)
		c.Set("user_id", userID)
		c.Next()
	}
}
