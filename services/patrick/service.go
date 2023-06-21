package patrick

import (
	"context"
	"io"

	"github.com/taubyte/go-interfaces/kvdb"
	"github.com/taubyte/go-interfaces/services"
)

type Service interface {
	services.Service
	KV() kvdb.KVDB
	NewGitRepository(provider string, repositoryId string, output io.Writer) (GitRepository, error)
}

type GitRepository interface {
	Url() *string
	Clone(ctx context.Context, path string, ref string) error
	Path() string
}
