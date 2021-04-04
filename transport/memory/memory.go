// Package memory is an in-memory transport
package memory

import (
	"github.com/go-alive/go-micro/config/cmd"
	"github.com/go-alive/go-micro/transport"
	"github.com/go-alive/go-micro/transport/memory"
)

func init() {
	cmd.DefaultTransports["memory"] = NewTransport
}

func NewTransport(opts ...transport.Option) transport.Transport {
	return memory.NewTransport(opts...)
}
