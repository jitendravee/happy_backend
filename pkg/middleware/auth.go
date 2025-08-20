package middleware

import (
	"net/http"
	"strings"

	"happy_backend/pkg/jwt"
	"happy_backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get token (prefer Cookie, fallback to Header)
		token, err := c.Cookie("access_token")
		if err != nil || token == "" {
			authHeader := c.GetHeader("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				token = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if token == "" {
			response.Error(c, http.StatusUnauthorized, "Missing access token")
			c.Abort()
			return
		}

		// 2. Validate token
		claims, err := jwt.ValidateToken(token, secret)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		// 3. Save user_id in context for handlers
		c.Set("user_id", claims["user_id"].(string))
		c.Next()
	}
}
