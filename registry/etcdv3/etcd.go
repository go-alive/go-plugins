// Package etcd provides an etcd v3 service registry
package etcdv3

import (
	"github.com/go-alive/go-micro/config/cmd"
	"github.com/go-alive/go-micro/registry"
	"github.com/go-alive/go-micro/registry/etcd"
)

func init() {
	cmd.DefaultRegistries["etcdv3"] = etcd.NewRegistry
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return etcd.NewRegistry(opts...)
}
