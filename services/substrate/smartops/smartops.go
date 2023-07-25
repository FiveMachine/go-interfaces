package smartops

import (
	"context"

	"github.com/ipfs/go-cid"
)

type EventCaller interface {
	Context() context.Context
	Type() uint32
	Application() string
	Project() (cid.Cid, error)
}

type Instance interface {
	Context() context.Context
	ContextCancel()

	Run(caller EventCaller) (uint32, error)
}

type SmartOpsCache interface {
	Close()
	Get(project, application, smartOpId string, ctx context.Context) (instance Instance, ok bool)
	Put(project, application, smartOpId string, ctx context.Context, instance Instance) error
}

// Util is the node utilities used by the smartOps
type Util interface {
	GPU() bool
}
