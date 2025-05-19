package emailer

import (
	"fmt"
	"net/smtp"
	"strings"
)

type EmailSender struct {
	From     string
	Password string
	Host     string
	Port     int
}

func NewEmailSender(from, password, host string, port int) *EmailSender {
	return &EmailSender{
		From:     from,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

func (e *EmailSender) Send(to []string, subject, body string) error {
	auth := smtp.PlainAuth("", e.From, e.Password, e.Host)

	// Join recipients with commas for the "To:" header
	toHeader := strings.Join(to, ", ")

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", toHeader, subject, body))

	addr := fmt.Sprintf("%s:%d", e.Host, e.Port)

	return smtp.SendMail(addr, auth, e.From, to, msg)
}
