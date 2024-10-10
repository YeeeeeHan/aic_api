package redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.Background()
)

func NewRedis(address string) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err := RedisClient.Ping(Ctx).Err(); err != nil {
		return err
	}
	return nil
}
