package client

import (
	pb "github.com/sunnydaytech/geiaus/service/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

type UserManageClient struct {
	userManageClient pb.UserManageClient
}

func (client *UserManageClient) CreateUser(request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	resp, _ := client.userManageClient.CreateUser(context.Background(), request)
	return resp, nil
}

func NewUserManageClient() *UserManageClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManageClient(conn)
	return &UserManageClient{c}
}
