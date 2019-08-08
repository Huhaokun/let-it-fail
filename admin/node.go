package admin

import (
	"context"
	"github.com/Huhaokun/let-it-fail/contract"
	. "github.com/Huhaokun/let-it-fail/log"
	"google.golang.org/grpc"
)

type Node interface {
	ListEndpoint(ctx context.Context, cmd *contract.ListCommand) (*contract.Endpoints, error)
}

type node struct {
	e contract.EndpointRegistryClient
}

func NewNode(target string) Node {
	cc, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		Log.Errorf("fail to dial agent %s", target)
	}
	return &node{
		e: contract.NewEndpointRegistryClient(cc),
	}
}

func (n *node) ListEndpoint(ctx context.Context, cmd *contract.ListCommand) (*contract.Endpoints, error) {
	return n.e.List(ctx, cmd)
}
