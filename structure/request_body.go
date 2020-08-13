package structure

import "github.com/gin-gonic/gin"

// SendConfirmationMailRequestBody will help to extract data from request body
type SendConfirmationMailRequestBody struct {
	ToEmail          string `json:"to_email"`
	VerificationLink string `json:"verification_link"`
}

// BindBody will bind request body to struct
func (rb *SendConfirmationMailRequestBody) BindBody(c *gin.Context) error {
	err := c.BindJSON(&rb)
	return err
}
