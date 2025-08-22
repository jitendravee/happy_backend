package http

import (
	"happy_backend/internal/entities"
	"happy_backend/internal/usecase"
	"happy_backend/pkg/helper"
	"happy_backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonColorHandler struct {
	uc *usecase.CommonColorUseCase
}

func NewCommonColorHandler(uc *usecase.CommonColorUseCase) *CommonColorHandler {
	return &CommonColorHandler{uc: uc}
}

func (h *CommonColorHandler) GetAllCommonColorsHandler(c *gin.Context) {
	colors, err := h.uc.GetAllTredingColorsUseCase()
	if err != nil {
		if err.Error() == "Common color list not found" {
			response.Error(c, http.StatusNotFound, err.Error())
		} else {
			response.Error(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	response.Success(c, http.StatusOK, "Common color list fetched successfully", colors)

}
func (h *CommonColorHandler) AddCommonColorHandler(c *gin.Context) {
	var req []entities.CommonColor
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Common color payload")
		return
	}
	addedColors, err := h.uc.AddCommonColorUseCase(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Colors added successfully", addedColors)

}
func (h *CommonColorHandler) UpdateCommonColorHandler(c *gin.Context) {
	var req entities.CommonColor
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Common color payload")
		return
	}
	CommonId := helper.GetParamStringOrAbort(c, "common_id")
	if CommonId == "" {
		return
	}
	updateTreadingColor, err := h.uc.UpdateCommonColorUseCase(CommonId, &req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Colors added successfully", updateTreadingColor)

}

func (h *CommonColorHandler) DeleteCommonColorByIDHandler(c *gin.Context) {
	CommonId := helper.GetParamStringOrAbort(c, "common_id")
	if CommonId == "" {
		return
	}
	err := h.uc.DeleteCommonColorByIDUseCase(CommonId)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
	}
	response.Success(c, http.StatusCreated, "Common color deleted successfully", nil)
}
