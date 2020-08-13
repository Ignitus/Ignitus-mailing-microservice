package structure

import "github.com/gin-gonic/gin"

// ErrorResponse is used to response error consistently
func ErrorResponse(message string) gin.H {
	return gin.H{
		"error":   true,
		"message": message,
	}
}

// SuccessResponse is used to response success consistently
func SuccessResponse(data interface{}) gin.H {
	return gin.H{
		"error": false,
		"data":  data,
	}
}
