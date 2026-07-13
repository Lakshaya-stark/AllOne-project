package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppEnv  string
	AppPort string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string

	RedisHost     string
	RedisPort     string
	RedisPassword string

	JWTSecret            string
	JWTExpiration        string
	JWTRefreshExpiration string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName: getEnv("APP_NAME"),
		AppEnv: getEnv("APP_ENV"),
		AppPort: getEnv("APP_PORT"),

		PostgresHost: getEnv("POSTGRES_HOST"),
		PostgresPort: getEnv("POSTGRES_PORT"),
		PostgresUser: getEnv("POSTGRES_USER"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD"),
		PostgresDB: getEnv("POSTGRES_DB"),

		RedisHost: getEnv("REDIS_HOST"),
		RedisPort: getEnv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),

		JWTSecret: getEnv("JWT_SECRET"),
		JWTExpiration: getEnv("JWT_EXPIRATION"),
		JWTRefreshExpiration: getEnv("JWT_REFRESH_EXPIRATION"),
	}

	return cfg
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}