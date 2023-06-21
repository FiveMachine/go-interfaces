package hoarder

import (
	"github.com/taubyte/go-interfaces/p2p/streams"
)

type Client interface {
	Rare() ([]string, error)
	Stash(cid string) (streams.Response, error)
	List() ([]string, error)
}
