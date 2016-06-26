package storage

import (
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"strconv"
)

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
	user, ok := s.userIdMap[userId]
	if !ok {
		return nil
	}
	return &user
}

func (s InMemUserStore) LookupUserByUserName(userName string) *pb.User {
	user, ok := s.usernameMap[userName]
	if !ok {
		return nil
	}
	return &user
}

func (s InMemUserStore) SetPassword(userId int64, hash []byte, salt string) *pb.User {
	user := s.LookupUserById(userId)
	if user == nil {
		panic("User not found: " + strconv.FormatInt(userId, 10))
	}
	user.PasswordHash = hash
	user.PasswordSalt = salt
	s.userIdMap[userId] = *user
	return user
}
