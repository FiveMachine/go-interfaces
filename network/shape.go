package network

type Shape struct {
	Protocols []string
	Ports     map[string]int `yaml:",omitempty"`
}
