package grpcserver

import (
	"context"
	"github.com/hisitra/hedron/src/comcn"
)

func (s *server) Get(ctx context.Context, message *comcn.Message) (*comcn.Message, error) {
	panic("implement me")
}

func (s *server) Set(ctx context.Context, message *comcn.Message) (*comcn.Message, error) {
	panic("implement me")
}
