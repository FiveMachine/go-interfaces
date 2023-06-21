package smartOps

import (
	"context"

	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-interfaces/services"
	"github.com/taubyte/go-interfaces/services/substrate/common"
	"github.com/taubyte/go-interfaces/services/substrate/counters"
	"github.com/taubyte/go-interfaces/services/tns"
	"github.com/taubyte/go-interfaces/vm"
)

type Service interface {
	services.Service
	Run(caller SmartOpEventCaller, smartOpIds []string) (uint32, error)

	Dev() bool
	Verbose() bool
	Branch() string
	Context() context.Context

	Vm() vm.Service
	Tns() tns.Client

	Counter() counters.Service
}

// Caller is the serviceable that called the smartOpEvent
type SmartOpEventCaller interface {
	Context() context.Context
	Type() uint32
	Application() string
	Project() (cid.Cid, error)
}

// Util is the node utilities used by the smartOps
type Util interface {
	GPU() bool
}

type Instance interface {
	Context() context.Context
	ContextCancel()

	Run(caller SmartOpEventCaller) (uint32, error)
}

type Cache interface {
	Close()
	Get(project, application, smartOpId string, ctx context.Context) (instance Instance, ok bool)
	Put(project, application, smartOpId string, ctx context.Context, instance Instance) error
}

type Serviceable interface {
	common.Serviceable
}
