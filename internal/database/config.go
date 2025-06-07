package database

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Host            string
	Port            int
	User            string
	Password        string
	Database        string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	Environment     string // "local", "production", "staging"
}

// LoadConfigFromEnv loads database configuration from environment variables
func LoadConfigFromEnv() (*Config, error) {
	environment := getEnvOrDefault("APP_ENV", "local")

	config := &Config{
		Environment: environment,
	}

	// Set defaults based on environment
	if environment == "local" {
		config.Host = getEnvOrDefault("DB_HOST", "localhost")
		config.User = getEnvOrDefault("DB_USER", "jonathanpetrone") // Your local user
		config.Database = getEnvOrDefault("DB_NAME", "aitarot")
		config.SSLMode = getEnvOrDefault("DB_SSLMODE", "disable")
	} else {
		// Production defaults
		config.Host = getEnvOrDefault("DB_HOST", "")
		config.User = getEnvOrDefault("DB_USER", "")
		config.Database = getEnvOrDefault("DB_NAME", "postgres")
		config.SSLMode = getEnvOrDefault("DB_SSLMODE", "require")
	}

	config.Password = os.Getenv("DB_PASSWORD") // Always from env

	// Parse port
	portStr := getEnvOrDefault("DB_PORT", "5432")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT value: %s", portStr)
	}
	config.Port = port

	// Parse connection pool settings
	config.MaxOpenConns = getEnvIntOrDefault("DB_MAX_OPEN_CONNS", 25)
	config.MaxIdleConns = getEnvIntOrDefault("DB_MAX_IDLE_CONNS", 25)

	// Parse connection timeouts
	maxLifetimeStr := getEnvOrDefault("DB_CONN_MAX_LIFETIME", "5m")
	maxLifetime, err := time.ParseDuration(maxLifetimeStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_CONN_MAX_LIFETIME value: %s", maxLifetimeStr)
	}
	config.ConnMaxLifetime = maxLifetime

	maxIdleTimeStr := getEnvOrDefault("DB_CONN_MAX_IDLE_TIME", "30s")
	maxIdleTime, err := time.ParseDuration(maxIdleTimeStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_CONN_MAX_IDLE_TIME value: %s", maxIdleTimeStr)
	}
	config.ConnMaxIdleTime = maxIdleTime

	// Note: Password can be empty for local development setups
	// No validation needed - empty password is valid for local PostgreSQL

	return config, nil
}

// ConnectionString builds a PostgreSQL connection string from the config
func (c *Config) ConnectionString() string {
	if c.Password == "" {
		// Omit password parameter entirely if empty
		return fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s sslmode=%s",
			c.Host, c.Port, c.User, c.Database, c.SSLMode,
		)
	}
	// Force IPv4 by adding prefer_simple_protocol=true
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s prefer_simple_protocol=true",
		c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode,
	)
}

// Helper functions
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvIntOrDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
