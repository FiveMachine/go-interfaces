package network

type Host struct {
	Addresses  []string `yaml:"addr"`
	SSH        SSH
	PublicAddr string     `yaml:"public_addr,omitempty"` // What to do with this now
	GPS        []string   `yaml:",omitempty"`
	Shapes     HostShapes `yaml:",omitempty"`
}

type SSH struct {
	Address string `yaml:"addr"`
	Port    int
	Key     string
}
