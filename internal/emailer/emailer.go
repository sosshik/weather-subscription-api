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

func (e *EmailSender) Send(to []string, subject, body string, isHTML bool) error {
	auth := smtp.PlainAuth("", e.From, e.Password, e.Host)
	toHeader := strings.Join(to, ", ")

	contentType := "text/plain; charset=\"UTF-8\""
	if isHTML {
		contentType = "text/html; charset=\"UTF-8\""
	}

	msg := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: %s\r\n\r\n"+
			"%s",
		toHeader, subject, contentType, body,
	))

	addr := fmt.Sprintf("%s:%d", e.Host, e.Port)

	return smtp.SendMail(addr, auth, e.From, to, msg)
}
