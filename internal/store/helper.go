package store

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	rdb "github.com/call-stack/copy_store.git/internal/redis"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Helper struct{}

func (h *Helper) getUniqueHash(ip string) string {
	ct_nano := time.Now().UnixNano()
	data_to_hash := ip + fmt.Sprint(ct_nano)
	md5_encoded := md5.Sum([]byte(data_to_hash))
	hash := base64.URLEncoding.EncodeToString(md5_encoded[:])[:URL_LENGTH]

	return hash
}

func (h *Helper) isHashPresentInRedis(ctx context.Context, hash string) (bool, string) {
	value, err := rdb.Client.Get(ctx, hash).Result()
	if err == redis.Nil {
		return false, ""
	} else if err != nil {
		panic(err)
	} else {
		return true, value
	}

}

func (h *Helper) setHashInRedis(ctx context.Context, hash string) {
	err := rdb.Client.Set(ctx, hash, 1, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (h *Helper) storeInMongo(ctx context.Context, bson_content []byte) {
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27018"))
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("content").Collection("data_store")
	_, insertErr := collection.InsertOne(ctx, bson_content)
	if insertErr != nil {
		log.Println("Error: ", insertErr)
		log.Fatal()
	}

}
