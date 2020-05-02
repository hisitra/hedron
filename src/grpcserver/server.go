package grpcserver

import (
	"context"
	"github.com/hisitra/hedron/src/comcn"
	"github.com/hisitra/hedron/src/configs"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	Port string
}

type Server interface {
	Start()

	Get(ctx context.Context, inMsg *comcn.InternalMessage) (*comcn.Output, error)
	Set(ctx context.Context, inMsg *comcn.InternalMessage) (*comcn.Output, error)

	Create(ctx context.Context, message *comcn.Input) (*comcn.Output, error)
	Read(ctx context.Context, message *comcn.Input) (*comcn.Output, error)
	Update(ctx context.Context, message *comcn.Input) (*comcn.Output, error)
	Delete(ctx context.Context, message *comcn.Input) (*comcn.Output, error)
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
	comcn.RegisterExternalServer(grpcServer, s)
	comcn.RegisterInternalServer(grpcServer, s)

	log.Println("Starting Hedron Node:", configs.Node.Name, "gRPC Server at PORT:", configs.Server.Port)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
