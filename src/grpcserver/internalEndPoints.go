package grpcserver

import (
	"context"
	"github.com/hisitra/hedron/src/almanac"
	"github.com/hisitra/hedron/src/comcn"
	eq "github.com/hisitra/hedron/src/executionqueue"
	iot "github.com/hisitra/hedron/src/iotranslator"
)

func (s *server) Get(ctx context.Context, inMsg *comcn.InternalMessage) (*comcn.Output, error) {
	req, err := iot.DecodeRequest(inMsg.Value)
	if err != nil {
		return iot.BadRequestResponse(""), nil
	}

	if !eq.ExecutionAllowable(req) {
		return iot.EarlierRequestFoundResponse(), nil
	}

	record, res := almanac.Get(req.Key)
	if res.Code != 200 {
		return res, nil
	}
	return iot.NewResponse(200, "success", record), nil
}

func (s *server) Set(ctx context.Context, inMsg *comcn.InternalMessage) (*comcn.Output, error) {
	res := almanac.Set(inMsg.Value)
	if res.Code != 200 {
		return res, nil
	}
	return iot.SuccessResponse(""), nil
}
