package cache

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/aic/aic_api/cache/helpers"
)

// This function is invalidate caching, such as when profile is updated. It uses the global default set hard and soft TTLs.
func ServiceCallWithInvalidation[params, response any](serviceFunc func(ctx context.Context, params *params) (*response, error), par *params, ctx context.Context, cacheKey ...string) (*response, error) {
	for _, key := range cacheKey {
		err := InvalidateServiceCallCache(ctx, key)
		if err != nil {
			return nil, err
		}
	}
	return wrapServiceCallFunc(serviceFunc, par, ctx)()
}

// This function is to cache functions by user id, where invalidation is required. It uses the global default set hard and soft TTLs.
func ServiceCallWithCacheKey[params, response any](serviceFunc func(ctx context.Context, params *params) (*response, error), par *params, ctx context.Context, cacheKey string, softTTL, hardTTL time.Duration) (*response, error) {
	return getData(ctx, wrapServiceCallFunc(serviceFunc, par, ctx), cacheKey, softTTL,
		hardTTL, func() bool { return true }, func(resp *response) bool { return true })
}

// This function is to cache functions by parameters, where invalidation is not required
func ServiceCallWithParams[params, response any](serviceFunc func(ctx context.Context, params *params) (*response, error), par *params, ctx context.Context) (*response, error) {
	return ServiceCallWithTTL(serviceFunc, par, ctx, defaultSoftTTL, defaultHardTTL)
}

// It uses user defined hard and soft TTLs.
func ServiceCallWithTTL[params, response any](serviceFunc func(ctx context.Context, par *params) (*response, error), par *params, ctx context.Context, softTTL, hardTTL time.Duration) (*response, error) {
	if serviceFunc == nil {
		return nil, errors.Errorf("serviceFunc is nil")
	}

	serviceFuncName := helpers.GetFunctionName(serviceFunc)

	var cacheKey string
	var err error

	cacheKey, err = helpers.GenerateCacheKeyWithParams(ctx, par, serviceFuncName, softTTL, hardTTL, version, Env)
	if err != nil {
		return nil, err
	}

	return getData(ctx, wrapServiceCallFunc(serviceFunc, par, ctx), cacheKey, softTTL,
		hardTTL, func() bool { return true }, func(resp *response) bool { return true })
}

func wrapServiceCallFunc[params, response any](serviceFunc func(ctx context.Context, par *params) (*response, error), par *params, ctx context.Context) func() (*response, error) {
	return func() (*response, error) {
		return serviceFunc(ctx, par)
	}
}

func InvalidateServiceCallCache(ctx context.Context, cacheKey string) error {
	return InvalidateCache(ctx, cacheKey)
}
