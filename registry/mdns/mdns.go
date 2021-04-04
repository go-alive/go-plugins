// Package mdns provides a multicast dns registry
package mdns

import (
	"github.com/go-alive/go-micro/config/cmd"
	"github.com/go-alive/go-micro/registry"
	"github.com/go-alive/go-micro/registry/mdns"
)

func init() {
	cmd.DefaultRegistries["mdns"] = NewRegistry
}

// NewRegistry returns a new mdns registry
func NewRegistry(opts ...registry.Option) registry.Registry {
	return registry.NewRegistry(opts...)
}

// Domain sets the mdnsDomain
func Domain(d string) registry.Option {
	return mdns.Domain(d)
}
