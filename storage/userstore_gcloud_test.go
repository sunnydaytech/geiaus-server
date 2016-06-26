package storage

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"github.com/sunnydaytech/geiaus-server/storage"
	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
	"reflect"
	"testing"
)

const (
	PROJECT_ID = "booming-octane-752" //name=shuwen
)

func TestCreateUser(t *testing.T) {
	client, err := datastore.NewClient(context.Background(), PROJECT_ID)
	if err != nil {
		t.Errorf(err.Error())
	}
	userStore := storage.NewGCloudUserStore(client)
	userId := int64(1234)
	userStore.DeleteUser(userId)
	user := userStore.LookupUserById(userId)
	if user != nil {
		t.Errorf("Deleting user failed")
	}
	user = &pb.User{
		UserName: "username",
		UserId:   userId,
	}
	userStore.CreateUser(user)
	user2 := userStore.LookupUserById(userId)
	fmt.Printf("user %v\nuser2 %v\n", user, user2)
	if !proto.Equal(user, user2) {
		t.Errorf("create lookup user failed.")
	}

	hash := []byte("hash")
	userStore.SetPassword(userId, hash, "salt")
	user = userStore.LookupUserById(userId)
	if !reflect.DeepEqual(user.PasswordHash, hash) {
		t.Errorf("storing password failed.")
	}
}
