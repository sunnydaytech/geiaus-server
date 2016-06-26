package main

import (
	"flag"
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"github.com/sunnydaytech/geiaus-server/server"
	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	projectId := flag.String("gcloud_project_id", "", "Google Cloud Storage project id")
	flag.Parse()
	var client *datastore.Client
	var err error
	if *projectId != "" {
		client, err = datastore.NewClient(context.Background(), *projectId)
		if err != nil {
			panic(err.Error())
		}
	}

	Start(":50051", server.NewUserServer(client), server.NewInMemSessionServer())
}

func Start(port string, userManagerServer pb.UserManageServer, sessionServer *server.SessionServer) {
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
