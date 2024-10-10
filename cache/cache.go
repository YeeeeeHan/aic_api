package cache

import (
	"context"
	"time"

	json "github.com/bytedance/sonic"
	"github.com/pkg/errors"
	"gitlab.com/aic/aic_api/logs"
)

var (
	jsonAPI = json.Config{UseNumber: true}.Froze() // used to maintain int64 and float64 precision when unmarshalling into an interface{}
)

// ToggleCache is a utility function that can be called to toggle caching on / off.
// Event listeners can be configured to toggle the cache on / off with a simple configuration change.
func ToggleCache(isCacheEnabled bool) {
	skipCache = isCacheEnabled
}

type CacheValue struct {
	UpdatedTS int64
	SoftTTL   time.Duration
	Data      string
}

func getData[response any](
	ctx context.Context,
	serviceFunc func() (*response, error),
	cacheKey string,
	softTTL time.Duration,
	hardTTL time.Duration,
	readFromCache func() bool,
	writeToCache func(*response) bool,
) (
	res *response,
	err error,
) {
	if isSkipCache() {
		return serviceFunc()
	}

	var result *CacheValue
	if !readFromCache() {
		result, err = handleCacheMiss(ctx, cacheKey, serviceFunc, softTTL, hardTTL, writeToCache)
		if err != nil {
			return nil, err
		}
		return generateResp[response](result)
	}

	result, err = fetchFromCache(ctx, cacheKey)
	if err != nil {
		result, err = handleCacheMiss(ctx, cacheKey, serviceFunc, softTTL, hardTTL, writeToCache)
		if err != nil {
			return nil, err
		}
	}

	if isPastSoftTTLThreshhold(result) {
		handleCacheSoftHit(ctx, cacheKey, serviceFunc, softTTL, hardTTL, writeToCache)
	}

	return generateResp[response](result)
}

func generateResp[response any](cacheVal *CacheValue) (res *response, err error) {
	resp := new(response)
	err = generateResponseStructFromCacheVal(cacheVal, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func fetchFromCache(ctx context.Context, key string) (*CacheValue, error) {
	cacheVal := &CacheValue{}
	val, err := cacheProvider.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	err = DecompressStruct(ctx, val, cacheVal, compressionLibrary)
	return cacheVal, err
}

func handleCacheMiss[response any](ctx context.Context, key string, serviceFunc func() (*response, error), softTTL,
	hardTTL time.Duration, writeToCache func(*response) bool) (*CacheValue, error) {
	resp, err := serviceFunc()
	if err != nil {
		return nil, errors.Wrap(err, "rpc call failed")
	}
	logs.CtxInfo(ctx, "cache missed for key %s", key)
	go updateCache(ctx, key, resp, softTTL, hardTTL, writeToCache)
	return makeCacheValue(resp, softTTL)
}

func handleCacheSoftHit[response any](ctx context.Context, key string, serviceFunc func() (*response, error), softTTL,
	hardTTL time.Duration, writeToCache func(*response) bool) {
	go func() {
		resp, err := serviceFunc()
		if err != nil {
			return // don't write to cache on error
		}
		logs.CtxInfo(ctx, "cache soft hit")

		updateCache(ctx, key, resp, softTTL, hardTTL, writeToCache)
	}()
}

func generateResponseStructFromCacheVal[response any](cacheVal *CacheValue, res *response) error {
	err := jsonAPI.UnmarshalFromString(cacheVal.Data, res)
	if err != nil {
		return errors.Wrap(err, "unable to marshal cache value to rpc response")
	}
	return nil
}

func updateCache[response any](ctx context.Context, key string, rpcCallResp *response,
	softTTL, hardTTL time.Duration, writeToCache func(*response) bool) {
	if rpcCallResp == nil || !writeToCache(rpcCallResp) {
		return
	}
	cacheVal, err := makeCacheValue(rpcCallResp, softTTL)
	if err != nil {
		return
	}
	compressedData, err := CompressStruct(ctx, cacheVal, compressionLibrary)
	if err != nil {
		return
	}
	err = cacheProvider.Set(ctx, key, compressedData, hardTTL)
	if err != nil {
		return
	}
}

func isPastSoftTTLThreshhold(cacheVal *CacheValue) bool {
	return cacheVal.UpdatedTS+int64(cacheVal.SoftTTL.Seconds()) < time.Now().Unix()
}

func isSkipCache() bool {
	return skipCache
}

func makeCacheValue(val any, softTTL time.Duration) (*CacheValue, error) {
	data, err := jsonAPI.MarshalToString(val)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal rpc response")
	}
	return &CacheValue{
		UpdatedTS: time.Now().Unix(),
		Data:      data,
		SoftTTL:   softTTL,
	}, nil
}

func InvalidateCache(ctx context.Context, cacheKey string) error {
	logs.CtxInfo(ctx, "cache key: %v invalidated", cacheKey)
	return cacheProvider.Del(ctx, cacheKey)
}

func AddToCache(ctx context.Context, cacheKey string, cacheVal any) error {
	logs.CtxInfo(ctx, "adding to cache key: %v , val: %v", cacheKey, cacheVal)
	cacheVal, err := makeCacheValue(cacheVal, defaultSoftTTL)
	if err != nil {
		return err
	}
	compressedData, err := CompressStruct(ctx, cacheVal, compressionLibrary)
	if err != nil {
		return err
	}
	return cacheProvider.Set(ctx, cacheKey, compressedData, defaultHardTTL)
}

func GetMarshalledCache(ctx context.Context, cacheKey string, object any) error {
	data, err := fetchFromCache(ctx, cacheKey)
	if err != nil {
		return err
	}

	err = jsonAPI.UnmarshalFromString(data.Data, object)
	if err != nil {
		return err
	}

	return nil
}
