package common

import peer "github.com/taubyte/go-interfaces/p2p/peer"

type Service interface {
	Node() peer.Node
	Close()
}
