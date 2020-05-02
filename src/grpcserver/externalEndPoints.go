package grpcserver

import (
	"context"
	"github.com/hisitra/hedron/src/comcn"
	eq "github.com/hisitra/hedron/src/executionqueue"
	iot "github.com/hisitra/hedron/src/iotranslator"
)

func (s *server) Create(ctx context.Context, input *comcn.Input) (*comcn.Output, error) {
	req, res := iot.NewRequest(input, "C")
	if res.Code != 200 {
		return res, nil
	}
	eq.ArrivalChan <- req
	return <-req.OnResponse(), nil
}

func (s *server) Read(ctx context.Context, input *comcn.Input) (*comcn.Output, error) {
	req, res := iot.NewRequest(input, "R")
	if res.Code != 200 {
		return res, nil
	}
	eq.ArrivalChan <- req
	return <-req.OnResponse(), nil
}

func (s *server) Update(ctx context.Context, input *comcn.Input) (*comcn.Output, error) {
	req, res := iot.NewRequest(input, "U")
	if res.Code != 200 {
		return res, nil
	}
	eq.ArrivalChan <- req
	return <-req.OnResponse(), nil
}

func (s *server) Delete(ctx context.Context, input *comcn.Input) (*comcn.Output, error) {
	req, res := iot.NewRequest(input, "D")
	if res.Code != 200 {
		return res, nil
	}
	eq.ArrivalChan <- req
	return <-req.OnResponse(), nil
}
