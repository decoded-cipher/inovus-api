package main

import (
	"log"

	"github.com/decoded-cipher/inovus-api/config"
	"github.com/decoded-cipher/inovus-api/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var queries *models.Queries

func main() {
	// Connect to database
	if err := config.ConnectDatabase(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer config.CloseDatabase()

	// Load queries
	var err error
	queries, err = models.LoadQueries(config.DB, "queries.sql")
	if err != nil {
		log.Fatalf("Failed to load queries: %v", err)
	}

	// Initialize server
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// Setup routes
	InitHTTPHandler(server)

	// Start server
	port := ":" + config.GetServerPort()
	log.Printf("Server starting on %s", port)
	server.Start(port)
}
