package seer

import (
	"context"

	"github.com/taubyte/go-interfaces/kvdb"
	"github.com/taubyte/go-interfaces/services"
	"github.com/taubyte/go-interfaces/services/http"
)

type Service interface {
	services.Service
	KV() kvdb.KVDB
	Resolver() Resolver
	GitHubTokenHTTPAuth(ctx http.Context) (interface{}, error)
	GitHubTokenHTTPAuthCleanup(ctx http.Context) (interface{}, error)
	ListNodes() ([]string, error)
}

type Resolver interface {
	LookupTXT(context.Context, string) ([]string, error)
	LookupCNAME(ctx context.Context, host string) (string, error)
}
