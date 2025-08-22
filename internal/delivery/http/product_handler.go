package http

import (
	"happy_backend/internal/entities"
	"happy_backend/internal/usecase"
	"happy_backend/pkg/helper"
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

func (h *ProductHandler) GetProductById(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		response.Error(c, http.StatusBadRequest, "Missing product id")
		return
	}

	product, err := h.uc.GetProductByID(productID)
	if err != nil {
		if err.Error() == "product not found" {
			response.Error(c, http.StatusNotFound, err.Error())
		} else {
			response.Error(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response.Success(c, http.StatusOK, "Product fetched successfully", product)
}
func (h *ProductHandler) DeleteProductByID(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		response.Error(c, http.StatusBadRequest, "Missing product id")
		return
	}
	err := h.uc.DeleteProductByID(productID)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
	}
	response.Success(c, http.StatusCreated, "Product Deleted successfully", nil)

}
func (h *ProductHandler) AddNewColorToProduct(c *gin.Context) {
	var req []entities.Color

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Color payload")
		return
	}

	productID := helper.GetParamStringOrAbort(c, "id")
	if productID == "" {
		return
	}

	addedColors, err := h.uc.AddProductColors(productID, &req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Colors added successfully", addedColors)
}
func (h *ProductHandler) UpdateProductColor(c *gin.Context) {
	var req entities.Color
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Color payload")
		return
	}
	productID := helper.GetParamStringOrAbort(c, "id")
	if productID == "" {
		return
	}
	colorId := helper.GetParamStringOrAbort(c, "color_id")
	if colorId == "" {
		return
	}
	updatedColor, err := h.uc.UpdateProductColor(productID, colorId, &req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Colors added successfully", updatedColor)

}
func (h *ProductHandler) UpdateTheProductById(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		response.Error(c, http.StatusBadRequest, "Missing product id")
		return
	}
	var req entities.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid product payload")
		return
	}

	product, err := h.uc.UpdateProductByID(productID, &req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Product Updated successfully", product)

}
func (h *ProductHandler) GetProductsList(c *gin.Context) {
	products, err := h.uc.GetAllProductsUseCase()
	if err != nil {
		if err.Error() == "products list not found" {
			response.Error(c, http.StatusNotFound, err.Error())
		} else {
			response.Error(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	response.Success(c, http.StatusOK, "Products list fetched successfully", products)

}
