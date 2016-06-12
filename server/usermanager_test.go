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
	username := "username"

	createUserResp, err := userManagerServer.CreateUser(context, &pb.CreateUserRequest{
		UserName: username})
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
	lookupUserResp, _ = userManagerServer.LookupUser(context, &pb.LookupUserRequest{
		UserName: username})
	if lookupUserResp.User.UserId != createdUserId {
		t.Error("Lookup user by username failed.")
	}

	nonExistUsername := "doens't exist"
	lookupUserResp, _ = userManagerServer.LookupUser(context, &pb.LookupUserRequest{
		UserName: nonExistUsername,
	})
	if lookupUserResp.User != nil {
		t.Errorf("lookup user resp should return nil User")
	}
}

func TestSetAndCheckPassword(t *testing.T) {
	userManagerServer := server.NewInMemUserServer()

	context := context.Background()
	createUserResp, err := userManagerServer.CreateUser(context, &pb.CreateUserRequest{
		UserName: "username"})
	if err != nil {
		t.Fatalf("CreateUser returns error!")
	}
	userId := createUserResp.CreatedUser.UserId
	password := "myPassword"
	userManagerServer.SetPassword(context, &pb.SetPasswordRequest{
		UserId:   userId,
		Password: password})
	checkPasswordResp, _ := userManagerServer.CheckPassword(context, &pb.CheckPasswordRequest{
		UserId:   userId,
		Password: password})
	if !checkPasswordResp.Match {
		t.Fatalf("Check Password fails!")
	}
	checkPasswordResp, _ = userManagerServer.CheckPassword(context, &pb.CheckPasswordRequest{
		UserId:   userId,
		Password: password + "wrongPassword"})
	if checkPasswordResp.Match {
		t.Fatalf("Check Password should fail!")
	}
}
