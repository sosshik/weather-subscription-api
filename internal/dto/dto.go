package dto

import (
	"errors"
	"fmt"
	"github.com/sosshik/weather-subscription-api/internal/repository"
	"github.com/sosshik/weather-subscription-api/internal/weather"
)

type SubscribeRequestDTO struct {
	Email     string `json:"email"`
	City      string `json:"city"`
	Frequency string `json:"frequency"`
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
