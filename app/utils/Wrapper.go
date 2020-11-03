package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WrapAPIError(c * gin.Context, message string, code int) {
	c.JSON(code,map[string]interface{}{
		"code":         code,
		"error_type":    http.StatusText(code),
		"error_details": message,
	})
}

func WrapAPISuccess(c *gin.Context, message string, code int) {
	c.JSON(code,map[string]interface{}{
		"code":   code,
		"status": message,
	})
}

func WrapAPIData(c *gin.Context, data interface{}, code int, message string) {
	c.JSON(code,map[string]interface{}{
		"code":   code,
		"status": message,
		"data":   data,
	})
}
