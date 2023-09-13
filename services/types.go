package services

import (
	"context"

	peer "github.com/taubyte/p2p/peer"
	streams "github.com/taubyte/p2p/streams/client"
)

type Service interface {
	Node() peer.Node
	Close() error
}

type Client interface {
	Close()
	Context() context.Context
	New(cmd string, opts ...streams.Option) *streams.Request
}
