package mocks

import (
	"bytes"
	"context"
	"io"
	"sync"

	"github.com/taubyte/go-interfaces/p2p/peer"
)

type MockedNode interface {
	peer.Node
}

type mockNode struct {
	mapDef map[string][]byte
	lock   sync.RWMutex
	peer.Node
	context  context.Context
	contextC context.CancelFunc
}

type MockedReadSeekCloser interface {
	peer.ReadSeekCloser
}

type mockReadSeekCloser struct {
	*bytes.Buffer
	io.Seeker
	io.Writer
}
