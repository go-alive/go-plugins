// Package memory provides an in-memory registry
package memory

import (
	"github.com/go-alive/go-micro/config/cmd"
	"github.com/go-alive/go-micro/registry"
	"github.com/go-alive/go-micro/registry/memory"
)

func init() {
	cmd.DefaultRegistries["memory"] = NewRegistry
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return memory.NewRegistry(opts...)
}
