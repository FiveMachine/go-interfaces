package console

import "github.com/taubyte/go-interfaces/p2p/streams"

type Injection func(streams.Body)

type Client interface {
	Close()
}
