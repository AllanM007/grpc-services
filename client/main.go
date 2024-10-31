package main

import (
	"context"
	"log"
	"time"

	"github.com/AllanM007/grpc-services/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GreeterClient struct {
	greeter.UnimplementedGreeterServer
}

func main() {

	clientConn, err := grpc.NewClient(":8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer clientConn.Close()

	c := greeter.NewGreeterClient(clientConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := c.SayHello(ctx, &greeter.HelloRequest{Name: "Ola"})
	if err != nil {
		log.Fatalf("couln't send greeting: %v", err)
	}

	log.Printf("Got message: %s", req.GetName())
}
