package handlers

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// GetWeatherHandler godoc
// @Summary      Get current weather
// @Description  Retrieves the current weather information for a specified city.
// @Tags         Weather
// @Accept       json
// @Produce      json
// @Param        city  query     string  true  "City name for weather forecast"
// @Success      200   {object}  dto.WeatherDTO  "Weather data returned successfully"
// @Failure      400   {string}  string          "Invalid request"
// @Failure      404   {string}  string          "City not found"
// @Failure      500   {string}  string          "Internal server error"
// @Router       /weather [get]
func (h *Handler) GetWeatherHandler(c echo.Context) error {
	city := c.QueryParam("city")
	if city == "" {
		log.Warnf("[GetWeatherHandler] City is empty")
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	weather, err := h.services.GetWeather(city)
	if err != nil {
		if err.Error() == "city not found" {
			return c.String(http.StatusNotFound, "City not found")
		}
		log.Warnf("[GetWeatherHandler] Error getting weather: %s", err.Error())
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, weather)
}
