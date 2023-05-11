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

	Name() string
	Get(string) (interface{}, bool)
	Set(key, value string)
	Delete(string)
	Raw() map[string]interface{}
}

type Response interface {
	Encode(s io.Writer) error
	Get(string) (interface{}, error)
}
