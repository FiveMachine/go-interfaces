package ipfs

import (
	"context"
	"io"

	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"

	// "github.com/ipfs/go-ipld-format"
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
	// Remove removes a node from this DAG.
	// Remove returns no error if the requested node is not present in this DAG.
	Remove(context.Context, cid.Cid) error
	// RemoveMany removes many nodes from this DAG.
	//
	// It returns success even if the nodes were not present in the DAG.
	RemoveMany(context.Context, []cid.Cid) error

	// Add adds a node to this DAG.
	Add(context.Context, ipld.Node) error
	// AddMany adds many nodes to this DAG.
	//
	// Consider using the Batch NodeAdder (`NewBatch`) if you make
	// extensive use of this function.
	AddMany(context.Context, []ipld.Node) error

	// Get retrieves nodes by CID. Depending on the NodeGetter
	// implementation, this may involve fetching the Node from a remote
	// machine; consider setting a deadline in the context.
	Get(context.Context, cid.Cid) (ipld.Node, error)

	// GetMany returns a channel of NodeOptions given a set of CIDs.
	GetMany(context.Context, []cid.Cid) <-chan *ipld.NodeOption
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
