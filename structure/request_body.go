package structure
import "github.com/gin-gonic/gin"

type RequestBody struct {
	ToEmail          string `json:"to_email"`
	VerificationLink string `json:"verification_link"`
}

/* Responsible for binding thed request body to struct. */
func (rb *RequestBody) BindBody(c *gin.Context) error {
	err := c.BindJSON(&rb)
	return err
}
