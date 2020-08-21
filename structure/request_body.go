package structure

import "github.com/gin-gonic/gin"

// RequestBody is body must be followed by /mail/confirmation endpoint
type RequestBody struct {
	RecipientAddress  string `json:"recipientAddress"`
	VerificationLink  string `json:"verificationLink"`
	RecipientUserName string `json:"recipientUserName"`
}

// BindBody - Responsible for binding thed request body to struct.
func (rb *RequestBody) BindBody(c *gin.Context) error {
	err := c.BindJSON(&rb)
	return err
}
