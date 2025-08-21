package http

import (
	"happy_backend/internal/entities"
	"happy_backend/internal/usecase"
	"happy_backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	uc *usecase.ProductUseCase
}

func NewProductHandler(uc *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (h *ProductHandler) AddProduct(c *gin.Context) {
	var req entities.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid product payload")
		return
	}

	product, err := h.uc.AddProduct(&req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Product created successfully", product)
}
