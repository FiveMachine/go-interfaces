package database

import (
	"context"

	"github.com/taubyte/go-interfaces/moody"
	"github.com/taubyte/go-interfaces/services/substrate/components"
)

type Service interface {
	components.ServiceComponent
	Database(context Context) (Database, error)
	Context() context.Context
	Logger() moody.Logger
	Databases() map[string]Database
	Global(projectID string) (Database, error)
}
