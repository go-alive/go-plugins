// Package memory provides a memory broker
package memory

import (
	"github.com/go-alive/go-micro/broker"
	"github.com/go-alive/go-micro/broker/nats"
	"github.com/go-alive/go-micro/config/cmd"
)

func init() {
	cmd.DefaultBrokers["nats"] = NewBroker
}

func NewBroker(opts ...broker.Option) broker.Broker {
	return nats.NewBroker(opts...)
}
