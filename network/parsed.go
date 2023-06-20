package network

type ParsedNetwork struct {
	Hosts   map[string]Host
	Keys    map[string]Key
	Network Network
	Shapes  Shapes
}
