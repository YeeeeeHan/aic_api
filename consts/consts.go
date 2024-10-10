package consts

import "time"

type ContextKey string

const (
	LogID          ContextKey    = "log-id"
	ProfileHeader  ContextKey    = "profile-header"
	DefaultSoftTTL time.Duration = 5 * time.Minute
	DefaultHardTTL time.Duration = 10 * time.Minute
)

const (
	PROD = "prod"
	TEST = "test"
)
