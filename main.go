package main

import (
	"context"
	"log"
	"net"

	"github.com/maorSeri/hellogrpc/messages"

	"google.golang.org/grpc"
)

const port = ":50000"

type server struct{}

//mandetory to call method sayHello
func (s *server) SayHello(ctx context.Context, req *messages.HelloRequest) (*messages.HelloResponse, error) {
	return &messages.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	// open a listener on tcp port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a new server using the grpc package
	s := grpc.NewServer()

	//regester hello worled server by passing a reference
	messages.RegisterHelloServiceServer(s, &server{})

	s.Serve(lis)
}
