package main

import (
	"context"
	"log"
	"net"

	"github.com/AllanM007/grpc-services/greeter"
	"google.golang.org/grpc"
)

type GreeterServer struct {
	greeter.UnimplementedGreeterServer
}

func (s GreeterServer) SayHello(ctx context.Context, in *greeter.HelloRequest) (*greeter.HelloReply, error) {
	return &greeter.HelloReply{
		Name: string("Hello ") + in.Name,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %v", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &GreeterServer{}

	greeter.RegisterGreeterServer(serverRegistrar, service)

	err = serverRegistrar.Serve(listener)
	if err != nil {
		log.Fatalf("not able to serve: %v", err)
	}
}
