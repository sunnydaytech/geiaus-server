package storage

import (
	pb "github.com/sunnydaytech/geiaus/service/proto"
)

type UserStore interface {
	CreateUser(user *pb.User) pb.User
	DeleteUser(userId int64) pb.User
	LookupUser(userId int64) pb.User
}

// InMemUserStore in menmory impl of interface UserStore
type InMemUserStore struct {
	userMap map[int64]pb.User
}

func (s InMemUserStore) CreateUser(user *pb.User) pb.User {
	s.userMap[user.UserId] = *user
	return s.userMap[user.UserId]
}

func (s InMemUserStore) DeleteUser(userId int64) pb.User {
	userToBeDeleted := s.userMap[userId]
	delete(s.userMap, userId)
	return userToBeDeleted
}

func (s InMemUserStore) LookupUser(userId int64) pb.User {
	return s.userMap[userId]
}

func NewInMemUserStore() UserStore {
	userMap := map[int64]pb.User{}
	return InMemUserStore{
		userMap: userMap,
	}
}
