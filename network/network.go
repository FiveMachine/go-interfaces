package network

import (
	"github.com/taubyte/go-interfaces/services/common"
)

type Network struct {
	Domain   Domain
	P2P      P2P
	Identity common.Identity
}

type P2P struct {
	Bootstrap map[string][]string `yaml:",omitempty"`
	Swarm     Swarm               `yaml:",omitempty"`
}

type Swarm struct {
	Key string `yaml:",omitempty"`
}

type Domain struct {
	Root       string
	Generated  string `yaml:"generated"`
	Validation Validation
}

type Validation struct {
	Key  DVKey
	Skip Skip `yaml:",omitempty"`
}

type Skip struct {
	Postfix []string `yaml:",omitempty"`
	Regex   []string `yaml:",omitempty"`
}
