package http

import (
	"happy_backend/internal/entities"
	"happy_backend/internal/usecase"
	"happy_backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckoutHandler struct {
	uc *usecase.CheckoutUseCase
}

func NewCheckoutHandler(uc *usecase.CheckoutUseCase) *CheckoutHandler {
	return &CheckoutHandler{
		uc: uc,
	}
}

func (h *CheckoutHandler) GetCheckoutSummaryHandler(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists || userID == "" {
		response.Error(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	var items []entities.CartItem
	if err := c.ShouldBindJSON(&items); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid CartItem payload")
		return
	}
	const deliveryCharge float32 = 50.0
	const taxPercent float32 = 10.0

	summary, err := h.uc.GetCheckoutSummaryUseCase(userID.(string), &items, deliveryCharge, taxPercent)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Checkout summary fetched successfully", summary)
}
