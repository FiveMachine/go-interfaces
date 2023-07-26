package http

import (
	"net/http"
	"time"

	"github.com/taubyte/go-interfaces/services/substrate/components"
	structureSpec "github.com/taubyte/go-specs/structure"
)

type Service interface {
	components.ServiceComponent
}

type Serviceable interface {
	components.Serviceable
	Handle(w http.ResponseWriter, r *http.Request, serv components.MatchDefinition) (time.Time, error)
}

type Function interface {
	Serviceable
	Config() *structureSpec.Function
}

type Website interface {
	Serviceable
	Config() *structureSpec.Website
}
