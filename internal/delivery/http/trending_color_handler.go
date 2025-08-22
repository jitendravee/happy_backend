package http

import (
	"happy_backend/internal/entities"
	"happy_backend/internal/usecase"
	"happy_backend/pkg/helper"
	"happy_backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrendingColorHandler struct {
	uc *usecase.TrendingColorUseCase
}

func NewTrendingColorHandler(uc *usecase.TrendingColorUseCase) *TrendingColorHandler {
	return &TrendingColorHandler{uc: uc}
}

func (h *TrendingColorHandler) GetAllTrendingColorsHandler(c *gin.Context) {
	colors, err := h.uc.GetAllTredingColorsUseCase()
	if err != nil {
		if err.Error() == "trending color list not found" {
			response.Error(c, http.StatusNotFound, err.Error())
		} else {
			response.Error(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	response.Success(c, http.StatusOK, "Trending color list fetched successfully", colors)

}
func (h *TrendingColorHandler) AddTrendingColorHandler(c *gin.Context) {
	var req []entities.TrendingColor
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid trending color payload")
		return
	}
	addedColors, err := h.uc.AddTrendingColorUseCase(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Colors added successfully", addedColors)

}
func (h *TrendingColorHandler) UpdateTrendingColorHandler(c *gin.Context) {
	var req entities.TrendingColor
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid trending color payload")
		return
	}
	trendingId := helper.GetParamStringOrAbort(c, "trending_id")
	if trendingId == "" {
		return
	}
	updateTreadingColor, err := h.uc.UpdateTrendingColorUseCase(trendingId, &req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Colors added successfully", updateTreadingColor)

}

func (h *TrendingColorHandler) DeleteTrendingColorByIDHandler(c *gin.Context) {
	trendingId := helper.GetParamStringOrAbort(c, "trending_id")
	if trendingId == "" {
		return
	}
	err := h.uc.DeleteTrendingColorByIDUseCase(trendingId)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
	}
	response.Success(c, http.StatusCreated, "Trending color deleted successfully", nil)
}
