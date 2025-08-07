package redisclient

import (
	"context"
	"log"

	"github.com/berylCAtieno/stoo-inventory/internal/config"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr: config.Config.RedisAddress,
	})

	if err := Client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}
