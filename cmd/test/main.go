package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jonathanpetrone/aitarot/internal/database"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Test config loading
	config, err := database.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Add these debug lines:
	log.Printf("Loaded DB_NAME from env: '%s'", os.Getenv("DB_NAME"))
	log.Printf("Config Database: '%s'", config.Database)

	// Print config (without password for security)
	log.Printf("Database config loaded successfully:")
	// ... rest of your existing code

	// Print config (without password for security)
	log.Printf("Database config loaded successfully:")
	log.Printf("Host: %s", config.Host)
	log.Printf("Port: %d", config.Port)
	log.Printf("User: %s", config.User)
	log.Printf("Database: %s", config.Database)

	// Test actual connection
	log.Printf("Testing database connection...")
	log.Printf("Connection string: '%s'", config.ConnectionString())
	db, err := database.Connect(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test health check
	log.Printf("Running health check...")
	err = db.HealthCheck(nil)
	if err != nil {
		log.Fatalf("Health check failed: %v", err)
	}

	// Show connection stats
	stats := db.GetStats()
	log.Printf("Connection successful!")
	log.Printf("Open connections: %d", stats.OpenConnections)
	log.Printf("In use: %d", stats.InUse)
	log.Printf("Idle: %d", stats.Idle)
}

func maskPassword(connStr string) string {
	// Simple password masking for logging
	return "host=localhost port=5432 user=jonathanpetrone password=*** dbname=aitarot sslmode=disable"
}
