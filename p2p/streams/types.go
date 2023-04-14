package streams

import (
	"io"

	"github.com/libp2p/go-libp2p/core/network"
)

type Connection interface {
	io.Closer
	network.ConnSecurity
	network.ConnMultiaddrs
}

type Command interface {
	Connection() (Connection, error)
	Encode(io.Writer) error
}

type Response interface {
	Encode(s io.Writer) error
	Get(string) (interface{}, error)
}
