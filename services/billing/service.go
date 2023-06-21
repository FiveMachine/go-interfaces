package billing

import (
	"github.com/taubyte/go-interfaces/kvdb"
	"github.com/taubyte/go-interfaces/services"
	"github.com/taubyte/go-interfaces/services/http"
)

type Service interface {
	services.Service
	KV() kvdb.KVDB
	GitHubTokenHTTPAuth(ctx http.Context) (interface{}, error)
	GitHubTokenHTTPAuthCleanup(ctx http.Context) (interface{}, error)
}
