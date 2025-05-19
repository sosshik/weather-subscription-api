package handlers

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) GetWeatherHandler(c echo.Context) error {
	city := c.QueryParam("city")
	if city == "" {
		log.Warnf("[GetWeatherHandler] City is empty")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	weather, err := h.services.GetWeather(city)
	if err != nil {
		if err.Error() == "city not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "City not found"})
		}
		log.Warnf("[GetWeatherHandler] Error getting weather: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusOK, weather)
}
