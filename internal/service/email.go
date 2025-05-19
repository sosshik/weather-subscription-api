package service

import (
	"fmt"
	"github.com/sosshik/weather-subscription-api/internal/config"
	"github.com/sosshik/weather-subscription-api/internal/emailer"
)

type EmailService struct {
	Emailer *emailer.EmailSender
}

func NewEmailService(sender *emailer.EmailSender) *EmailService {
	return &EmailService{sender}
}

func (s *EmailService) SendConfirmationEmail(to, token string) error {
	cfg := config.GetConfig()
	body := fmt.Sprintf(`
    <p>Click here to confirm your subscription:</p>
    <p><a href="%s/confirm/%s">Confirm subscription</a></p>
    <p>Your token is: %s</p>
    <p>You can use this token to unsubscribe or click here to unsubscribe</p>
	<p><a href="%s/unsubscribe/%s">Unsubscribe</a></p>
`, cfg.ServiceDomain, token, token, cfg.ServiceDomain, token)
	return s.Emailer.Send(
		[]string{to},
		"Confirm your subscription",
		body,
	)
}
