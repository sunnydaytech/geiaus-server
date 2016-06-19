package storage

import (
	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
	"log"
)

const (
	PROJECT_ID = "nich01as-com"
)

func NewDataStoreClient() *datastore.Client {
	client, err := datastore.NewClient(context.Background(), PROJECT_ID)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
