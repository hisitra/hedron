package grpcserver

import (
	"context"
	"github.com/hisitra/hedron/src/comcn"
	iot "github.com/hisitra/hedron/src/iotranslator"
	"github.com/hisitra/ventager"
)

func (s *server) Create(ctx context.Context, message *comcn.Message) (*comcn.Message, error) {
	newReq, res := iot.NewRequest(message.Value, "C")
	if !res.IsSuccessful() {
		return &comcn.Message{Value: res.Marshal()}, nil
	}
	ventager.Fire("request-arrival", newReq)
	finalRes, ok := (<-ventager.Listen("request-executed-"+newReq.ID)).(*iot.Response)
	if !ok {
		return &comcn.Message{Value: iot.InternalServerResponse("").Marshal()}, nil
	}
	return &comcn.Message{Value: finalRes.Marshal()}, nil
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
