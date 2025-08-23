package http

import (
	"happy_backend/internal/entities"
	"happy_backend/internal/usecase"
	"happy_backend/pkg/helper"
	"happy_backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	uc *usecase.CartUseCase
}

func NewCartHandler(uc *usecase.CartUseCase) *CartHandler {
	return &CartHandler{
		uc: uc,
	}
}

func (h *CartHandler) GetCartDetailsHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	cart, err := h.uc.GetCartDetailsUseCase(userId.(string))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Cart fetched successfully", cart)

}

func (h *CartHandler) AddCartItemHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	if userId == "" {
		return
	}
	var req entities.CartItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid CartItem payload")
		return
	}
	err := h.uc.AddItemToCartUseCase(userId.(string), &req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Item added to the cart", nil)
}
func (h *CartHandler) GetCartItemByIdHandler(c *gin.Context) {
	itemId := helper.GetParamStringOrAbort(c, "cart_item_id")
	if itemId == "" {
		return
	}
	cartItem, err := h.uc.GetCartItemById(itemId)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Item added to the cart", cartItem)

}
func (h *CartHandler) DeleteCartItemByIdHandler(c *gin.Context) {
	itemId := helper.GetParamStringOrAbort(c, "cart_item_id")
	if itemId == "" {
		return
	}

	err := h.uc.DeleteCartItemById(itemId)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Cart item deleted successfully", nil)
}
func (h *CartHandler) UpdateCartItemHandler(c *gin.Context) {
	itemId := helper.GetParamStringOrAbort(c, "cart_item_id")
	if itemId == "" {
		return
	}
	var req entities.CartItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid CartItem payload")
		return
	}
	updatedItem, err := h.uc.UpdateTheCartItemUseCase(itemId, &req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Item added to the cart", updatedItem)
}
