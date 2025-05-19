package service

import (
	"github.com/sosshik/weather-subscription-api/internal/dto"
	"github.com/sosshik/weather-subscription-api/internal/weather"
)

type WeatherService struct {
	Weather *weather.Weather
}

func NewWeatherService(weather *weather.Weather) *WeatherService {
	return &WeatherService{
		Weather: weather,
	}
}

func (s *WeatherService) GetWeather(city string) (*dto.WeatherDTO, error) {
	resp, err := s.Weather.GetWeather(city)
	if err != nil {
		return nil, err
	}
	w := dto.WeatherDTO{}
	w.FromWeatherResponse(resp)
	return &w, nil
}
