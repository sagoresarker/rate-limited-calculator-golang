package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(addr string) (*RedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}

	return &RedisStorage{client: client}, nil
}

func (s *RedisStorage) IncrementCount(windowStart time.Time, count int) error {
	key := fmt.Sprintf("ratelimit:%d", windowStart.Unix())

	ctx := context.Background()

	err := s.client.SetNX(ctx, key, count, time.Minute).Err()

	if err != nil {
		return err
	}

	return s.client.Incr(ctx, key).Err()
}
