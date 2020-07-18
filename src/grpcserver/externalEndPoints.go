package grpcserver

import (
	"context"
	protocom "github.com/hisitra/hedron/src/protos"
)

func (s *server) Get(ctx context.Context, request *protocom.ExternalGetRequest) (*protocom.Response, error) {
	panic("implement me")
}

func (s *server) Set(ctx context.Context, request *protocom.ExternalSetRequest) (*protocom.Response, error) {
	panic("implement me")
}
