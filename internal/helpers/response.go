package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Response structure
type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`        // Data bisa diabaikan jika nil
	Validation interface{} `json:"validations,omitempty"` // Data bisa diabaikan jika nil
}

// APIResponse is a helper to standardize API responses
func APIResponse(c *gin.Context, status int, success bool, message string, data interface{}, validation interface{}) {
	if !success {
		log.Printf("Error: %s", message) // Logging jika gagal
	}
	res := Response{
		Success:    success,
		Message:    message,
		Data:       data,
		Validation: validation,
	}
	c.JSON(status, res)
}

func SuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	APIResponse(c, status, true, message, data, nil)
}
func ErrorResponse(c *gin.Context, status int, message string) {
	APIResponse(c, status, false, message, nil, nil)
}
func ErrorValidation(c *gin.Context, status int, validation interface{}) {
	message := "validation fail"
	APIResponse(c, status, false, message, nil, validation)
}
