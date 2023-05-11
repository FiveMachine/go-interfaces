package common

import (
	"context"
	"fmt"
	"time"

	p2pPeer "bitbucket.org/taubyte/p2p/peer"
	"github.com/taubyte/go-interfaces/p2p/peer"
)

var WaitForSwamDuration = 10 * time.Second
var Deployment = Taubyte

func (config *GenericConfig) NewNode(ctx context.Context, databaseName string) (node peer.Node, err error) {
	var p2pNode *p2pPeer.Node
	if config.DevMode {
		return config.NewLiteNode(ctx, databaseName)
	} else {
		boostrapParam, err := convertBootstrap(config.Peers, config.DevMode)
		if err != nil {
			return nil, fmt.Errorf("getting bootstrap perms in NewNode failed with: %s", err)
		}

		p2pNode, err = p2pPeer.NewPublic(
			ctx,
			config.Root+databaseName,
			config.PrivateKey,
			config.SwarmKey,
			config.P2PListen,
			config.P2PAnnounce,
			boostrapParam,
		)
		if err != nil {
			return nil, err
		}

	}

	err = p2pNode.WaitForSwarm(WaitForSwamDuration)
	if err != nil {
		return nil, err
	}

	return p2pNode, nil
}

func (config *GenericConfig) NewLiteNode(ctx context.Context, databaseName string) (peer.Node, error) {
	boostrapParam, err := convertBootstrap(config.Peers, config.DevMode)
	if err != nil {
		return nil, fmt.Errorf("getting bootstrap perms in NewLiteNode failed with: %s", err)
	}

	node, err := p2pPeer.NewLitePublic(
		ctx,
		config.Root+databaseName,
		config.PrivateKey,
		config.SwarmKey,
		config.P2PListen,
		config.P2PAnnounce,
		boostrapParam,
	)
	if err != nil {
		return nil, err
	}

	err = node.WaitForSwarm(WaitForSwamDuration)
	if err != nil {
		return nil, err
	}

	return node, nil
}
