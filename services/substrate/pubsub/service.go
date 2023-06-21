package pubsub

import (
	"context"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/taubyte/go-interfaces/services/substrate/common"
	"github.com/taubyte/go-interfaces/services/substrate/counters"
	smartOps "github.com/taubyte/go-interfaces/services/substrate/smartops"
	structureSpec "github.com/taubyte/go-specs/structure"
)

type Service interface {
	common.Service
	Subscribe(projectId, appId, channel string) error
	Publish(ctx context.Context, projectId, appId, channel string, data []byte) error
	WebSocketURL(projectId, appId, channel string) (string, error)

	Counter() counters.Service
	SmartOps() smartOps.Service
}

type Messaging interface {
	Config() *structureSpec.Messaging
}

type Serviceable interface {
	common.Serviceable
	Config() *structureSpec.Function
	HandleMessage(msg *pubsub.Message) (time.Time, error)
	Name() string
}

type Channel interface {
	Context() context.Context
	SmartOps(smartOps []string) (uint32, error)
	Type() uint32
	Messaging
}
