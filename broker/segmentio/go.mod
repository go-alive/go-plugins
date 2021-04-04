module github.com/go-alive/go-plugins/broker/segmentio

go 1.13

require (
	github.com/google/uuid v1.1.1
	github.com/go-alive/go-micro/v2 v2.9.1
	github.com/go-alive/go-plugins/broker/kafka/v2 v2.3.0
	github.com/go-alive/go-plugins/codec/segmentio/v2 v2.3.0
	github.com/segmentio/kafka-go v0.3.7
)

replace github.com/go-alive/go-plugins/codec/segmentio/v2 => ../../codec/segmentio
