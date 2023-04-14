package p2p

import (
	"context"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/taubyte/go-interfaces/p2p/streams"
	"github.com/taubyte/go-interfaces/services/substrate/common"
	"github.com/taubyte/go-interfaces/services/substrate/counters"
	smartOps "github.com/taubyte/go-interfaces/services/substrate/smartops"
	structureSpec "github.com/taubyte/go-specs/structure"
)

type Command interface {
	Send(ctx context.Context, body map[string]interface{}) (streams.Response, error)
	SendTo(ctx context.Context, cid cid.Cid, body map[string]interface{}) (streams.Response, error)
}

type Stream interface {
	Listen() (protocol string, err error)
	Command(command string) (Command, error)
	Close()
}

type StreamHandler func(cmd streams.Command) (resp streams.Response, err error)

type CommandService interface {
	Close()
}

type MatchDefinition struct {
	Project     string
	Application string
	Protocol    string
	Command     string
}

func (m *MatchDefinition) String() string {
	return m.Project + m.Application + m.Protocol + m.Command
}

func (m *MatchDefinition) CachePrefix() string {
	return m.Project
}

type Service interface {
	common.Service
	Stream(ctx context.Context, projectID, applicationID, protocol string) (Stream, error)
	StartStream(name, protocol string, handler StreamHandler) (CommandService, error)
	SmartOps() smartOps.Service
	Counter() counters.Service
	LookupService(matcher *MatchDefinition) (config *structureSpec.Service, application string, err error)
	Discover(ctx context.Context, max int, timeout time.Duration) ([]peer.AddrInfo, error)
}

type Serviceable interface {
	common.Serviceable
	Handle(data streams.Command) (time.Time, streams.Response, error)
	Name() string
	Close()
}
