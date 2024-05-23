package config

import (
	"os"
	"strconv"
)

type Config struct {
	RateLimitPerMinute int
	RedisAddr          string
}

func LoadConfig() (*Config, error) {
	rateLimitPerMinute := loadIntFromEnv("RATE_LIMIT_PER_MINUTE", 10)
	redisAddr := loadStringFromEnv("REDIS_ADDR", "localhost:6379")

	return &Config{
		RateLimitPerMinute: rateLimitPerMinute,
		RedisAddr:          redisAddr,
	}, nil
}

func loadIntFromEnv(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}

func loadStringFromEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
