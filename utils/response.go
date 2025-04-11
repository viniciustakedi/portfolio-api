package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Message(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, gin.H{
		"message":     message,
		"status_code": statusCode,
	})
}

func Data(c *gin.Context, data any, message string, statusCode int) {
	c.JSON(statusCode, gin.H{
		"data":        data,
		"message":     message,
		"status_code": statusCode,
	})
}

func Error(c *gin.Context, message string, statusCode ...int) {
	code := http.StatusBadRequest
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	c.JSON(code, gin.H{
		"message":     message,
		"status_code": code,
	})
}
