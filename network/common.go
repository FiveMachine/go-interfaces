package network

type (
	Shapes     map[string]Shape
	HostShapes map[string]Node
)

type Node struct {
	Key string
	Id  string
}

type DVKey struct {
	Private string `yaml:"private"`
	Public  string `yaml:"public"`
}
