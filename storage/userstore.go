package storage

import (
	pb "github.com/sunnydaytech/geiaus-server/proto"
)

type UserStore interface {
	CreateUser(user *pb.User) *pb.User
	DeleteUser(userId int64) *pb.User
	LookupUserById(userId int64) *pb.User
	LookupUserByUserName(userName string) *pb.User
	SetPassword(userId int64, hash []byte, salt string) *pb.User
}

// InMemUserStore in menmory impl of interface UserStore
type InMemUserStore struct {
	userIdMap   map[int64]pb.User
	usernameMap map[string]pb.User
}

func (s InMemUserStore) CreateUser(user *pb.User) *pb.User {
	s.userIdMap[user.UserId] = *user
	if user.UserName != "" {
		s.usernameMap[user.UserName] = *user
	}
	return user
}

func (s InMemUserStore) DeleteUser(userId int64) *pb.User {
	userToBeDeleted := s.userIdMap[userId]
	delete(s.userIdMap, userId)
	return &userToBeDeleted
}

func (s InMemUserStore) LookupUserById(userId int64) *pb.User {
	user := s.userIdMap[userId]
	return &user
}

func (s InMemUserStore) LookupUserByUserName(userName string) *pb.User {
	user := s.usernameMap[userName]
	return &user
}

func (s InMemUserStore) SetPassword(userId int64, hash []byte, salt string) *pb.User {
	user := s.LookupUserById(userId)
	user.AuthMethod = append(user.AuthMethod, &pb.AuthMethod{
		Value: &pb.AuthMethod_Password{
			Password: &pb.Password{
				Hash: hash,
				Salt: salt}}})
	s.userIdMap[userId] = *user
	return user
}

func NewInMemUserStore() UserStore {
	userIdMap := map[int64]pb.User{}
	usernameMap := map[string]pb.User{}
	return InMemUserStore{
		userIdMap:   userIdMap,
		usernameMap: usernameMap,
	}
}
