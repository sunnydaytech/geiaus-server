package server

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"github.com/sunnydaytech/geiaus-server/server"
	"golang.org/x/net/context"
	"testing"
)

func TestCreateISession(t *testing.T) {
	sessionServer := server.NewInMemSessionServer()
	context := context.Background()
	createISessionRequest := &pb.CreateISessionRequest{
		UserId: 1,
	}
	createISessionResp, err := sessionServer.CreateISession(context, createISessionRequest)
	if err != nil {
		t.Fatalf("Failed to create isession")
	}
	iSession := createISessionResp.ISession
	if iSession.UserId != 1 {
		t.Error("userId doesn't match")
	}
	if iSession.Id == "" {
		t.Error("sessionId missing")
	}

	lookupISessionRequest := &pb.LookupISessionRequest{
		Id: iSession.Id,
	}
	lookupISessionResp, err := sessionServer.LookupISession(context, lookupISessionRequest)
	if !proto.Equal(iSession, lookupISessionResp.ISession) {
		t.Errorf("Looked ISession failed.")
	}
}
