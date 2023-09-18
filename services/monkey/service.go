package monkey

import (
	"github.com/taubyte/go-interfaces/services"
	"github.com/taubyte/go-interfaces/services/hoarder"
	"github.com/taubyte/go-interfaces/services/patrick"
)

type Service interface {
	services.Service
	Delete(jid string)
	Dev() bool

	Hoarder() hoarder.Client
	Patrick() patrick.Client
}
