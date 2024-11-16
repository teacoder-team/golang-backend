package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	NodeEnv         string
	ApplicationPort int
	ApplicationURL  string
	AllowedOrigin   string
	SiteURL         string

	SwaggerEnabled bool
	SwaggerPrefix  string

	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresURI      string

	RedisUser     string
	RedisPassword string
	RedisHost     string
	RedisPort     int
	RedisURI      string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	applicationPort, err := strconv.Atoi(getEnv("APPLICATION_PORT", "4000"))
	if err != nil {
		return nil, err
	}

	postgresPort, err := strconv.Atoi(getEnv("POSTGRES_PORT", "5433"))
	if err != nil {
		return nil, err
	}

	redisPort, err := strconv.Atoi(getEnv("REDIS_PORT", "5433"))
	if err != nil {
		return nil, err
	}

	swaggerEnabled, err := strconv.ParseBool(getEnv("SWAGGER_ENABLED", "false"))
	if err != nil {
		return nil, err
	}

	config := &Config{
		NodeEnv:         getEnv("NODE_ENV", "development"),
		ApplicationPort: applicationPort,
		ApplicationURL:  os.ExpandEnv(getEnv("APPLICATION_URL", "http://localhost:4000")),
		AllowedOrigin:   getEnv("ALLOWED_ORIGIN", "http://localhost:3000"),
		SiteURL:         getEnv("SITE_URL", "http://localhost:3000"),

		SwaggerEnabled: swaggerEnabled,
		SwaggerPrefix:  getEnv("SWAGGER_PREFIX", "/swagger"),

		PostgresUser:     getEnv("POSTGRES_USER", "root"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", ""),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     postgresPort,
		PostgresDatabase: getEnv("POSTGRES_DATABASE", ""),
		PostgresURI:      os.ExpandEnv(getEnv("POSTGRES_URI", "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}")),

		RedisUser:     getEnv("REDIS_USER", "root"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     redisPort,
		RedisURI:      os.ExpandEnv(getEnv("REDIS_URI", "redis://${REDIS_PASSWORD}@${REDIS_HOST}:${REDIS_PORT}")),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
