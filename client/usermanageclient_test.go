package client

import (
	"geia.us/service/client"
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	userManageClient := client.NewUserManageClient()
	resp, _ := userManageClient.CreateUser(&pb.CreateUserRequest{})
	log.Printf("%v", resp)
}
