package storage

import (
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"google.golang.org/cloud/datastore"
)

type UserStore interface {
	CreateOrUpdateUser(user *pb.User) *pb.User
	DeleteUser(userId int64) *pb.User
	LookupUserById(userId int64) *pb.User
	LookupUserByUserName(userName string) *pb.User
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

func NewGCloudUserStore(client *datastore.Client) UserStore {
	return GCloudUserStore{
		client: client,
	}
}
