package server

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/sunnydaytech/geiaus/service/proto"
	"github.com/sunnydaytech/geiaus/service/server"
	"golang.org/x/net/context"
	"testing"
)

func TestCreateUser(t *testing.T) {
	userManagerServer := server.NewInMemUserServer()

	context := context.Background()
	userId := int64(12)
	user := &pb.User{
		UserId: userId,
	}
	createUserResp, err := userManagerServer.CreateUser(context, &pb.CreateUserRequest{
		UserToCreate: user})
	if err != nil {
		t.Fatalf("CreateUser returns error!")
	}
	if !proto.Equal(user, createUserResp.CreatedUser) {
		t.Errorf("Created user doesn't match.")
	}
	lookupUserResp, err := userManagerServer.LookupUser(context, &pb.LookupUserRequest{
		UserId: userId})
	if err != nil {
		t.Fatalf("Lookup user returns error!")
	}
	if !proto.Equal(user, lookupUserResp.User) {
		t.Errorf("Looked up user doesn't match.")
	}

}

func TestSetAndCheckPassword(t *testing.T) {
	userManagerServer := server.NewInMemUserServer()

	context := context.Background()
	userId := int64(12)
	user := &pb.User{
		UserId: userId,
	}
	_, err := userManagerServer.CreateUser(context, &pb.CreateUserRequest{
		UserToCreate: user})
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
