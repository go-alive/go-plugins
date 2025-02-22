// Package cache provides a registry cache
package cache

import (
	"github.com/go-alive/go-micro/registry"
	"github.com/go-alive/go-micro/registry/cache"
)

// New returns a new cache
func New(r registry.Registry, opts ...cache.Option) cache.Cache {
	return cache.New(r, opts...)
}
