// Package service provides the broker service client
package service

import (
	"github.com/go-alive/go-micro/broker"
	"github.com/go-alive/go-micro/broker/service"
	"github.com/go-alive/go-micro/config/cmd"
)

func init() {
	cmd.DefaultBrokers["service"] = NewBroker
}

func NewBroker(opts ...broker.Option) broker.Broker {
	return service.NewBroker(opts...)
}
