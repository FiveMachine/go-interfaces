package network

type ParsedNetwork struct {
	Hosts   map[string]Host
	Keys    map[string]Key
	Network Network
	Shapes  Shapes
}

type Key struct {
	User  string
	Pem   []string `yaml:"files"`
	Value string
}
