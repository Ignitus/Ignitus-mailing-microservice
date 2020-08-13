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
	fromName  = "Team Ignitus"
	fromEmail = os.Getenv("FROM_EMAIL")
)

/* Mail fn will rely on sendGrid server. */
func Mail(to, subject, htmlMessage string) error {
	utils.LogMessage(fmt.Sprintf("Sending mail to %v", to))

	from := mail.NewEmail(fromName, fromEmail)
	mailTo := mail.NewEmail("", to)
	message := mail.NewSingleEmail(from, subject, mailTo, "Hi,\n", htmlMessage)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		utils.LogError(fmt.Sprintf("Error sending mail to %v on subject %v", to, subject), err)
	}

	if response.StatusCode != http.StatusAccepted {
		utils.LogMessage(fmt.Sprintf("Response: %v", response.Body))
	}

	return err
}
