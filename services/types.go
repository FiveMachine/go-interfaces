package services

import (
	"github.com/taubyte/go-interfaces/kvdb"
	http "github.com/taubyte/http"
	peer "github.com/taubyte/p2p/peer"
)

type Service interface {
	Node() peer.Node
	Close() error
}

type DBService interface {
	Service
	KV() kvdb.KVDB
}

type HttpService interface {
	Service
	Http() http.Service
}

type GitHubAuth interface {
	GitHubTokenHTTPAuth(ctx http.Context) (interface{}, error)
	GitHubTokenHTTPAuthCleanup(ctx http.Context) (interface{}, error)
}
