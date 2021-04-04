package grpc

import (
	"time"

	"github.com/go-alive/go-micro/config/source"
	proto "github.com/go-alive/go-plugins/config/source/grpc/v2/proto"
)

func toChangeSet(c *proto.ChangeSet) *source.ChangeSet {
	return &source.ChangeSet{
		Data:      c.Data,
		Checksum:  c.Checksum,
		Format:    c.Format,
		Timestamp: time.Unix(c.Timestamp, 0),
		Source:    c.Source,
	}
}
