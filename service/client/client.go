package main

import (
	pb "geia.us/service/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloWorldClient(conn)

	r, er := c.Print(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not print: %v", err)
	}
	log.Printf(r.Context)
}
