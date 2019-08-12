package admin

import (
	"context"
	"fmt"
	"github.com/Huhaokun/let-it-fail/contract"
	. "github.com/Huhaokun/let-it-fail/log"
	"google.golang.org/grpc"
)

type Node interface {
	Id() string
	ListEndpoint(ctx context.Context, cmd *contract.ListCommand) (*contract.Endpoints, error)
	contract.StatusKillerClient
}

type node struct {
	e contract.EndpointRegistryClient
	contract.StatusKillerClient
	id string
}

func NewNode(ip string, port int) Node {
	target := fmt.Sprintf("%s:%d", ip, port)
	cc, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		Log.Errorf("fail to dial agent %s", target)
	}
	return &node{
		e:                  contract.NewEndpointRegistryClient(cc),
		StatusKillerClient: contract.NewStatusKillerClient(cc),
		id:                 ip,
	}
}

func (n *node) ListEndpoint(ctx context.Context, cmd *contract.ListCommand) (*contract.Endpoints, error) {
	endpoints, err := n.e.List(ctx, cmd)
	if err == nil {
		for _, endpoint := range endpoints.Endpoints {
			endpoint.Host = n.Id()
		}
	}

	return endpoints, err
}

//func (n *node) Stop(ctx context.Context, filter *contract.EndpointFilter) (*contract.OpResult, error) {
//	return n.k.Stop(ctx, filter)
//}
//
//func (n *node) Kill(ctx context.Context, filter *contract.EndpointFilter) (*contract.OpResult, error) {
//	return n.k.Kill(ctx, filter)
//}
//
//func (n *node) Pause(ctx context.Context, filter *contract.EndpointFilter) (*contract.OpResult, error) {
//	return n.k.Pause(ctx, filter)
//}

func (n *node) Id() string {
	return n.id
}
