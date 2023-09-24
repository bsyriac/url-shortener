package store

import (
	"context"
	"os"

	"github.com/go-redis/redis"
)

var Ctx = context.Background()

func NewRedisConnection() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB_HOST"),
		Password: os.Getenv("REDIS_DB_PASS"),
		DB:       0,
	})

	return client

}
