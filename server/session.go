package server

import (
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"github.com/sunnydaytech/geiaus-server/storage"
	"golang.org/x/net/context"
)

type SessionServer struct {
	sessionStore storage.SessionStore
}

func (s *SessionServer) CreateISession(context context.Context, request *pb.CreateISessionRequest) (*pb.CreateISessionResponse, error) {
	iSession := s.sessionStore.CreateISession(request.UserId)
	return &pb.CreateISessionResponse{
		ISession: iSession,
	}, nil
}

func (s *SessionServer) CreateBSession(context context.Context, request *pb.CreateBSessionRequest) (*pb.CreateBSessionResponse, error) {
	return &pb.CreateBSessionResponse{}, nil
}

func (s *SessionServer) CreateUSession(context context.Context, request *pb.CreateUSessionRequest) (*pb.CreateUSessionResponse, error) {
	return &pb.CreateUSessionResponse{}, nil
}

func NewInMemSessionServer() *SessionServer {
	return &SessionServer{
		sessionStore: storage.NewInMemSessionStore(),
	}
}
