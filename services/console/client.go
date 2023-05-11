package console

import (
	"bitbucket.org/taubyte/p2p/streams/command"
)

type Injection func(command.Body)

type Client interface {
	Close()
}
