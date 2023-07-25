package billing

import "github.com/taubyte/go-interfaces/services/substrate/counters"

type Client interface {
	Close()
	List() ([]string, error)
	New(project string) (string, error)
	Report(report map[string]counters.Metric) error
}
