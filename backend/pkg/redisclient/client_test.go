package redisclient

import (
	"context"
	"testing"
)

func TestRedisConnection(t *testing.T) {

	InitRedis()

	err := Client.Ping(context.Background()).Err()
	if err != nil {
		t.Fatalf("Redis connection failed: %v", err)
	}
}
