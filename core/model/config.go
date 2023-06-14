package model

import "time"

// ApplicationConfig ...
type ApplicationConfig struct {
	AppName          string
	GRPCTimeout      time.Duration
	CacheExpiry      time.Duration
	CacheCleanup     time.Duration
	DefaultPageLimit int
}
