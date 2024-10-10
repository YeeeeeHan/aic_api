package service

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"gitlab.com/aic/aic_api/cache/constants"
// )

// var (
// 	sampleInvalidCustomCacheConfig = &Config{
// 		CacheProvider: constants.CustomCacheType,
// 	}
// 	sampleValidCustomCacheConfig = &Config{
// 		CacheProvider:       constants.CustomCacheType,
// 		CustomConfiguration: &CustomConfig{Client: &testCustomCache{}},
// 	}
// )

// func TestCacheInit(t *testing.T) {
// 	c, _ := newCustom(sampleValidCustomCacheConfig.CustomConfiguration)
// 	tests := []struct {
// 		name          string
// 		config        *Config
// 		validateError bool
// 		client        *Client
// 		freezeError   bool
// 	}{
// 		{
// 			name:          "invalid custom cache config",
// 			config:        sampleInvalidCustomCacheConfig,
// 			validateError: true,
// 			client:        nil,
// 			freezeError:   true,
// 		}, {
// 			name:          "valid custom cache config",
// 			config:        sampleValidCustomCacheConfig,
// 			validateError: false,
// 			client:        c,
// 			freezeError:   false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := tt.config.Validate()
// 			assert.Equal(t, tt.validateError, err != nil)

// 			client, err := tt.config.Freeze()
// 			assert.Equal(t, tt.freezeError, err != nil)
// 			assert.Equal(t, tt.client, client)
// 		})
// 	}
// }

// type testCustomCache struct{}

// func (c *testCustomCache) Get(ctx context.Context, key string) ([]byte, error) {
// 	return nil, nil
// }
// func (c *testCustomCache) Set(ctx context.Context, key string, val any, ttl time.Duration) error {
// 	return nil
// }
