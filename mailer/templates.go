package mailer

import "fmt"

// GenerateTemplate - Responsible for generating content in Email Body
func GenerateTemplate(recipientUserName string, verificationLink string) string {
	return fmt.Sprintf(`
		<p>Hi! %v </p>
		<p>Thanks for joining the community.</p>
		<p>Welcome to Ignitus! To verify your email, click the following link: </p>
		<a href='%v' target=\"__blank\">%v</a>`, recipientUserName, verificationLink, verificationLink,
	)
}
