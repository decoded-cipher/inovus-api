package main

import (
	"log"

	"github.com/decoded-cipher/inovus-api/config"
	"github.com/labstack/echo/v4"
)

func main() {
	// Connect to database
	// if err := config.ConnectDatabase(); err != nil {
	// 	log.Fatalf("Database connection failed: %v", err)
	// }
	// defer config.CloseDatabase()

	// Initialize server
	server := echo.New()
	// server.Use(middleware.Logger())
	// server.Use(middleware.Recover())

	// Setup routes
	InitHTTPHandler(server)

	// Start server
	port := ":" + config.GetServerPort()
	log.Printf("Server starting on %s", port)
	server.Start(port)
}
