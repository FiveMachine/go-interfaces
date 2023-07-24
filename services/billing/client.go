package billing

import "github.com/taubyte/go-interfaces/services/substrate"

type Client interface {
	Close()
	List() ([]string, error)
	New(project string) (string, error)
	Report(report map[string]substrate.Metric) error
}
