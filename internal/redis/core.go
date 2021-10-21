package redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type RedisDB struct {
	Client *redis.Client
}

var RedisClient *RedisDB

var (
	errNil = errors.New("no matching record found in redis database.")
	ctx    = context.TODO()
)

func NewDatabase(address string) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisDB{
		Client: client,
	}, nil

}

// func SetClient() {
// 	RedisClient, _ = newDatabase("localhost:6379")
// }

// func GetClient() *RedisDB {
// 	return RedisClient
// }
