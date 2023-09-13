package hoarder

import (
	"github.com/taubyte/go-interfaces/services"
	"github.com/taubyte/p2p/streams/command/response"
)

type Client interface {
	services.Client
	Rare() ([]string, error)
	Stash(cid string) (response.Response, error)
	List() ([]string, error)
}
