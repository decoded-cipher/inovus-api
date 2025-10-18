package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Initialize the HTTP server with middleware
func initHTTPServer() *echo.Echo {
	var server = echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	return server
}

// Entry point of the application
func main() {
	server := initHTTPServer()
	InitHTTPHandler(server)
	server.Start(":8080")
}
