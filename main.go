package main

import (
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"github.com/sunnydaytech/geiaus-server/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	Start(":50051", server.NewInMemUserServer(), server.NewInMemSessionServer())
}

func Start(port string, userManagerServer *server.UserManagerServer, sessionServer *server.SessionServer) {
	log.Printf("Starting server on port %s\n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManageServer(s, userManagerServer)
	pb.RegisterSessionServer(s, sessionServer)
	s.Serve(lis)
}
