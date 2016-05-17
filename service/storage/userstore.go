package storage

import (
	pb "github.com/sunnydaytech/geiaus/service/proto"
)

type UserStore interface {
	CreateUser(user *pb.User) *pb.User
	DeleteUser(userId int64) *pb.User
	LookupUser(userId int64) *pb.User
	SetPassword(userId int64, hash string, salt string) *pb.User
}

// InMemUserStore in menmory impl of interface UserStore
type InMemUserStore struct {
	userMap map[int64]pb.User
}

func (s InMemUserStore) CreateUser(user *pb.User) *pb.User {
	s.userMap[user.UserId] = *user
	return user
}

func (s InMemUserStore) DeleteUser(userId int64) *pb.User {
	userToBeDeleted := s.userMap[userId]
	delete(s.userMap, userId)
	return &userToBeDeleted
}

func (s InMemUserStore) LookupUser(userId int64) *pb.User {
	user := s.userMap[userId]
	return &user
}

func (s InMemUserStore) SetPassword(userId int64, hash string, salt string) *pb.User {
	user := s.LookupUser(userId)
	user.AuthMethod = append(user.AuthMethod, &pb.AuthMethod{
		Value: &pb.AuthMethod_Password{
			Password: &pb.Password{
				Hash: hash,
				Salt: salt}}})
	return user
}

func NewInMemUserStore() UserStore {
	userMap := map[int64]pb.User{}
	return InMemUserStore{
		userMap: userMap,
	}
}
