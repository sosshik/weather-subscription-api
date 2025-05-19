package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Weather struct {
	ApiKey string
}

type WeatherResponse struct {
	Main struct {
		Temperature float64 `json:"temp"`
		Humidity    int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

func NewWeather(apiKey string) *Weather {
	return &Weather{
		ApiKey: apiKey,
	}
}

func (w *Weather) GetWeather(city string) (*WeatherResponse, error) {
	fmt.Printf(w.ApiKey)
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city,
		w.ApiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if resp.StatusCode == http.StatusNotFound {
			return nil, errors.New("city not found")
		}
		return nil, fmt.Errorf("failed to fetch weather, status: %d", resp.StatusCode)
	}

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, err
	}

	return &weatherResp, nil
}
