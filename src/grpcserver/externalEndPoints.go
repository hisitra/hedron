package grpcserver

import (
	"context"
	"github.com/hisitra/hedron/src/logan"
	protocom "github.com/hisitra/hedron/src/protos"
)

func (s *server) Get(ctx context.Context, request *protocom.ExternalGetRequest) (*protocom.Response, error) {
	logan.Info.Println("Get request arrived")
	return &protocom.Response{
		StatusCode: 0,
		CustomCode: "",
		Message:    "",
		Data:       nil,
	}, nil
}

func (s *server) Set(ctx context.Context, request *protocom.ExternalSetRequest) (*protocom.Response, error) {
	logan.Info.Println("Set request arrived")
	return &protocom.Response{
		StatusCode: 0,
		CustomCode: "",
		Message:    "",
		Data:       nil,
	}, nil
}
