package redis

import (
	"github.com/go-redis/redis/v8"
)

type RedisCore struct {
}

var Client *redis.Client

func (r *RedisCore) SetRedisClient() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
