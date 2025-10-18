package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HealthCheckResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message,omitempty"`
	Timestamp string `json:"timestamp"`
}

func HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthCheckResponse{
		Status:    http.StatusText(http.StatusOK),
		Message:   "Service is healthy",
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func baseHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Welcome to the Inovus API"})
}