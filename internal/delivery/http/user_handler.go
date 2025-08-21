// internal/delivery/http/user_handler.go
package http

import (
	"net/http"

	"happy_backend/internal/entities"
	"happy_backend/internal/usecase"
	"happy_backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
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

	// Set HttpOnly cookies
	c.SetCookie("access_token", accessToken, 900, "/", "", false, true)         // 15 min
	c.SetCookie("refresh_token", refreshToken, 7*24*3600, "/", "", false, true) // 7 days

	// Don’t return tokens in JSON
	response.Success(c, http.StatusCreated, "User created successfully", gin.H{
		"email": user.Email,
		"id":    user.ID,
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

	// Set HttpOnly cookies
	c.SetCookie("access_token", accessToken, 900, "/", "", false, true)         // 15 min
	c.SetCookie("refresh_token", refreshToken, 7*24*3600, "/", "", false, true) // 7 days

	// Don’t return tokens in JSON
	response.Success(c, http.StatusOK, "Login successful", gin.H{
		"email": user.Email,
		"id":    user.ID,
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

	response.Success(c, http.StatusOK, "Token refreshed", nil)
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
