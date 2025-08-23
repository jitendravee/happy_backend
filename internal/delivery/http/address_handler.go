package http

import (
	"happy_backend/internal/entities"
	"happy_backend/internal/usecase"
	"happy_backend/pkg/helper"
	"happy_backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	uc *usecase.AddressUseCase
}

func NewAddressHandler(uc *usecase.AddressUseCase) *AddressHandler {
	return &AddressHandler{
		uc: uc,
	}
}

// Create a new address
func (h *AddressHandler) CreateAddressHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	if userId == "" {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req entities.Address
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid address payload")
		return
	}

	address, err := h.uc.CreateAddressUseCase(userId.(string), &req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Address created successfully", address)
}

// Get all addresses for a user
func (h *AddressHandler) GetAllAddressesHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	if userId == "" {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	addresses, err := h.uc.GetAllAddressesUseCase(userId.(string))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Addresses fetched successfully", addresses)
}

// Get a single address by ID
func (h *AddressHandler) GetAddressByIDHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	addressId := helper.GetParamStringOrAbort(c, "address_id")
	if userId == "" || addressId == "" {
		return
	}

	address, err := h.uc.GetAddressByIDUseCase(userId.(string), addressId)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Address fetched successfully", address)
}

// Update an address
func (h *AddressHandler) UpdateAddressHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	addressId := helper.GetParamStringOrAbort(c, "address_id")
	if userId == "" || addressId == "" {
		return
	}

	var req entities.Address
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid address payload")
		return
	}

	updated, err := h.uc.UpdateAddressUseCase(userId.(string), addressId, &req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Address updated successfully", updated)
}

// Delete an address
func (h *AddressHandler) DeleteAddressHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	addressId := helper.GetParamStringOrAbort(c, "address_id")
	if userId == "" || addressId == "" {
		return
	}

	err := h.uc.DeleteAddressUseCase(userId.(string), addressId)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Address deleted successfully", nil)
}
