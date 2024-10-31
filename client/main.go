package main

import (
	"context"
	"flag"
	"log"

	"github.com/AllanM007/grpc-services/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var defaultName string = "Tata"

var (
	name = flag.String("name", defaultName, "Name to greet")
)

type GreeterClient struct {
	greeter.UnimplementedGreeterServer
}

func main() {

	flag.Parse()

	clientConn, err := grpc.NewClient(":8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer clientConn.Close()

	client := greeter.NewGreeterClient(clientConn)

	response, err := client.SayHello(context.Background(), &greeter.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("couln't send greeting: %v", err)
	}

	log.Printf("Got message: %s", response.GetName())
}
