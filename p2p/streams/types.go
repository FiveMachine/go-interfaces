package streams

import (
	"io"

	"github.com/libp2p/go-libp2p/core/network"
)

type Body map[string]interface{}

type Stream network.Stream
type StreamHandler func(Stream)

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
	SetName(interface{}) error
	Delete(string)
	Raw() map[string]interface{}
}

type Response interface {
	Encode(s io.Writer) error
	Get(string) (interface{}, error)
	Set(key string, data interface{})
}
