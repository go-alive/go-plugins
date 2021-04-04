// Package grpc provides a grpc transport
package grpc

import (
	"github.com/go-alive/go-micro/config/cmd"
	"github.com/go-alive/go-micro/transport"
	"github.com/go-alive/go-micro/transport/grpc"
)

func init() {
	cmd.DefaultTransports["grpc"] = NewTransport
}

func NewTransport(opts ...transport.Option) transport.Transport {
	return grpc.NewTransport(opts...)
}
