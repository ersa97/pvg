package helpers

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(sub, address, name, key string) {

	from := mail.NewEmail(os.Getenv("SENDER_NAME"), os.Getenv("SENDER_EMAIL"))
	subject := sub
	to := mail.NewEmail(name, address)
	plainTextContent := "your key"
	htmlContent := "Here is your key <strong>" + key + "</strong> please copy this key and use it to verify your account!"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_APIKEY"))
	response, err := client.Send(message)
	if err != nil {
		fmt.Println("Unable to send email")
	} else {
		fmt.Println("Email sent")
	}
	fmt.Println(response)
	fmt.Println(response.Body)
}
