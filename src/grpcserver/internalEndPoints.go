package grpcserver

import (
	"context"
	protocom "github.com/hisitra/hedron/src/protos"
)

func (s *server) InternalGet(ctx context.Context, request *protocom.InternalGetRequest) (*protocom.Response, error) {
	panic("implement me")
}

func (s *server) InternalSet(ctx context.Context, request *protocom.InternalGetRequest) (*protocom.Response, error) {
	panic("implement me")
}

func (s *server) SendLeaderHeartbeat(ctx context.Context, request *protocom.LeaderHeartbeatRequest) (*protocom.Response, error) {
	panic("implement me")
}

func (s *server) AskVote(ctx context.Context, request *protocom.AskVoteRequest) (*protocom.Response, error) {
	panic("implement me")
}