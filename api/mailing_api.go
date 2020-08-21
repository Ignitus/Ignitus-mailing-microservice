package api

import (
	"net/http"

	"github.com/Ignitus/ignitus-mailing-microservice/mailer"
	"github.com/Ignitus/ignitus-mailing-microservice/structure"
	"github.com/Ignitus/ignitus-mailing-microservice/utils"
	"github.com/gin-gonic/gin"
)

// MailingAPI is endpoint to send confirmation mail to user.
func MailingAPI(c *gin.Context) {
	var userData structure.RequestBody
	err := userData.BindBody(c)
	if err != nil {
		utils.LogError("Error binding json body.", err)
		c.JSON(http.StatusBadRequest, structure.ErrorResponse("Invalid body."))
		return
	}
	recipientAddress := userData.RecipientAddress
	recipientUserName := userData.RecipientUserName
	verificationLink := userData.VerificationLink
	template := mailer.GenerateTemplate(recipientUserName, verificationLink)

	err = mailer.Mail(recipientAddress, recipientUserName, mailer.Subject, template)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structure.ErrorResponse("Email was not sent."))
		return
	}

	c.JSON(http.StatusOK, structure.SuccessResponse("Successfully sent mail."))
}
