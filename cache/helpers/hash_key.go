package helpers

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"time"

	json "github.com/bytedance/sonic"
	"gitlab.com/aic/aic_api/biz/util/errors"
)

// GenerateCacheKey generates a cache key for a given function name and request encoded in SHA512. With the following format:
// SHA512(functionName:marshalledRequest:softTTL:hardTTL)
// func GenerateCacheKeyWithUserId(ctx context.Context, userId string, invalidatingFunc string, softTTL, hardTTL time.Duration, version string) (string, error) {
// 	if userId == "" {
// 		return "", errors.NewInvalidParamsError("need user id")
// 	}

// 	key := constructUnhashedKey(invalidatingFunc, userId, softTTL, hardTTL, version)

// 	hasher := sha512.New()
// 	hasher.Write([]byte(key))
// 	hashedKey := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

// 	return hashedKey, nil
// }

// GenerateCacheKey generates a cache key for a given function name and request encoded in SHA512. With the following format:
// SHA512(functionName:marshalledRequest:softTTL:hardTTL)
func GenerateCacheKeyWithParams(ctx context.Context, params any, functionName string, softTTL, hardTTL time.Duration, version, env string) (string, error) {
	var (
		err           error
		marshalledReq string
	)
	if params == nil {
		return "", errors.NewInvalidParamsError("need user id")
	}
	if params == nil {
		marshalledReq = ""
	} else {
		marshalledReq, err = json.ConfigStd.MarshalToString(params) // ensures Map's keys are sorted for unique key generation
		if err != nil {
			return "", errors.NewInternalError("")
		}
	}

	key := constructUnhashedKeyParams(functionName, marshalledReq, softTTL, hardTTL, version, env)
	hasher := sha512.New()
	hasher.Write([]byte(key))
	hashedKey := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return hashedKey, nil
}

func GenerateCacheKey(prefix, id, env string) string {
	hasher := sha512.New()
	hasher.Write([]byte(fmt.Sprintf("cache-key:%s:%s:%s", prefix, id, env)))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

// func constructUnhashedKey(invalidatingFunc string, userId string, softTTL, hardTTL time.Duration, version string) string {
// 	return fmt.Sprintf("%v:%v:%d:%d:%v", invalidatingFunc, userId, int64(softTTL.Seconds()), int64(hardTTL.Seconds()), version)
// }

func constructUnhashedKeyParams(functionName string, marshalledReq string, softTTL, hardTTL time.Duration, version, env string) string {
	return fmt.Sprintf("%v:%v:%d:%d:%v:%v", functionName, marshalledReq, int64(softTTL.Seconds()), int64(hardTTL.Seconds()), version, env)
}
