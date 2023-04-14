package ipfs

import (
	"context"
	"io"

	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	ipld "github.com/ipfs/go-ipld-format"
	ufsio "github.com/ipfs/go-unixfs/io"
	"github.com/libp2p/go-libp2p-core/peer"
)

type Peer interface {
	// Bootstrap is an optional helper to connect to the given peers and bootstrap
	// the Peer DHT (and Bitswap). This is a best-effort function. Errors are only
	// logged and a warning is printed when less than half of the given peers
	// could be contacted. It is fine to pass a list where some peers will not be
	// reachable.
	Bootstrap(peers []peer.AddrInfo) ([]peer.AddrInfo, error)
	DHTBootstrap() error
	// Session returns a session-based NodeGetter.
	Session(ctx context.Context) ipld.NodeGetter
	// AddFile chunks and adds content to the DAGService from a reader. The content
	// is stored as a UnixFS DAG (default for IPFS). It returns the root
	// ipld.Node.
	AddFile(ctx context.Context, r io.Reader, params *AddParams) (ipld.Node, error)
	// GetFile returns a reader to a file as identified by its root CID. The file
	// must have been added as a UnixFS DAG (default for IPFS).
	GetFile(ctx context.Context, c cid.Cid) (ufsio.ReadSeekCloser, error)
	// BlockStore offers access to the blockstore underlying the Peer's DAGService.
	BlockStore() blockstore.Blockstore
	// HasBlock returns whether a given block is available locally. It is
	// a shorthand for .Blockstore().Has().
	HasBlock(c cid.Cid) (bool, error)
}

// AddParams contains all of the configurable parameters needed to specify the
// importing process of a file.
type AddParams struct {
	Layout    string
	Chunker   string
	RawLeaves bool
	Hidden    bool
	Shard     bool
	NoCopy    bool
	HashFun   string
}
