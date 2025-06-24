package utility

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
)

func SendEmail(toEmail, subject, body string) error {
	from := mail.Address{Name: "Blog API Platform", Address: os.Getenv("SMTP_EMAIL")}
	pass := os.Getenv("SMTP_PASS")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	to := mail.Address{Name: "", Address: toEmail}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = subject

	massage := ""
	for k, v := range header {
		massage += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	massage += "\r\n" + body

	auth := smtp.PlainAuth("", from.Address, pass, host)
	err := smtp.SendMail(host+":"+port, auth, from.Address, []string{to.Address}, []byte(massage))
	return err
}
