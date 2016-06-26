package storage

import (
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
	"log"
	"strconv"
)

const (
	KIND_USER = "Geiaus_User"
)

type GCloudUserStore struct {
	client *datastore.Client
}

func (s GCloudUserStore) CreateUser(user *pb.User) *pb.User {
	ctx := context.Background()
	key := datastore.NewKey(ctx, KIND_USER, "", user.UserId, nil)
	_, err := s.client.Put(ctx, key, user)
	if err != nil {
		panic("Failed to create user." + err.Error())
	}
	return user
}

func (s GCloudUserStore) DeleteUser(userId int64) *pb.User {
	ctx := context.Background()
	key := datastore.NewKey(ctx, KIND_USER, "", userId, nil)
	user := s.LookupUserById(userId)
	err := s.client.Delete(ctx, key)
	if err != nil {
		panic("Failed to delete user." + err.Error())
	}
	return user
}

func (s GCloudUserStore) LookupUserById(userId int64) *pb.User {
	ctx := context.Background()
	key := datastore.NewKey(ctx, KIND_USER, "", userId, nil)
	user := &pb.User{}
	err := s.client.Get(ctx, key, user)
	if err != nil {
		return nil
	}
	return user
}

func (s GCloudUserStore) LookupUserByUserName(userName string) *pb.User {
	ctx := context.Background()
	query := datastore.NewQuery(KIND_USER).Filter("UserName =", userName)
	users := &[]*pb.User{}
	_, err := s.client.GetAll(ctx, query, users)
	if err != nil {
		log.Printf("Failed to lookup: " + err.Error())
	}
	if len(*users) == 1 {
		return (*users)[0]
	}
	return nil
}

func (s GCloudUserStore) SetPassword(userId int64, hash []byte, salt string) *pb.User {
	user := s.LookupUserById(userId)
	if user == nil {
		panic("User not found: " + strconv.FormatInt(userId, 10))
	}
	user.PasswordHash = hash
	user.PasswordSalt = salt
	ctx := context.Background()
	key := datastore.NewKey(ctx, KIND_USER, "", userId, nil)
	s.client.Put(ctx, key, user)
	return user
}
