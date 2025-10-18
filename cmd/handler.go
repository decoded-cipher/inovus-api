package main

import "github.com/labstack/echo/v4"

func InitHTTPHandler(server *echo.Echo) {
	server.GET("/", baseHandler)
	server.GET("/health", HealthCheckHandler)
}
