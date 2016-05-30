package main

import (
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"github.com/sunnydaytech/geiaus-server/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	Start(":50051", server.NewInMemUserServer())
}

func Start(port string, userManagerServer *server.UserManagerServer) {
	log.Printf("Starting server on port %s\n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManageServer(s, userManagerServer)
	s.Serve(lis)
}
