package common

import (
	"github.com/taubyte/go-interfaces/p2p/peer"
	"github.com/taubyte/go-interfaces/services/http"
	seerIface "github.com/taubyte/go-interfaces/services/seer"
)

type GenericConfig struct {
	Shape        string
	Node         peer.Node
	Http         http.Service
	ClientNode   peer.Node
	DVPrivateKey []byte
	DVPublicKey  []byte

	Root string

	Protocols []string

	PrivateKey []byte `yaml:"privatekey"`
	SwarmKey   []byte `yaml:"swarmkey"`

	Ports map[string]int

	P2PListen   []string `yaml:"p2p-listen"`
	P2PAnnounce []string `yaml:"p2p-announce"`

	Location *seerIface.Location `yaml:"location"`

	Bootstrap bool
	Peers     []string `yaml:"peers"`
	DevMode   bool     `yaml:"devmode"`

	HttpListen string `yaml:"http-listen"`
	HttpSecure bool   `yaml:"http-secure"`
	NetworkUrl string `yaml:"network-url"`
	DnsPort    int    `yaml:"dns"`

	Verbose bool   `yaml:"verbose"`
	Branch  string `yaml:"branch"`

	TLS     TLSConfig      `yaml:"tls"`
	Domains HttpDomainInfo `yaml:"domains"`

	Plugins []string
}

type TLSConfig struct {
	Certificate string `yaml:"certificate"`
	Key         string `yaml:"key"`
}

type HttpDomainInfo struct {
	Key         DVKey              `yaml:"key"`
	Whitelisted WhiteListedDomains `yaml:"whitelist"`
	Services    string             `yaml:"services"`
	Generated   string             `yaml:"generated"`
}

// TODO: combine with the structure inside mycelium
type WhiteListedDomains struct {
	Postfix []string
	Regex   []string
}

type DVKey struct {
	Private string `yaml:"private"`
	Public  string `yaml:"public"`
}
