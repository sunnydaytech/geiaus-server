package server

import (
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"github.com/sunnydaytech/geiaus-server/storage"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"math/rand"
	"time"
)

type UserManagerServer struct {
	userStore storage.UserStore
}

func (s *UserManagerServer) CreateUser(context context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	createdUser := s.userStore.CreateUser(request.UserToCreate)
	return &pb.CreateUserResponse{
		CreatedUser: createdUser}, nil
}

func (s *UserManagerServer) DeleteUser(context context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	deletedUser := s.userStore.DeleteUser(request.UserId)
	return &pb.DeleteUserResponse{
		DeletedUser: deletedUser}, nil
}

func (s *UserManagerServer) LookupUser(context context.Context, request *pb.LookupUserRequest) (*pb.LookupUserResponse, error) {
	user := s.userStore.LookupUser(request.UserId)
	return &pb.LookupUserResponse{
		User: user}, nil
}

func (s *UserManagerServer) SetPassword(context context.Context, request *pb.SetPasswordRequest) (*pb.SetPasswordResponse, error) {
	salt := newSalt()
	passwordBytes := []byte(request.Password + salt)
	hash, _ := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	user := s.userStore.SetPassword(request.UserId, hash, salt)
	return &pb.SetPasswordResponse{
		UpdatedUser: user}, nil
}

func (s *UserManagerServer) CheckPassword(context context.Context, request *pb.CheckPasswordRequest) (*pb.CheckPasswordResponse, error) {
	user := s.userStore.LookupUser(request.UserId)
	authMethod := user.AuthMethod[0]
	passwordData := authMethod.GetPassword()
	passwordBytes := []byte(request.Password + passwordData.Salt)
	err := bcrypt.CompareHashAndPassword(passwordData.Hash, passwordBytes)
	if err == nil {
		return &pb.CheckPasswordResponse{
			Match: true}, nil
	} else {
		return &pb.CheckPasswordResponse{
			Match: false}, nil
	}

}

var letterRunes = []rune("!@#$%^&*()_+~1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func newSalt() string {
	rand.Seed(time.Now().UnixNano())
	s := make([]rune, 10)
	for i := range s {
		s[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(s)
}

func NewInMemUserServer() *UserManagerServer {
	return &UserManagerServer{
		userStore: storage.NewInMemUserStore()}
}
