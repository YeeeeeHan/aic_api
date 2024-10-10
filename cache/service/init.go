package service

import (
	"github.com/pkg/errors"
	"gitlab.com/aic/aic_api/cache/constants"
	"gitlab.com/aic/aic_api/cache/helpers"
)

// Config is a configuration struct for the cache client. It allows the user to specify the type of cache they want to use
// as well as the configuration for that cache type.
type Config struct {
	// CacheProvider is the type of cache to use. These are constants in the cache package.
	CacheProvider constants.CacheType
	// CustomConfiguration is a configuration for a custom cache client. This field is required only if CacheProvider is set to
	// CustomCacheType.
	CustomConfiguration *CustomConfig
	// RedisConfiguration is a configuration for a redis cache client. This field is required only if CacheProvider is set to
	// RedisCacheType.
	RedisConfiguration *RedisConfig
}

// Validate validates the cache configuration.
func (c *Config) Validate() error {
	switch c.CacheProvider {
	case constants.CustomCacheType:
		if c.CustomConfiguration == nil {
			return helpers.TernaryOp(c.CustomConfiguration == nil, errors.Errorf("custom cache config is nil"), nil)
		}
	case constants.RedisCacheType:
		return helpers.TernaryOp(c.RedisConfiguration == nil, errors.Errorf("redis configuration is nil"), nil)
	default:
		return errors.Errorf("cache type is not supported")
	}

	return nil
}

// Freeze freezes the cache configuration and generates the respective cache clients.
func (c *Config) Freeze() (*Client, error) {
	if c == nil {
		return nil, errors.Errorf("config, is nil")
	}
	switch c.CacheProvider {
	case constants.CustomCacheType:
		return newCustom(c.CustomConfiguration)
	case constants.RedisCacheType:
		return newRedis(c.RedisConfiguration)
	default:
		return nil, errors.Errorf("cache type %d is not supported", c.CacheProvider)
	}
}
