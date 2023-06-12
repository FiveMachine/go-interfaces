package database

import (
	"context"

	"github.com/taubyte/go-interfaces/services"
	smartOps "github.com/taubyte/go-interfaces/services/substrate/smartops"
)

type Service interface {
	services.Service
	Database(context Context) (Database, error)
	Context() context.Context
	Databases() map[string]Database
	Global(projectID string) (Database, error)
	SmartOps() smartOps.Service
}
