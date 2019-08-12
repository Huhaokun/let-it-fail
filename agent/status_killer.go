package agent

import (
	"context"
	"github.com/Huhaokun/let-it-fail/contract"
	"github.com/docker/docker/client"
	"time"
)

type StatusKiller struct {
	docker *client.Client
}

func NewStatusKiller(dockerCli *client.Client) *StatusKiller {
	return &StatusKiller{
		docker: dockerCli,
	}
}

func (k *StatusKiller) Stop(ctx context.Context, filter *contract.EndpointFilter) (*contract.OpResult, error) {
	timeout := 30 * time.Second
	return &contract.OpResult{}, k.docker.ContainerStop(ctx, filter.Id, &timeout)
}
func (k *StatusKiller) Kill(ctx context.Context, filter *contract.EndpointFilter) (*contract.OpResult, error) {
	return &contract.OpResult{}, k.docker.ContainerKill(ctx, filter.Id, "SIGKILL")
}
func (k *StatusKiller) Pause(ctx context.Context, filter *contract.EndpointFilter) (*contract.OpResult, error) {
	return &contract.OpResult{}, k.docker.ContainerPause(ctx, filter.Id)
}
