// Package http returns a http2 transport using net/http
package http

import (
	"github.com/go-alive/go-micro/config/cmd"
	"github.com/go-alive/go-micro/transport"
)

func init() {
	cmd.DefaultTransports["http"] = NewTransport
}

// NewTransport returns a new http transport using net/http and supporting http2
func NewTransport(opts ...transport.Option) transport.Transport {
	return transport.NewTransport(opts...)
}
