package constants

// CacheType is the type of cache to use
type CacheType int32

type CompressionLibraryType int32

const (
	// CustomCacheType allows the user to BYOC (Bring your own cache) as long as the user's cache client
	// implements the ClientAPIs interface
	CustomCacheType CacheType = iota + 1
	// RedisCacheType uses redis as a cache. It currently supports a single redis instance or a cluster.
	// Under the hood, it uses the go-redis library. The user has to pass in required attributes to connect
	// to redis itself. Supports redis 7.
	RedisCacheType
)

// RedisType is the type of redis server configuration the user is using.
type RedisType int32

const (
	// SingularRedisType is a default redis server.
	SingularRedisType RedisType = iota + 1
	// ClusterRedisType is a redis server that is set up in cluster mode.
	ClusterRedisType
)

var supportedCacheTypes = map[CacheType]bool{
	RedisCacheType:  true,
	CustomCacheType: true,
}

const (
	// NoCompression will disable compression and uncompressed values are stored in the cache.
	NoCompressionType CompressionLibraryType = iota
	// GzipCompression will enable compression and values are compressed with the GZIP library and stored in the cache.
	GzipCompressionType
	// SnappyCompression will enable compression and values are compressed with the Snappy library and stored in the cache.
	SnappyCompressionType
)
