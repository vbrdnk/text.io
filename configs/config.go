package configs

import (
	"os"
	"strconv"
)

type Config struct {
	Port    int
	Timeout int
	Debug   bool

	// Database configuration
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() Config {
	config := Config{
		// Default values
		Port:    8080,
		Timeout: 30,
		Debug:   false,

		// Default database values (for Docker)
		DBHost:     "localhost",
		DBPort:     5432,
		DBUser:     "postgres",
		DBPassword: "postgres",
		DBName:     "goapi",
	}

	if port := os.Getenv("PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.Port = p
		}
	}

	if timeout := os.Getenv("TIMEOUT"); timeout != "" {
		if t, err := strconv.Atoi(timeout); err == nil {
			config.Timeout = t
		}
	}

	if debug := os.Getenv("DEBUG"); debug == "true" {
		config.Debug = true
	}

	// Database config from environment
	if host := os.Getenv("DB_HOST"); host != "" {
		config.DBHost = host
	}

	if port := os.Getenv("DB_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.DBPort = p
		}
	}

	if user := os.Getenv("DB_USER"); user != "" {
		config.DBUser = user
	}

	if password := os.Getenv("DB_PASSWORD"); password != "" {
		config.DBPassword = password
	}

	if name := os.Getenv("DB_NAME"); name != "" {
		config.DBName = name
	}

	return config
}
