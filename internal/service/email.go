package service

import (
	"fmt"
	"github.com/sosshik/weather-subscription-api/internal/emailer"
)

type EmailService struct {
	Emailer *emailer.EmailSender
}

func NewEmailService(sender *emailer.EmailSender) *EmailService {
	return &EmailService{sender}
}

func (s *EmailService) SendConfirmationEmail(to, token string) error {
	return s.Emailer.Send(
		[]string{to},
		"Confirm your subscription",
		fmt.Sprintf("Click here to confirm: http://localhost:8090/confirm/%s", token),
	)
}
