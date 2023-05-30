package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	SERVER_HOST    string
	SERVER_PORT    string
	ENABLE_LIMITER bool
	ENABLE_LOGGER  bool
	ENABLE_MONITOR bool
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		SERVER_HOST:    os.Getenv("SERVER_HOST"),
		SERVER_PORT:    os.Getenv("SERVER_PORT"),
		ENABLE_LIMITER: parseEnvBool("ENABLE_LIMITER", true),
		ENABLE_LOGGER:  parseEnvBool("ENABLE_LOGGER", true),
		ENABLE_MONITOR: parseEnvBool("ENABLE_LIMITER", false),
	}
}

func parseEnvInt(key string, defaultValue int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}

func parseEnvBool(key string, defaultValue bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	parsedValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}
