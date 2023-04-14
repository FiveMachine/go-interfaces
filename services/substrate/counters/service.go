package counters

import (
	"context"

	"github.com/taubyte/go-interfaces/services"
)

type Service interface {
	services.Service
	Context() context.Context
	Push(...*WrappedMetric)
}
