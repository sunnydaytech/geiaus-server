package main

import (
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"github.com/sunnydaytech/geiaus-server/server"
	"github.com/sunnydaytech/geiaus-server/storage"
	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Post struct {
	Title       string
	Body        string `datastore:",noindex"`
	PublishedAt time.Time
}

func main() {
	ctx := context.Background()
	client := storage.NewDataStoreClient()
	key := datastore.NewKey(ctx, "Post", "post1", 0, nil)
	post := &Post{
		Title:       "title",
		Body:        "body",
		PublishedAt: time.Now(),
	}
	_, err := client.Put(ctx, key, post)
	if err != nil {
		log.Fatalln("Failed to put data")
	}

	Start(":50051", server.NewInMemUserServer(), server.NewInMemSessionServer())
}

func Start(port string, userManagerServer *server.UserManagerServer, sessionServer *server.SessionServer) {
	log.Printf("Starting server on port %s\n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManageServer(s, userManagerServer)
	pb.RegisterSessionServer(s, sessionServer)
	s.Serve(lis)
}
