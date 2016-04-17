package main

import (
	"fmt"
	pb "geia.us/service/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type helloWorldServer struct {
}

func (s *helloWorldServer) Print(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{Context: "Hello world!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &helloWorldServer{})
	s.Serve(lis)
	fmt.Printf("Hello,world.\n")
}
