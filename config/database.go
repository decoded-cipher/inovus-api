package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectDatabase establishes a database connection
func ConnectDatabase() error {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "inovus"),
		getEnv("DB_PASSWORD", "inovus123"),
		getEnv("DB_NAME", "inovus"),
		getEnv("DB_SSLMODE", "disable"),
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	DB = db
	log.Println("Database connected successfully")
	return nil
}

// CloseDatabase closes the database connection
func CloseDatabase() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// getEnv returns environment variable or default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetServerPort returns the server port
func GetServerPort() string {
	return getEnv("SERVER_PORT", "8080")
}
