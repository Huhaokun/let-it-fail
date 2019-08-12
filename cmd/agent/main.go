package main

import (
	"fmt"
	"github.com/Huhaokun/let-it-fail/agent"
	"github.com/Huhaokun/let-it-fail/contract"
	. "github.com/Huhaokun/let-it-fail/log"
	"github.com/docker/docker/client"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var port = 7999

func main() {
	// start grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		Log.Errorf("fail to listen %v", err)
	}

	var opts [] grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	dockerCli, err := client.NewEnvClient()
	if err != nil {
		Log.Fatalf("fail to new docker client %v", err)
	}

	contract.RegisterEndpointRegistryServer(grpcServer, agent.NewEndpointRegistry(dockerCli))
	contract.RegisterStatusKillerServer(grpcServer, agent.NewStatusKiller(dockerCli))

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<- c
		os.Exit(1)
	}()

	err = grpcServer.Serve(lis)
	if err != nil {
		Log.Errorf("fail to serve %v", err)
	}

}
