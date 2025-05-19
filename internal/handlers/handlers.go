package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sosshik/weather-subscription-api/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service is healthy:)")
	})

	// subscriptions
	{
		e.POST("/subscribe", h.SubscribeHandler)
		e.GET("/confirm/:token", h.ConfirmHandler)
		e.GET("/unsubscribe/:token", h.UnsubscribeHandler)
	}

	// weather
	{
		e.GET("/weather", h.GetWeatherHandler)
	}

	return e
}
