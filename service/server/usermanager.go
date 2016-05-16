package server

import (
	pb "github.com/sunnydaytech/geiaus/service/proto"
	"github.com/sunnydaytech/geiaus/service/storage"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserManagerServer struct {
	userStore storage.UserStore
}

func (s *UserManagerServer) CreateUser(context context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	createdUser := s.userStore.CreateUser(request.UserToCreate)
	return &pb.CreateUserResponse{
		CreatedUser: &createdUser}, nil
}

func (s *UserManagerServer) DeleteUser(context context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	deletedUser := s.userStore.DeleteUser(request.UserId)
	return &pb.DeleteUserResponse{
		DeletedUser: &deletedUser}, nil
}

func NewInMemUserServer() *UserManagerServer {
	return &UserManagerServer{
		userStore: storage.NewInMemUserStore()}
}

func (s *UserManagerServer) LookupUser(context context.Context, request *pb.LookupUserRequest) (*pb.LookupUserResponse, error) {
	user := s.userStore.LookupUser(request.UserId)
	return &pb.LookupUserResponse{
		User: &user}, nil
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
