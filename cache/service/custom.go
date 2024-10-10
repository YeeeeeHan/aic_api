package service

// import (
// 	"github.com/pkg/errors"
// )

// // ClientAPIs consolidates the APIs that a custom cache client must implement.
// type ClientAPIs interface {
// 	// IGet is an interface for all cache clients that support Get operations.
// 	IGet
// 	// ISet is an interface for all cache clients that support Set operations.
// 	ISet
// }

// // CustomConfig is a configuration struct for a custom cache client.
// type CustomConfig struct {
// 	Client ClientAPIs
// }

// func newCustom(cfg *CustomConfig) (*Client, error) {
// 	if cfg == nil {
// 		return nil, errors.Errorf("nil ptr passed in for custom cache config")
// 	}

// 	return &Client{
// 		GetAPI: cfg.Client,
// 		SetAPI: cfg.Client,
// 	}, nil
// }
