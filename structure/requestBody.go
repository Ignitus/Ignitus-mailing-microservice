package structure
import "github.com/gin-gonic/gin"

type RequestBody struct {
	RecipientAddress  string `json:"recipientAddress"`
	VerificationLink  string `json:"verificationLink"`
	RecipientUserName string `json:"recipientUserName"`
}

/* Responsible for binding thed request body to struct. */
func (rb *RequestBody) BindBody(c *gin.Context) error {
	err := c.BindJSON(&rb)
	return err
}
