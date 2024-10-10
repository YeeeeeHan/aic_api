package service

import (
	"github.com/go-redis/redis/v8"
	"gitlab.com/aic/aic_api/cache/constants"
)

var (
	sampleSingularRedisInstanceConfig = &RedisConfig{
		RedisServerType: constants.SingularRedisType,
		SingularConfig: &redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default DB

		},
	}
)

type TestValue struct {
	Hello string `json:"hello"`
}
