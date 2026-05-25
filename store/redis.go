package store

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx = context.Background()

func Init(host, port string) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})
}

func Push(key, value string) error {
	return client.LPush(ctx, key, value).Err()
}

func Pop(key string) (string, error) {
	return client.LPop(ctx, key).Result()
}

func Enqueue(key, value string) error {
	return client.RPush(ctx, key, value).Err()
}

func Dequeue(key string) (string, error) {
	return client.LPop(ctx, key).Result()
}

func Append(key, value string) error {
	return client.RPush(ctx, key, value).Err()
}

func Remove(key string) (string, error) {
	return client.LPop(ctx, key).Result()
}

func GetState(key string) ([]string, error) {
	return client.LRange(ctx, key, 0, -1).Result()
}
