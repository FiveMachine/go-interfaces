package http

import (
	goHttp "net/http"
	"time"

	"github.com/taubyte/go-interfaces/services/http"
	"github.com/taubyte/go-interfaces/services/substrate/common"
	"github.com/taubyte/go-interfaces/services/substrate/counters"
	smartOps "github.com/taubyte/go-interfaces/services/substrate/smartops"
)

type Service interface {
	common.Service
	Http() http.Service

	Counter() counters.Service
	SmartOps() smartOps.Service
}

type Serviceable interface {
	common.Serviceable
	Handle(w goHttp.ResponseWriter, r *goHttp.Request, serv common.MatchDefinition) (time.Time, error)
}
