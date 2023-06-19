package network

type KeyConfig struct {
	User  string
	Files []string // TODO: should only change to a single file?
}

// TODO: delete this as duplicate as above?
type Key struct {
	User  string
	Pem   []string `yaml:"files"`
	Value string
}

type KeyFile struct {
	File     string
	Data     []byte
	FileName string
}
