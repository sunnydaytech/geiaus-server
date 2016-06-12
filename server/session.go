package server

import (
	pb "github.com/sunnydaytech/geiaus-server/proto"
	"github.com/sunnydaytech/geiaus-server/storage"
	"golang.org/x/net/context"
)

var (
	SESSION_RUNES = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

const (
	SESSION_LENGTH = 20
)

type SessionServer struct {
	sessionStore storage.SessionStore
}

func (s *SessionServer) CreateISession(context context.Context, request *pb.CreateISessionRequest) (*pb.CreateISessionResponse, error) {
	iSessionId := randStr(SESSION_RUNES, SESSION_LENGTH)
	iSession := s.sessionStore.CreateISession(request.UserId, iSessionId)
	return &pb.CreateISessionResponse{
		ISession: iSession,
	}, nil
}

func (s *SessionServer) LookupISession(context context.Context, request *pb.LookupISessionRequest) (*pb.LookupISessionResponse, error) {
	iSession := s.sessionStore.LookupISessionById(request.Id)
	return &pb.LookupISessionResponse{
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
