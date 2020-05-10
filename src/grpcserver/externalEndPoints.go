package grpcserver

import (
	"context"
	"github.com/hisitra/hedron/src/comcn"
)

func (s *server) Create(ctx context.Context, message *comcn.Message) (*comcn.Message, error) {
	panic("implement me")
}

func (s *server) Read(ctx context.Context, message *comcn.Message) (*comcn.Message, error) {
	panic("implement me")
}

func (s *server) Update(ctx context.Context, message *comcn.Message) (*comcn.Message, error) {
	panic("implement me")
}

func (s *server) Delete(ctx context.Context, message *comcn.Message) (*comcn.Message, error) {
	panic("implement me")
}
