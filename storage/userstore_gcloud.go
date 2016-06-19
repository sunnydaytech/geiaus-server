package storage

import (
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
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
		panic("Failed to create user.")
	}
	return user
}

func (s GCloudUserStore) DeleteUser(user *pb.User) *pb.User {
	ctx := context.Background()
	key := datastore.NewKey(ctx, KIND_USER, "", user.UserId, nil)
	err := s.client.Delete(ctx, key)
	if err != nil {
		panic("Failed to delete user.")
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
	users := []*pb.User{}
	s.client.GetAll(ctx, query, users)
	if len(users) == 1 {
		return users[0]
	}
	return nil
}

func (s GCloudUserStore) SetPassword(userId int64, hash []byte, salt string) *pb.User {
	user := s.LookupUserById(userId)
	if user == nil {
		panic("User not found: " + strconv.FormatInt(userId, 10))
	}
	user.AuthMethod = append(user.AuthMethod, &pb.AuthMethod{
		Value: &pb.AuthMethod_Password{
			Password: &pb.Password{
				Hash: hash,
				Salt: salt}}})
	ctx := context.Background()
	key := datastore.NewKey(ctx, KIND_USER, "", userId, nil)
	s.client.Put(ctx, key, user)
	return user
}
