// Package quic provides a QUIC based transport
package quic

import (
	"github.com/go-alive/go-micro/config/cmd"
	"github.com/go-alive/go-micro/transport"
	"github.com/go-alive/go-micro/transport/quic"
)

func init() {
	cmd.DefaultTransports["quic"] = NewTransport
}

func NewTransport(opts ...transport.Option) transport.Transport {
	return quic.NewTransport(opts...)
}
