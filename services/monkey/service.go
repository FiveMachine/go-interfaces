package monkey

import (
	"github.com/taubyte/go-interfaces/services"
	"github.com/taubyte/go-interfaces/services/patrick"
)

type Service interface {
	services.Service
	Patrick() patrick.Client
	Delete(jid string)

	Dev() bool
}
