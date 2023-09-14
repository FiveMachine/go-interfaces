package monkey

import (
	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-interfaces/services"
	"github.com/taubyte/p2p/streams/command/response"
)

type Client interface {
	services.Client
	Status(jid string) (*StatusResponse, error)
	Update(jid string, body map[string]interface{}) (string, error)
	List() ([]string, error)
	Cancel(cid cid.Cid, jid string) (response.Response, error)
}
