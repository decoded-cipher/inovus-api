package main

import "github.com/labstack/echo/v4"

func InitHTTPHandler(server *echo.Echo) {

	// Base routes
	server.GET("/", baseHandler)
	server.GET("/health", HealthCheckHandler)
	
	// User routes
	userGroup := server.Group("/users")
	userGroup.GET("/all", getAllUsers)
}
