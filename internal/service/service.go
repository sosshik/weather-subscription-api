package service

import (
	"github.com/sosshik/weather-subscription-api/internal/dto"
	"github.com/sosshik/weather-subscription-api/internal/emailer"
	"github.com/sosshik/weather-subscription-api/internal/repository"
	"github.com/sosshik/weather-subscription-api/internal/weather"
)

type Subscription interface {
	Subscribe(sub dto.SubscribeRequestDTO) (string, error)
	Confirm(token string) error
	Unsubscribe(token string) error
}

type Email interface {
	SendConfirmationEmail(to, token string) error
}

type Weather interface {
	GetWeather(city string) (*dto.WeatherDTO, error)
}

type Service struct {
	Subscription
	Email
	Weather
}

func NewService(repo *repository.Repository, sender *emailer.EmailSender, weather *weather.Weather) *Service {
	return &Service{
		Subscription: NewSubscriptionService(repo),
		Email:        NewEmailService(sender),
		Weather:      NewWeatherService(weather),
	}
}
