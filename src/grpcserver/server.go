package grpcserver

import (
	"context"
	"github.com/hisitra/hedron/src/configs"
	"github.com/hisitra/hedron/src/logan"
	protocom "github.com/hisitra/hedron/src/protos"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	Port string
}

type Server interface {
	Start()

	Get(ctx context.Context, request *protocom.ExternalGetRequest) (*protocom.Response, error)
	Set(ctx context.Context, request *protocom.ExternalSetRequest) (*protocom.Response, error)
}

func New() Server {
	return &server{Port: configs.Server.Port}
}

func (s *server) Start() {
	listener, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	// register the servers here
	protocom.RegisterExternalServer(grpcServer, s)

	logan.Info.Println("Starting gRPC Server at PORT:", s.Port)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
