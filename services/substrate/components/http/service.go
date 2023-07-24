package http

import (
	goHttp "net/http"
	"time"

	"github.com/taubyte/go-interfaces/services/substrate/components"
	structureSpec "github.com/taubyte/go-specs/structure"
)

type Service interface {
	components.ServiceComponent
}

type Serviceable interface {
	components.Serviceable
	Handle(w goHttp.ResponseWriter, r *goHttp.Request, serv components.MatchDefinition) (time.Time, error)
}

type Function interface {
	Serviceable
	Config() *structureSpec.Function
}

type Website interface {
	Serviceable
	Config() *structureSpec.Website
}
