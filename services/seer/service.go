package seer

import (
	"context"

	http "bitbucket.org/taubyte/http"
	kv "bitbucket.org/taubyte/kvdb/database"
	"github.com/taubyte/go-interfaces/services"
)

type Service interface {
	services.Service
	KV() *kv.KVDatabase
	Resolver() Resolver
	GitHubTokenHTTPAuth(ctx *http.Context) (interface{}, error)
	GitHubTokenHTTPAuthCleanup(ctx *http.Context) (interface{}, error)
	ListNodes() ([]string, error)
}

type Resolver interface {
	LookupTXT(context.Context, string) ([]string, error)
	LookupCNAME(ctx context.Context, host string) (string, error)
}
