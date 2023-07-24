package components

import (
	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-interfaces/services/substrate"
	tns "github.com/taubyte/go-interfaces/services/tns"
	matcherSpec "github.com/taubyte/go-specs/matcher"
)

type ServiceComponent interface {
	substrate.Service
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
	Service() ServiceComponent
}

type MatchDefinition interface {
	String() string
	CachePrefix() string
}
