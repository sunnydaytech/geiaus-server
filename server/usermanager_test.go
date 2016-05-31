package server

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"github.com/sunnydaytech/geiaus-server/server"
	"golang.org/x/net/context"
	"testing"
)

func TestCreateUser(t *testing.T) {
	userManagerServer := server.NewInMemUserServer()

	context := context.Background()

	createUserResp, err := userManagerServer.CreateUser(context, &pb.CreateUserRequest{
		UserName: "username"})
	if err != nil {
		t.Fatalf("CreateUser returns error!")
	}
	createdUserId := createUserResp.CreatedUser.UserId
	lookupUserResp, err := userManagerServer.LookupUser(context, &pb.LookupUserRequest{
		UserId: createdUserId})
	if err != nil {
		t.Fatalf("Lookup user returns error!")
	}
	expectedUser := &pb.User{
		UserId:   createdUserId,
		UserName: "username",
	}
	if !proto.Equal(expectedUser, lookupUserResp.User) {
		t.Errorf("Looked up user doesn't match.")
	}

}

func TestSetAndCheckPassword(t *testing.T) {
	userManagerServer := server.NewInMemUserServer()

	context := context.Background()
	_, err := userManagerServer.CreateUser(context, &pb.CreateUserRequest{
		UserName: "username"})
	if err != nil {
		t.Fatalf("CreateUser returns error!")
	}
	password := "myPassword"
	userManagerServer.SetPassword(context, &pb.SetPasswordRequest{
		UserId:   12,
		Password: password})
	checkPasswordResp, _ := userManagerServer.CheckPassword(context, &pb.CheckPasswordRequest{
		UserId:   12,
		Password: password})
	if !checkPasswordResp.Match {
		t.Fatalf("Check Password fails!")
	}
	checkPasswordResp, _ = userManagerServer.CheckPassword(context, &pb.CheckPasswordRequest{
		UserId:   12,
		Password: password + "wrongPassword"})
	if checkPasswordResp.Match {
		t.Fatalf("Check Password should fail!")
	}
}
