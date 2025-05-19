package dto

import (
	"errors"
	"fmt"
	"github.com/sosshik/weather-subscription-api/internal/repository"
	"github.com/sosshik/weather-subscription-api/internal/weather"
	"net/mail"
)

type SubscribeRequestDTO struct {
	Email     string `form:"email" json:"email"`
	City      string `form:"city" json:"city"`
	Frequency string `form:"frequency" json:"frequency"`
}

func (d *SubscribeRequestDTO) ToSubscriptionModel() repository.Subscription {
	return repository.Subscription{
		Email:     d.Email,
		City:      d.City,
		Frequency: d.Frequency,
	}
}

func (d *SubscribeRequestDTO) Validate() error {
	if d.Email == "" {
		return errors.New("email is required")
	}
	if !isEmailValid(d.Email) {
		return errors.New("invalid email")
	}
	if d.City == "" {
		return errors.New("city is required")
	}
	if d.Frequency == "" {
		return errors.New("frequency is required")
	}

	if d.Frequency != "hourly" && d.Frequency != "daily" {
		return errors.New("frequency is invalid")
	}

	return nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

type WeatherDTO struct {
	Temperature int    `json:"temperature"`
	Humidity    int    `json:"humidity"`
	Description string `json:"description"`
}

func (w *WeatherDTO) FromWeatherResponse(resp *weather.WeatherResponse) {
	w.Humidity = resp.Main.Humidity
	w.Temperature = int(resp.Main.Temperature)
	if resp.Weather != nil {
		w.Description = fmt.Sprintf("%s - %s", resp.Weather[0].Main, resp.Weather[0].Description)
	}
}
