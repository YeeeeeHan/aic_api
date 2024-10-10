// Code generated by hertz generator.

package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	redisv8 "github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gitlab.com/aic/aic_api/biz/dal"
	"gitlab.com/aic/aic_api/biz/handler"
	"gitlab.com/aic/aic_api/biz/redis"
	"gitlab.com/aic/aic_api/biz/router"
	"gitlab.com/aic/aic_api/cache"
	"gitlab.com/aic/aic_api/cache/constants"
	"gitlab.com/aic/aic_api/cache/service"
	"gitlab.com/aic/aic_api/consts"
	"gitlab.com/aic/aic_api/logs"
	"gitlab.com/aic/aic_api/monitoring"
)

type CustomClient struct {
	RedisClient *redisv8.Client
}

func (c *CustomClient) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := c.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

func (c *CustomClient) Set(ctx context.Context, key string, val any, ttl time.Duration) error {
	err := c.RedisClient.Set(ctx, key, val, ttl).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		logs.CtxError(ctx, "error setting value: %v", err)
		return err
	}
	return nil
}

func (c *CustomClient) Del(ctx context.Context, key string) error {
	err := c.RedisClient.Del(ctx, key).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		logs.CtxError(ctx, "error setting value: %v", err)
		return err
	}
	return nil
}

var environment string

func main() {
	h := server.Default()

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal("Fatal error when retrieving env variables", err)
		os.Exit(1)
	}

	register(h)


	if err := dal.InitDB(isTestEnv()); err != nil {
		panic(err)
	}

	//Initialises log writer
	logWriter := monitoring.InitWriter()
	go logWriter.Run()

	RedisAddr := os.Getenv("REDISADDR")
	err = redis.NewRedis(RedisAddr)
	if err != nil {
		logrus.Fatal("Fatal error when initialising redis", err)
		panic(err)
	}

	CustomClient := &CustomClient{}
	CustomClient.RedisClient = redis.RedisClient

	cacheConfig := &cache.Config{ // Configuration used to initialise cache
		DefaultSoftTTL: consts.DefaultSoftTTL, // Global softTTL if one is not provided at the time of function call. (Refer to concepts and notes for more info)
		DefaultHardTTL: consts.DefaultHardTTL, // Global hardTTL if one is not provided at the time of functon call. (Refer to concepts and notes for more info)
		CacheConfig: service.Config{ // Cache configuration, this must be provided.
			CacheProvider: constants.CustomCacheType, // Choose cache provider as CustomCacheType
			CustomConfiguration: &service.CustomConfig{ // Custom configuration is required if CacheProvider is set to CustomConfiguration
				Client: CustomClient, // Custom client that implements Get and Set
			},
		},
		SkipCache: false,    // Global toggle on whether or not to skip cache. Useful for toggling between environments.
		Version:   "v1.0.0", // Version number set. If there are any upgrades, this prevents breaking changes as old keys will not be re-used.
		Env:       environment,
	}

	if err := cache.Init(cacheConfig); err != nil {
		panic(err) // panic if caching fails to initialise
	}

	h.Spin()
}

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
}

// register registers all routers.
func register(r *server.Hertz) {

	router.GeneratedRegister(r)

	customizedRegister(r)
}

func isTestEnv() bool {
	env := flag.String("env", "test", "Server environment")
	flag.Parse()

	if *env == "test" {
		environment = consts.TEST
		logrus.Info("environment set to test")
		return true
	} else if *env == "prod" {
		environment = consts.PROD
		logrus.Info("environment set to prod")
		return false
	} else {
		logrus.Info("invalid environment, defaulting to test environment")
	}
	environment = consts.TEST

	return true
}
