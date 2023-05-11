package common

import (
	"fmt"

	"bitbucket.org/taubyte/p2p/peer"
)

func convertBootstrap(peers []string, devMode bool) (peer.BootstrapParams, error) {
	// To Bypass having no peers if launching not from odo
	if Deployment == Taubyte {
		if devMode && len(peers) < 1 {
			return peer.StandAlone(), nil
		}
		return peer.Bootstrap(), nil
	}

	if len(peers) > 0 {
		peers, err := convertToAddrInfo(peers)
		if err != nil {
			return peer.BootstrapParams{}, fmt.Errorf("converting peers to libp2p addr info failed with: %s", err)
		}

		return peer.Bootstrap(peers...), nil
	}

	return peer.StandAlone(), nil
}
