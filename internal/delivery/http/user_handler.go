// internal/delivery/http/user_handler.go
package http

import (
	"net/http"

	"happy_backend/internal/entities"
	"happy_backend/internal/usecase"
	"happy_backend/pkg/middleware"
	"happy_backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc *usecase.UserUsecase
}

func NewUserHandler(r *gin.Engine, uc *usecase.UserUsecase) {
	h := &UserHandler{uc: uc}

	// Public routes
	r.POST("/signup", h.SignUp)
	r.POST("/signin", h.SignIn)
	r.POST("/refresh", h.Refresh)

	// Protected routes
	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware(uc.Secret())) // inject JWT secret
	{
		auth.GET("/me", h.Me)
	}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var req entities.User
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Email == "" || req.Password == "" {
		response.Error(c, http.StatusBadRequest, "Email and password required")
		return
	}

	user, accessToken, refreshToken, err := h.uc.Register(&req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
		return
	}

	c.SetCookie("access_token", accessToken, 900, "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, 604800, "/", "", false, true)

	response.Success(c, http.StatusCreated, "User created successfully", gin.H{
		"email":         user.Email,
		"id":            user.ID,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
func (h *UserHandler) Me(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.uc.GetByID(userID.(string))
	if err != nil || user == nil {
		response.Error(c, http.StatusNotFound, "User not found")
		return
	}

	response.Success(c, http.StatusOK, "User profile", gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

func (h *UserHandler) SignIn(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, accessToken, refreshToken, err := h.uc.Login(req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	// HttpOnly cookies (set Secure=true in production HTTPS)
	c.SetCookie("access_token", accessToken, 900, "/", "", false, true)         // 15 min
	c.SetCookie("refresh_token", refreshToken, 7*24*3600, "/", "", false, true) // 7 days

	response.Success(c, http.StatusOK, "Login successful", gin.H{
		"email":         user.Email,
		"id":            user.ID,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *UserHandler) Refresh(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil || rt == "" {
		response.Error(c, http.StatusUnauthorized, "Missing refresh token")
		return
	}

	newAT, err := h.uc.RefreshAccessToken(rt)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.SetCookie("access_token", newAT, 900, "/", "", false, true)

	response.Success(c, http.StatusOK, "Token refreshed", gin.H{
		"access_token": newAT,
	})
}
