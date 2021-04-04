// Package memory provides a memory broker
package memory

import (
	"github.com/go-alive/go-micro/broker"
	"github.com/go-alive/go-micro/broker/memory"
	"github.com/go-alive/go-micro/config/cmd"
)

func init() {
	cmd.DefaultBrokers["memory"] = NewBroker
}

func NewBroker(opts ...broker.Option) broker.Broker {
	return memory.NewBroker(opts...)
}
