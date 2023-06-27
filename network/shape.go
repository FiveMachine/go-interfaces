package network

type Shape struct {
	Protocols []string `yaml:",omitempty"`
	Ports     Ports    `yaml:",omitempty"`
}

type Ports struct {
	Main map[string]int
	Lite map[string]int
}
