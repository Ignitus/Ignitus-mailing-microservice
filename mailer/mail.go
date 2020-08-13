package mailer

import (
	"fmt"
	"net/http"
	"os"
	"github.com/Ignitus/ignitus-mailing-microservice/utils"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	senderName  = "Team Ignitus"
	senderEmail = os.Getenv("SENDER_EMAIL")
)

/* Mail fn will rely on sendGrid server. */
func Mail(recipientAddress, recipientUserName, subject, template string) error {
	utils.LogMessage(fmt.Sprintf("Sending mail to %v", recipientAddress))

	sender := mail.NewEmail(senderName, senderEmail)
	recipient := mail.NewEmail(recipientUserName, recipientAddress)
	message := mail.NewSingleEmail(sender, subject, recipient, "Hi,\n", template)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		utils.LogError(fmt.Sprintf("Error sending mail to %v on subject %v", recipientAddress, subject), err)
	}

	if response.StatusCode != http.StatusAccepted {
		utils.LogMessage(fmt.Sprintf("Response: %v", response.Body))
	}

	return err
}
