package store

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStore(host, port string) *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})
	return &RedisStore{client: client, ctx: context.Background()}
}

func (r *RedisStore) Push(key, value string) error {
	return r.client.LPush(r.ctx, key, value).Err()
}

func (r *RedisStore) Pop(key string) (string, error) {
	return r.client.LPop(r.ctx, key).Result()
}

func (r *RedisStore) Enqueue(key, value string) error {
	return r.client.RPush(r.ctx, key, value).Err()
}

func (r *RedisStore) Dequeue(key string) (string, error) {
	return r.client.LPop(r.ctx, key).Result()
}

func (r *RedisStore) Append(key, value string) error {
	return r.client.RPush(r.ctx, key, value).Err()
}

func (r *RedisStore) Remove(key string) (string, error) {
	return r.client.LPop(r.ctx, key).Result()
}

func (r *RedisStore) GetState(key string) ([]string, error) {
	return r.client.LRange(r.ctx, key, 0, -1).Result()
}
