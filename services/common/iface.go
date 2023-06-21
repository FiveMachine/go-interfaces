package common

import (
	"context"

	"github.com/taubyte/go-interfaces/services"
)

type Package interface {
	Config() *GenericConfig
	New(context.Context, *GenericConfig) (services.Service, error)
}
