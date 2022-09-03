package database

import (
	"context"
	"eaza-go/internal/common"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"time"
)

var RedisClient *Redis

type Redis struct {
	client *redis.Client
}

func NewRedisClient(opts *redis.Options) fiber.Storage {
	client := redis.NewClient(opts)
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	if RedisClient == nil {
		RedisClient = &Redis{
			client,
		}
	}
	return RedisClient
}

func (r *Redis) Get(key string) ([]byte, error) {
	res := r.client.Get(context.Background(), key)
	val, err := res.Bytes()
	if err != nil {
		return nil, err
	}
	return val, err
}

// DefaultSet sets key and value expire in  time.Hour * 24 * 30 time
func (r *Redis) DefaultSet(key string, val []byte) error {
	return r.Set(key, val, common.RedisCacheExpiration)
}

func (r *Redis) Set(key string, val []byte, exp time.Duration) error {
	return r.client.Set(context.Background(), key, val, exp).Err()
}

func (r *Redis) Delete(key string) error {
	return r.client.Del(context.Background(), key).Err()

}

func (r *Redis) Reset() error {
	return r.client.FlushAll(context.Background()).Err()
}

func (r *Redis) Close() error {
	return r.client.Close()
}
