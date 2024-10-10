package cache

import (
	"time"

	"gitlab.com/aic/aic_api/cache/constants"
	"gitlab.com/aic/aic_api/cache/service"
)

var (
	defaultSoftTTL time.Duration
	defaultHardTTL time.Duration

	cacheProvider *service.Client

	skipCache bool

	compressionLibrary constants.CompressionLibraryType

	version string

	Env string
)

func InjectCacheProvider(c *service.Client) {
	cacheProvider = c
}

func InjectSkipCache(b bool) {
	skipCache = b
}

func InjectSoftTTL(softTTL time.Duration) {
	defaultSoftTTL = softTTL
}
func InjectHardTTL(hardTTL time.Duration) {
	defaultHardTTL = hardTTL
}
func InjectCompressionLibrary(t constants.CompressionLibraryType) {
	compressionLibrary = t
}
func InjectVersion(v string) {
	version = v
}
func InjectEnv(e string) {
	Env = e
}
