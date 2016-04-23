package client

import (
	"geia.us/service/client"
	pb "geia.us/service/proto"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	userManageClient := client.NewUserManageClient()
	resp, _ := userManageClient.CreateUser(&pb.CreateUserRequest{})
	log.Printf("%v", resp)
}
