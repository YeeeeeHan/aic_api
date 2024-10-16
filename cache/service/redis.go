package service

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"gitlab.com/aic/aic_api/cache/constants"
)

// RedisConfig is the configuration for the redis client.
type RedisConfig struct {
	// RedisServerType is a constant that defines which redis server type the user is using.
	RedisServerType constants.RedisType
	// SingularConfig is the configuration for a singular redis server. This is required only if RedisServerType
	// is set to SingularRedisType.
	SingularConfig *redis.Options
	// ClusterConfig is the configuration for a redis cluster. This is required only if RedisServerType
	// is set to ClusterRedisType.
	ClusterConfig *redis.ClusterOptions
}

func newRedis(config *RedisConfig) (*Client, error) {
	if config == nil {
		return nil, errors.Errorf("nil ptr redis config passed in")
	}

	var (
		rdb *wrappedRedisClient
		err error
	)

	switch config.RedisServerType {
	case constants.SingularRedisType:
		if config.SingularConfig == nil {
			return nil, errors.Errorf("nil ptr redis singular config passed in")
		}

		rdb, err = newRedisSingular(config.SingularConfig)
	case constants.ClusterRedisType:
		if config.ClusterConfig == nil {
			return nil, errors.Errorf("nil ptr redis cluster config passed in")
		}
		rdb, err = newRedisCluster(config.ClusterConfig)
	}

	return &Client{
		GetAPI: rdb,
		SetAPI: rdb,
	}, err
}

func newRedisCluster(cfg *redis.ClusterOptions) (*wrappedRedisClient, error) {
	rdb := redis.NewClusterClient(cfg)
	err := rdb.ForEachShard(context.Background(), func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to reach a shard")
	}

	return &wrappedRedisClient{client: rdb}, nil
}

func newRedisSingular(cfg *redis.Options) (*wrappedRedisClient, error) {
	rdb := redis.NewClient(cfg)
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, errors.Wrap(err, "unable to successfully connect to redis")
	}
	return &wrappedRedisClient{client: rdb}, nil
}

type wrappedRedisClient struct {
	client redis.UniversalClient
}

func (c *wrappedRedisClient) Get(ctx context.Context, key string) ([]byte, error) {
	return c.client.Get(ctx, key).Bytes()
}

func (c *wrappedRedisClient) Set(ctx context.Context, key string, val any, ttl time.Duration) error {
	return c.client.Set(ctx, key, val, ttl).Err()
}
