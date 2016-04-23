package server

import (
	pb "geia.us/service/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserManagerServer struct {
}

func (s *UserManagerServer) CreateUser(context context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{request.UserToCreate}, nil
}

func (s *UserManagerServer) DeleteUser(context context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}

func Start(port string, userManagerServer *UserManagerServer) {
	log.Printf("Starting server on port %s\n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManageServer(s, userManagerServer)
	s.Serve(lis)
}
