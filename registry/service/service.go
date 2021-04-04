// Package service uses the registry service
package service

import (
	"github.com/go-alive/go-micro/config/cmd"
	"github.com/go-alive/go-micro/registry"
	"github.com/go-alive/go-micro/registry/service"
)

func init() {
	cmd.DefaultRegistries["service"] = NewRegistry
}

// NewRegistry returns a new registry service client
func NewRegistry(opts ...registry.Option) registry.Registry {
	return service.NewRegistry(opts...)
}
