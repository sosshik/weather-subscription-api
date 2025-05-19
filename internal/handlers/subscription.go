package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/sosshik/weather-subscription-api/internal/dto"
	"net/http"
	"strings"
)

// SubscribeHandler godoc
// @Summary      Subscribe to weather updates
// @Description  Subscribe an email to receive weather updates for a specific city with a chosen frequency.
// @Tags         Subscription
// @Accept       json
// @Produce      json
// @Param        subscribeRequest  body      dto.SubscribeRequestDTO  true  "Subscription request"
// @Success      200  {object}  map[string]string  "Subscription successful. Confirmation email sent."
// @Failure      400  {object}  map[string]string  "Invalid input"
// @Failure      409  {object}  map[string]string  "Email already subscribed"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /subscribe [post]
func (h *Handler) SubscribeHandler(c echo.Context) error {
	var req dto.SubscribeRequestDTO
	if err := c.Bind(&req); err != nil {
		log.Warnf("[SubscribeHandler] Unable to decode JSON: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	err := req.Validate()
	if err != nil {
		log.Warnf("[SubscribeHandler] Invalid request payload: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	token, err := h.services.Subscribe(req)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Email already subscribed"})
		}
		log.Warnf("[SubscribeHandler] Unable to subscribe: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Unable to subscribe: %s", err)})
	}

	err = h.services.SendConfirmationEmail(req.Email, token)
	if err != nil {
		_ = h.services.Unsubscribe(token)
		log.Warnf("[SubscribeHandler] Unable to send email: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Unable to subscribe: %s", err)})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Subscription successful. Confirmation email sent."})
}

// ConfirmHandler godoc
// @Summary      Confirm email subscription
// @Description  Confirms a subscription using the token sent via confirmation email.
// @Tags         Subscription
// @Produce      json
// @Param        token  path      string  true  "Confirmation token"
// @Success      200  {object}  map[string]string  "Subscription confirmed successfully"
// @Failure      400  {object}  map[string]string  "Invalid token"
// @Failure      404  {object}  map[string]string  "Token not found"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /confirm/{token} [get]
func (h *Handler) ConfirmHandler(c echo.Context) error {
	token := c.Param("token")
	if token == "" {
		log.Warnf("[ConfirmHandler] Invalid request payload: token cannot be empty")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid token"})
	}

	if err := h.services.Confirm(token); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Token not found"})
		}
		log.Warnf("[ConfirmHandler] Unable to confirm subscription: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Unable to confirm subscription: %s", err)})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Subscription confirmed successfully"})
}

// UnsubscribeHandler godoc
// @Summary      Unsubscribe from weather updates
// @Description  Unsubscribes an email from weather updates using the token sent in emails.
// @Tags         Subscription
// @Produce      json
// @Param        token  path      string  true  "Unsubscribe token"
// @Success      200  {object}  map[string]string  "Unsubscribed successfully"
// @Failure      400  {object}  map[string]string  "Invalid token"
// @Failure      404  {object}  map[string]string  "Token not found"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /unsubscribe/{token} [get]
func (h *Handler) UnsubscribeHandler(c echo.Context) error {
	token := c.Param("token")
	if token == "" {
		log.Warnf("[UnsubscribeHandler] Invalid request payload: token cannot be empty")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid token"})
	}

	if err := h.services.Unsubscribe(token); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Token not found"})
		}
		log.Warnf("[UnsubscribeHandler] Unable to unsubscribe: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Unable to unsubscribe: %s", err)})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Unsubscribed successfully"})
}
