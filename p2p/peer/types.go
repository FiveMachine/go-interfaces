package peer

import (
	"context"
	"io"
	"time"

	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/config"
	"github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/taubyte/utils/fs/dir"
)

type ReadSeekCloser interface {
	io.ReadSeekCloser
	io.WriterTo
}

type Node interface {
	AddFile(r io.Reader) (string, error)
	AddFileForCid(r io.Reader) (cid.Cid, error)
	BootstrapPeers() []peer.AddrInfo
	Close()
	Context() context.Context
	DAG() *ipfslite.Peer
	DeleteFile(id string) error
	Discovery() discovery.Discovery
	Done() <-chan struct{}
	GetFile(ctx context.Context, id string) (ReadSeekCloser, error)
	GetFileFromCid(ctx context.Context, cid cid.Cid) (ReadSeekCloser, error)
	ID() peer.ID
	Messaging() *pubsub.PubSub
	NewChildContextWithCancel() (context.Context, context.CancelFunc)
	NewFolder(name string) (dir.Directory, error)
	NewPubSubKeepAlive(ctx context.Context, cancel context.CancelFunc, name string) error
	Peer() host.Host
	Peering() PeeringService
	Ping(pid string, count int) (int, time.Duration, error)
	PubSubPublish(ctx context.Context, name string, data []byte) error
	PubSubSubscribe(name string, handler PubSubConsumerHandler, err_handler PubSubConsumerErrorHandler) error
	PubSubSubscribeContext(ctx context.Context, name string, handler PubSubConsumerHandler, err_handler PubSubConsumerErrorHandler) error
	PubSubSubscribeToTopic(topic *pubsub.Topic, handler PubSubConsumerHandler, err_handler PubSubConsumerErrorHandler) error
	SimpleAddrsFactory(announce []string, override bool) config.Option
	Store() datastore.Batching
	WaitForSwarm(timeout time.Duration) error
}

type PubSubConsumerHandler func(msg *pubsub.Message)
type PubSubConsumerErrorHandler func(err error)

type PeeringService interface {
	Start() error
	Stop() error
	AddPeer(peer.AddrInfo)
	RemovePeer(peer.ID)
}
