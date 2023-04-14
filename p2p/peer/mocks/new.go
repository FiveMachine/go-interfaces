package mocks

import "context"

func New(ctx context.Context) MockedNode {
	ctx, ctxC := context.WithCancel(ctx)
	return &mockNode{
		mapDef:   make(map[string][]byte, 0),
		context:  ctx,
		contextC: ctxC,
	}
}
