package database

import (
	"github.com/taubyte/go-interfaces/services/substrate/components"
)

type Service interface {
	components.ServiceComponent
	Database(context Context) (Database, error)
	Databases() map[string]Database
	Global(projectID string) (Database, error)
}
