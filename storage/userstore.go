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

// Returns in-mem impl of UserStore.
func NewInMemUserStore() UserStore {
	userIdMap := map[int64]pb.User{}
	usernameMap := map[string]pb.User{}
	return InMemUserStore{
		userIdMap:   userIdMap,
		usernameMap: usernameMap,
	}
}
