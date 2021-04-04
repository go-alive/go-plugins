package etcd

import (
	"github.com/go-alive/go-micro/registry"
	"github.com/go-alive/go-micro/registry/etcd"
)

// Auth allows you to specify username/password
func Auth(username, password string) registry.Option {
	return etcd.Auth(username, password)
}
