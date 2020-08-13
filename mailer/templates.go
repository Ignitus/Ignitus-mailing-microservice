package mailer

import "fmt"

// GenerateConfirmationHTMLTemplate will generate confirmation html string to pass in Mail function
func GenerateConfirmationHTMLTemplate(urlLink string) string {
	return fmt.Sprintf("<h2>Please verify your email by visiting given url.</h2><br /><a href=\"%v\" target=\"__blank\">%v</a>", urlLink, urlLink)
}
