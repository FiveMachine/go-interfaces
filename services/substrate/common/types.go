package common

import (
	"context"

	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-interfaces/services"
	"github.com/taubyte/go-interfaces/services/substrate/counters"
	tns "github.com/taubyte/go-interfaces/services/tns"
	"github.com/taubyte/go-interfaces/vm"
	matcherSpec "github.com/taubyte/go-specs/matcher"
)

type Service interface {
	services.Service
	Dev() bool
	Vm() vm.Service
	Verbose() bool
	Branch() string
	Tns() tns.Client
	Context() context.Context
	CheckTns(MatchDefinition) ([]Serviceable, error)
	Cache() Cache
}

type Cache interface {
	Add(Serviceable) (Serviceable, error)
	Get(MatchDefinition) ([]Serviceable, error)
	Remove(Serviceable)
	Validate(serviceables []Serviceable, branch string, tns tns.Client) error
	Close()
}

type Serviceable interface {
	Match(MatchDefinition) matcherSpec.Index
	Validate(MatchDefinition) error
	Project() (cid.Cid, error)
	Commit() string
	Matcher() MatchDefinition
	Id() string
	Ready() error
	Close()
	Counter() counters.Service
	Service() Service
}

type MatchDefinition interface {
	String() string
	CachePrefix() string
}
