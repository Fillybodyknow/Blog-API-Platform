package utility

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to, subject, body string) error {
	from := os.Getenv("SMTP_FROM")
	pass := os.Getenv("SMTP_PASS")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", from, to, subject, body)

	auth := smtp.PlainAuth("", from, pass, host)
	addr := fmt.Sprintf("%s:%s", host, port)

	err := smtp.SendMail(addr, auth, from, []string{to}, []byte(msg))
	return err
}
