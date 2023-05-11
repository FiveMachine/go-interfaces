package billing

import (
	http "bitbucket.org/taubyte/http"
	kv "bitbucket.org/taubyte/kvdb/database"
	"github.com/taubyte/go-interfaces/services"
)

type Service interface {
	services.Service
	KV() *kv.KVDatabase
	GitHubTokenHTTPAuth(ctx *http.Context) (interface{}, error)
	GitHubTokenHTTPAuthCleanup(ctx *http.Context) (interface{}, error)
}
