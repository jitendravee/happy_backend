package response

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, APIResponse{
		StatusCode: code,
		Message:    msg,
		Data:       data,
	})
}
func ErrorWithData(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, APIResponse{
		StatusCode: code,
		Message:    msg,
		Data:       data,
	})
	c.Abort()
}
func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, APIResponse{
		StatusCode: code,
		Message:    msg,
	})
	c.Abort()
}
