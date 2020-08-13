package mailer

import "fmt"

/* Responsible for generating content in Email Body */
func GenerateTemplate(recipientAddress string, verificationLink string) string {
	return fmt.Sprintf(`
		<p>Hi! %v </p>
		<p>Thanks for joining the community.</p>
		<p>Welcome to Ignitus! To verify your email, click the following link: </p>
		<a href=\"%v\" target=\"__blank\">%v</a>,`, recipientAddress, verificationLink, verificationLink,
	)
}
