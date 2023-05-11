package billing

import (
	kv "bitbucket.org/taubyte/kvdb/database"
	"github.com/taubyte/go-interfaces/services"
	http "github.com/taubyte/http"
)

type Service interface {
	services.Service
	KV() *kv.KVDatabase
	GitHubTokenHTTPAuth(ctx http.Context) (interface{}, error)
	GitHubTokenHTTPAuthCleanup(ctx http.Context) (interface{}, error)
}
