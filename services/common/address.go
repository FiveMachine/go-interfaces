package common

import (
	"fmt"

	libp2p "github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func convertToAddrInfo(peers []string) ([]libp2p.AddrInfo, error) {
	addr := make([]libp2p.AddrInfo, 0)
	for _, _addr := range peers {
		addrInfo, err := convertToMultiAddr(_addr)
		if err != nil {
			return nil, fmt.Errorf("converting `%s` to multi addr failed with: %s", _addr, err)
		}

		addr = append(addr, *addrInfo)
	}

	return addr, nil
}

func convertToMultiAddr(addr string) (*libp2p.AddrInfo, error) {
	_multiaddr, err := multiaddr.NewMultiaddr(addr)
	if err != nil {
		return nil, fmt.Errorf("converting `%s` to a multi address failed with: %s", addr, err)
	}

	addrInfo, err := libp2p.AddrInfoFromP2pAddr(_multiaddr)
	if err != nil {
		return nil, fmt.Errorf("getting addr from p2p addr failed with: %s", err)
	}

	return addrInfo, nil

}
