package helper

import (
	"happy_backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetParamStringOrAbort fetches a URL param from context.
// If the param is empty, it returns a JSON error using response.Error and aborts the request.
func GetParamStringOrAbort(c *gin.Context, key string) string {
	param := c.Param(key)
	if param == "" {
		response.Error(c, http.StatusBadRequest, "Missing "+key)
		return ""
	}
	return param
}
