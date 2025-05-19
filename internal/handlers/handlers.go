package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/sosshik/weather-subscription-api/docs"
	"github.com/sosshik/weather-subscription-api/internal/service"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	e.Use(middleware.CORS())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service is healthy:)")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

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
