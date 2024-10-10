package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/aic/aic_api/cache/constants"
	"gitlab.com/aic/aic_api/logs"
)

// Client is a generic client structure for caches. All caches must fullfil the APIs that Client provides.
// Client is structured in this way to allow for easy extension of the cache package as well as ease of mocking.
type Client struct {
	// GetAPI is any cache client that can get items from the cache.
	GetAPI IGet
	// SetAPI is any cache client that can set items
	SetAPI ISet
	// DelAPI is any cache client that can del items
	DelAPI IDel
}

var CompressionLibrary constants.CompressionLibraryType

// IGet is an interface for all cache clients that support Get operations.
type IGet interface {
	Get(ctx context.Context, key string) ([]byte, error)
}

// ISet is an interface for all cache clients that support Set operations.
type ISet interface {
	Set(ctx context.Context, key string, val any, ttl time.Duration) error
}

// ISet is an interface for all cache clients that support Set operations.
type IDel interface {
	Del(ctx context.Context, key string) error
}

// Get simply gets an item from the cache based on the API provided by the cache client.
func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	compressedData, err := c.GetAPI.Get(ctx, key)
	logs.CtxInfo(ctx, "getting from cache")
	if err != nil {
		return nil, errors.Wrap(err, "unable to pull from cache")
	}

	return compressedData, nil
}

// Set simply sets an item in the cache based on the API provided by the cache client.
func (c *Client) Set(ctx context.Context, key string, val any, ttl time.Duration) error {
	err := c.SetAPI.Set(ctx, key, val, ttl)
	logs.CtxInfo(ctx, "cache set for key: %v", key)
	if err != nil {
		return errors.Wrap(err, "unable to set into cache")
	}
	return nil
}

// Del simply deletes an item in the cache based on the API provided by the cache client.
func (c *Client) Del(ctx context.Context, key string) error {
	err := c.DelAPI.Del(ctx, key)
	if err != nil {
		return errors.Wrap(err, "unable to del cache")
	}
	return nil
}

// ClientAPIs consolidates the APIs that a custom cache client must implement.
type ClientAPIs interface {
	// IGet is an interface for all cache clients that support Get operations.
	IGet
	// ISet is an interface for all cache clients that support Set operations.
	ISet
	// IDel is an interface for all cache clients that support Del operations.
	IDel
}

// CustomConfig is a configuration struct for a custom cache client.
type CustomConfig struct {
	Client ClientAPIs
}

func newCustom(cfg *CustomConfig) (*Client, error) {
	if cfg == nil {
		return nil, errors.Errorf("nil ptr passed in for custom cache config")
	}

	return &Client{
		GetAPI: cfg.Client,
		SetAPI: cfg.Client,
		DelAPI: cfg.Client,
	}, nil
}
