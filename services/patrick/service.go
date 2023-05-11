package patrick

import (
	"context"
	"io"

	kv "bitbucket.org/taubyte/kvdb/database"
	"github.com/taubyte/go-interfaces/services"
)

type Service interface {
	services.Service
	KV() *kv.KVDatabase
	NewGitRepository(provider string, repositoryId string, output io.Writer) (GitRepository, error)
}

type GitRepository interface {
	Url() *string
	Clone(ctx context.Context, path string, ref string) error
	Path() string
}
