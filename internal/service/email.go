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
	return s.Emailer.Send(
		[]string{to},
		"Confirm your subscription",
		fmt.Sprintf("Click here to confirm: %s/confirm/%s", cfg.ServiceDomain, token),
	)
}
