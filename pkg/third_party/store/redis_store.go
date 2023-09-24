package store

import (
	"fmt"

	"github.com/go-redis/redis"
)

type RedisStore interface {
	Create()
}

type RedisStruct struct {
	Client *redis.Client
}

func (r *RedisStruct) Create() {

	fmt.Println(r.Client.Ping())
}
