package mocks

import (
	"bytes"
	"context"
	"fmt"
	"io"

	ipfslite "github.com/hsanjuan/ipfs-lite"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	libp2pPerr "github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/routing"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
	"github.com/taubyte/go-interfaces/p2p/peer"
)

func (m *mockNode) add(r io.Reader) (_cid cid.Cid, err error) {
	var data []byte

	if data, err = io.ReadAll(r); err != nil {
		return
	}

	prefix := cid.Prefix{
		Version:  1,
		Codec:    uint64(mc.Raw),
		MhType:   mh.SHA2_256,
		MhLength: -1,
	}

	if _cid, err = prefix.Sum(data); err != nil {
		return
	}

	m.lock.Lock()
	m.mapDef[_cid.String()] = data
	m.lock.Unlock()

	return
}

func (m *mockNode) get(_cid string) (peer.ReadSeekCloser, error) {
	m.lock.RLock()
	data, exists := m.mapDef[_cid]
	m.lock.RUnlock()
	if !exists {
		return nil, fmt.Errorf("file cid `%s` does not exist", _cid)
	}

	return &mockReadSeekCloser{
		Buffer: bytes.NewBuffer(data),
	}, nil
}

func (m *mockNode) AddFile(r io.Reader) (string, error) {
	_cid, err := m.add(r)
	if err != nil {
		return "", err
	}

	return _cid.String(), nil
}

func (m *mockNode) AddFileForCid(r io.Reader) (cid.Cid, error) {
	return m.add(r)
}

func (m *mockNode) GetFile(_ctx context.Context, id string) (peer.ReadSeekCloser, error) {
	return m.get(id)
}

func (m *mockNode) GetFileFromCid(_ctx context.Context, cid cid.Cid) (peer.ReadSeekCloser, error) {
	return m.get(cid.String())
}

func (m *mockNode) DeleteFile(id string) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, exists := m.mapDef[id]; !exists {
		return fmt.Errorf("file cid `%s` does not exist", id)
	}

	delete(m.mapDef, id)

	return nil
}

func (m *mockNode) DAG() *ipfslite.Peer {
	mockedDag, err := ipfslite.New(context.TODO(), &mockedDataStore{}, &mockedBlockStore{}, &mockedHost{}, &mockedDHT{}, nil)
	if err != nil {
		panic(err)
	}

	return mockedDag
}

type mockedDataStore struct {
	datastore.Batching
}

func (m *mockedDataStore) Query(ctx context.Context, q query.Query) (query.Results, error) {
	return &mockedQueryResults{}, nil
}

type mockedQueryResults struct {
	query.Results
}

func (m *mockedQueryResults) Close() error { return nil }
func (m *mockedQueryResults) NextSync() (query.Result, bool) {
	return query.Result{}, false
}

type mockedHost struct {
	host.Host
}

func (m *mockedHost) ConnManager() connmgr.ConnManager {
	return connmgr.NullConnMgr{}
}

func (m *mockedHost) ID() libp2pPerr.ID {
	return ""
}

func (m *mockedHost) SetStreamHandler(pid protocol.ID, handler network.StreamHandler) {}

func (m *mockedHost) Network() network.Network {
	return &mockedNetwork{}
}

type mockedNetwork struct {
	network.Network
}

func (m *mockedNetwork) Notify(network.Notifiee) {}

type mockedDHT struct {
	routing.Routing
}

func (m *mockNode) Close() {
	for k := range m.mapDef {
		delete(m.mapDef, k)
	}

}
func (m *mockNode) Context() context.Context {
	return m.context
}

func (m *mockReadSeekCloser) Close() error {
	return nil
}

func (m *mockReadSeekCloser) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func (m *mockedBlockStore) Has(context.Context, cid.Cid) (bool, error) {
	return false, nil
}

func (m *mockedBlockStore) DeleteBlock(context.Context, cid.Cid) error              { return nil }
func (m *mockedBlockStore) Get(context.Context, cid.Cid) (blocks.Block, error)      { return nil, nil }
func (m *mockedBlockStore) GetSize(context.Context, cid.Cid) (int, error)           { return 0, nil }
func (m *mockedBlockStore) Put(context.Context, blocks.Block) error                 { return nil }
func (m *mockedBlockStore) PutMany(context.Context, []blocks.Block) error           { return nil }
func (m *mockedBlockStore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) { return nil, nil }
func (m *mockedBlockStore) HashOnRead(enabled bool)                                 {}
