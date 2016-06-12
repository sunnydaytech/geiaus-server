package storage

import (
	pb "github.com/sunnydaytech/geiaus-server/proto"
)

type SessionStore interface {
	CreateISession(userId int64, iSessionId string) *pb.ISession
	LookupISessionById(iSessionId string) *pb.ISession
}

type InMemSessionStore struct {
	iSessionIdMap map[string]pb.ISession
}

func (s InMemSessionStore) CreateISession(userId int64, iSessionId string) *pb.ISession {
	iSession := &pb.ISession{
		Id:     iSessionId,
		UserId: userId,
	}
	s.iSessionIdMap[iSessionId] = *iSession
	return iSession
}

func (s InMemSessionStore) LookupISessionById(iSessionId string) *pb.ISession {
	iSession, ok := s.iSessionIdMap[iSessionId]
	if !ok {
		return nil
	} else {
		return &iSession
	}
}

func NewInMemSessionStore() SessionStore {
	iSessionIdMap := map[string]pb.ISession{}
	return InMemSessionStore{
		iSessionIdMap: iSessionIdMap,
	}
}
