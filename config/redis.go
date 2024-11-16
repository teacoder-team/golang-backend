package config

import (
	"log"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var RedisClient *redis.Client

func ConnectRedis(cfg *Config) {
	options, err := redis.ParseURL(cfg.RedisURI)
	if err != nil {
		log.Fatalf("❌ Failed to parse Redis URI: %v", err)
	}

	RedisClient = redis.NewClient(options)

	_, err = RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("❌ Failed to connect to Redis: %v", err)
	}

	log.Println("✅ Redis connected successfully")
}
