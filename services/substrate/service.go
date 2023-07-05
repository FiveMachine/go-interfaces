package substrate

import (
	"github.com/taubyte/go-interfaces/moody"
	"github.com/taubyte/go-interfaces/services"
	"github.com/taubyte/go-interfaces/services/http"
	"github.com/taubyte/go-interfaces/services/substrate/counters"
	"github.com/taubyte/go-interfaces/services/substrate/p2p"
	smartOps "github.com/taubyte/go-interfaces/services/substrate/smartops"
	"github.com/taubyte/go-interfaces/services/tns"
	"github.com/taubyte/go-interfaces/vm"
)

type Service interface {
	services.Service
	// Http Returns the http service attached to the Substrate
	Http() http.Service
	// Vm Returns the  VM service attached to the Substrate
	Vm() vm.Service
	// Logger returns the logger for the Substrate
	Logger() moody.Logger
	// Tns returns the Tns client attached to the Substrate
	Tns() tns.Client
	// Branch returns the branch the Substrate listens to
	Branch() string
	// Counter returns the counter service attached to the Substrate s
	Counter() counters.Service
	// SmartOps returns the smartops service attached to the Substrate
	SmartOps() smartOps.Service
	// P2P returns the p2p service attached to the Substrate
	P2P() p2p.Service

	// Orbitals() []vm.Plugin
}
