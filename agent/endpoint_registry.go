package agent

import (
	"context"
	"github.com/Huhaokun/let-it-fail/contract"
	. "github.com/Huhaokun/let-it-fail/log"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type EndpointRegistry struct {
	Docker *client.Client
}

func NewEndpointRegistry() *EndpointRegistry {
	dockerCli, err := client.NewEnvClient()
	if err != nil {
		Log.Fatalf("fail to new docker client %v", err)
	}
	return &EndpointRegistry{
		Docker: dockerCli,
	}
}

func (e *EndpointRegistry) List(ctx context.Context, cmd *contract.ListCommand) (*contract.Endpoints, error) {
	containers, err := e.Docker.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	var endpoints [] *contract.Endpoint
	for _, container := range containers {
		endpoints = append(endpoints, containerToEndpoint(&container))
	}

	return &contract.Endpoints{
		Endpoints: endpoints,
	}, nil
}

func containerToEndpoint(container *types.Container) *contract.Endpoint {
	return &contract.Endpoint{
		Id: container.ID,
	}
}
