package network

type Host struct {
	Address    string `yaml:"addr"`
	PublicAddr string `yaml:"public_addr,omitempty"`
	Port       int
	Key        string
	GPS        []string   `yaml:",omitempty"`
	Shapes     HostShapes `yaml:",omitempty"`
}
