package storage

import (
	"context"
	"io"

	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-interfaces/kvdb"
	"github.com/taubyte/go-interfaces/services/substrate/components"
	structureSpec "github.com/taubyte/go-specs/structure"
	peer "github.com/taubyte/p2p/peer"
)

// Context holds the configuration and context for storage operations.
type Context struct {
	context.Context
	ProjectId     string
	ApplicationId string
	Matcher       string
	Config        *structureSpec.Storage
}

// Meta provides metadata operations for storage items.
type Meta interface {
	Get() (io.ReadSeekCloser, error)
	Cid() cid.Cid
	Version() int
}

// Service defines the interface for storage services.
type Service interface {
	components.ServiceComponent
	Storages() map[string]Storage
	Get(context Context) (Storage, error)
	Storage(context Context) (Storage, error)
	Add(r io.Reader) (cid.Cid, error)
	GetFile(ctx context.Context, cid cid.Cid) (peer.ReadSeekCloser, error)
}

// Storage defines the interface for individual storage operations.
type Storage interface {
	AddFile(ctx context.Context, r io.ReadSeeker, name string, replace bool) (int, error)
	DeleteFile(ctx context.Context, name string, version int) error
	Meta(ctx context.Context, name string, version int) (Meta, error)
	ListVersions(ctx context.Context, name string) ([]string, error)
	GetLatestVersion(ctx context.Context, name string) (int, error)
	List(ctx context.Context, prefix string) ([]string, error)
	Close()
	Used(ctx context.Context) (int, error)
	Capacity() int
	Id() string
	Kvdb() kvdb.KVDB
	ContextConfig() Context
	UpdateCapacity(size uint64)
	Config() *structureSpec.Storage
}
